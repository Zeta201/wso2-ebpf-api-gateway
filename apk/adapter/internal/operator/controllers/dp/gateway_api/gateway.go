package gateway_api

import (
	"context"

	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/sirupsen/logrus"
	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/helpers"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const owningGatewayLabel = "io.cilium.gateway/owning-gateway"
const lastTransitionTime = "LastTransitionTime"

// gatewayReconciler reconciles a Gateway object
type gatewayReconciler struct {
	client.Client
	Scheme             *runtime.Scheme
	SecretsNamespace   string
	IdleTimeoutSeconds int
}

func newGatewayReconciler(mgr ctrl.Manager, secretsNamespace string, idleTimeoutSeconds int) *gatewayReconciler {
	return &gatewayReconciler{
		Client:             mgr.GetClient(),
		Scheme:             mgr.GetScheme(),
		SecretsNamespace:   secretsNamespace,
		IdleTimeoutSeconds: idleTimeoutSeconds,
	}
}

// SetupWithManager sets up the controller with the Manager.
// The reconciler will be triggered by Gateway, or any cilium-managed GatewayClass events
func (r *gatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	hasMatchingControllerFn := hasMatchingController(context.Background(), r.Client, controllerName)
	return ctrl.NewControllerManagedBy(mgr).
		// Watch its own resource
		For(&gwapiv1b1.Gateway{},
			builder.WithPredicates(predicate.NewPredicateFuncs(hasMatchingControllerFn))).
		// Watch GatewayClass resources, which are linked to Gateway
		Watches(&gwapiv1b1.GatewayClass{},
			r.enqueueRequestForOwningGatewayClass(),
			builder.WithPredicates(predicate.NewPredicateFuncs(hasMatchingControllerFn))).
		// Watch related LB service for status
		Watches(&corev1.Service{},
			r.enqueueRequestForOwningResource(),
			builder.WithPredicates(predicate.NewPredicateFuncs(func(object client.Object) bool {
				_, found := object.GetLabels()[owningGatewayLabel]
				return found
			}))).
		// Watch HTTP Route status changes, there is one assumption that any change in spec will
		// always update status always at least for observedGeneration value.
		Watches(&gwapiv1b1.HTTPRoute{},
			r.enqueueRequestForOwningHTTPRoute(),
			builder.WithPredicates(onlyStatusChanged())).
		// Watch TLS Route status changes, there is one assumption that any change in spec will
		// always update status always at least for observedGeneration value.
		Watches(&gwapiv1a2.TLSRoute{},
			r.enqueueRequestForOwningTLSRoute(),
			builder.WithPredicates(onlyStatusChanged())).
		// Watch GRPCRoute status changes, there is one assumption that any change in spec will
		// always update status always at least for observedGeneration value.
		Watches(&gwapiv1a2.GRPCRoute{},
			r.enqueueRequestForOwningGRPCRoute(),
			builder.WithPredicates(onlyStatusChanged())).
		// Watch related secrets used to configure TLS
		Watches(&corev1.Secret{},
			r.enqueueRequestForTLSSecret(),
			builder.WithPredicates(predicate.NewPredicateFuncs(r.usedInGateway))).
		// Watch related namespace in allowed namespaces
		Watches(&corev1.Namespace{},
			r.enqueueRequestForAllowedNamespace()).
		// Watch created and owned resources
		Owns(&ciliumv2.CiliumEnvoyConfig{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.Endpoints{}).
		Complete(r)
}

// enqueueRequestForOwningGatewayClass returns an event handler for all Gateway objects
// belonging to the given GatewayClass.
func (r *gatewayReconciler) enqueueRequestForOwningGatewayClass() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		scopedLog := log.WithFields(logrus.Fields{
			logfields.Controller: gateway,
			logfields.Resource:   a.GetName(),
		})
		var reqs []reconcile.Request
		gwList := &gwapiv1b1.GatewayList{}
		if err := r.Client.List(ctx, gwList); err != nil {
			scopedLog.Error("Unable to list Gateways")
			return nil
		}

		for _, gw := range gwList.Items {
			if gw.Spec.GatewayClassName != gwapiv1b1.ObjectName(a.GetName()) {
				continue
			}
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: gw.Namespace,
					Name:      gw.Name,
				},
			}
			reqs = append(reqs, req)
			scopedLog.WithFields(logrus.Fields{
				logfields.K8sNamespace: gw.GetNamespace(),
				logfields.Resource:     gw.GetName(),
			}).Info("Queueing gateway")
		}
		return reqs
	})
}

// enqueueRequestForOwningResource returns an event handler for all Gateway objects having
// owningGatewayLabel
func (r *gatewayReconciler) enqueueRequestForOwningResource() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		scopedLog := log.WithFields(logrus.Fields{
			logfields.Controller: "gateway",
			logfields.Resource:   a.GetName(),
		})

		key, found := a.GetLabels()[owningGatewayLabel]
		if !found {
			return nil
		}

		scopedLog.WithFields(logrus.Fields{
			logfields.K8sNamespace: a.GetNamespace(),
			logfields.Resource:     a.GetName(),
			"gateway":              key,
		}).Info("Enqueued gateway for owning service")

		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Namespace: a.GetNamespace(),
					Name:      key,
				},
			},
		}
	})
}

// enqueueRequestForOwningHTTPRoute returns an event handler for any changes with HTTP Routes
// belonging to the given Gateway
func (r *gatewayReconciler) enqueueRequestForOwningHTTPRoute() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		hr, ok := a.(*gwapiv1b1.HTTPRoute)
		if !ok {
			return nil
		}

		return getReconcileRequestsForRoute(context.Background(), r.Client, a, hr.Spec.CommonRouteSpec)
	})
}

// enqueueRequestForOwningTLSRoute returns an event handler for any changes with TLS Routes
// belonging to the given Gateway
func (r *gatewayReconciler) enqueueRequestForOwningTLSRoute() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		hr, ok := a.(*gwapiv1a2.TLSRoute)
		if !ok {
			return nil
		}

		return getReconcileRequestsForRoute(context.Background(), r.Client, a, hr.Spec.CommonRouteSpec)
	})
}

func getReconcileRequestsForRoute(ctx context.Context, c client.Client, object metav1.Object, route gwapiv1b1.CommonRouteSpec) []reconcile.Request {
	var reqs []reconcile.Request

	scopedLog := log.WithFields(logrus.Fields{
		logfields.Controller: gateway,
		logfields.Resource: types.NamespacedName{
			Namespace: object.GetNamespace(),
			Name:      object.GetName(),
		},
	})

	for _, parent := range route.ParentRefs {
		if !helpers.IsGateway(parent) {
			continue
		}

		ns := helpers.NamespaceDerefOr(parent.Namespace, object.GetNamespace())

		gw := &gwapiv1b1.Gateway{}
		if err := c.Get(ctx, types.NamespacedName{
			Namespace: ns,
			Name:      string(parent.Name),
		}, gw); err != nil {
			if !k8serrors.IsNotFound(err) {
				scopedLog.WithError(err).Error("Failed to get Gateway")
			}
			continue
		}

		if !hasMatchingController(ctx, c, controllerName)(gw) {
			scopedLog.Debug("Gateway does not have matching controller, skipping")
			continue
		}

		scopedLog.WithFields(logrus.Fields{
			logfields.K8sNamespace: ns,
			logfields.Resource:     parent.Name,
			logfields.Route:        object.GetName(),
		}).Info("Enqueued gateway for Route")

		reqs = append(reqs, reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: ns,
				Name:      string(parent.Name),
			},
		})
	}

	return reqs
}

// enqueueRequestForOwningTLSCertificate returns an event handler for any changes with TLS secrets
func (r *gatewayReconciler) enqueueRequestForTLSSecret() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		gateways := getGatewaysForSecret(ctx, r.Client, a)
		reqs := make([]reconcile.Request, 0, len(gateways))
		for _, gw := range gateways {
			reqs = append(reqs, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: gw.GetNamespace(),
					Name:      gw.GetName(),
				},
			})
		}
		return reqs
	})
}

// enqueueRequestForAllowedNamespace returns an event handler for any changes
// with allowed namespaces
func (r *gatewayReconciler) enqueueRequestForAllowedNamespace() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, ns client.Object) []reconcile.Request {
		gateways := getGatewaysForNamespace(ctx, r.Client, ns)
		reqs := make([]reconcile.Request, 0, len(gateways))
		for _, gw := range gateways {
			reqs = append(reqs, reconcile.Request{
				NamespacedName: gw,
			})
		}
		return reqs
	})
}

func (r *gatewayReconciler) usedInGateway(obj client.Object) bool {
	return len(getGatewaysForSecret(context.Background(), r.Client, obj)) > 0
}

// enqueueRequestForOwningGRPCRoute returns an event handler for any changes with GRPC Routes
// belonging to the given Gateway
func (r *gatewayReconciler) enqueueRequestForOwningGRPCRoute() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, a client.Object) []reconcile.Request {
		gr, ok := a.(*gwapiv1a2.GRPCRoute)
		if !ok {
			return nil
		}

		return getReconcileRequestsForRoute(ctx, r.Client, a, gr.Spec.CommonRouteSpec)
	})
}
