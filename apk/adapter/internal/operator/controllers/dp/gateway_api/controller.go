package gateway_api

import (
	"context"

	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/helpers"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	// controllerName is the gateway controller name used in cilium.
	controllerName      = "io.cilium/gateway-controller"
	backendServiceIndex = "backendServiceIndex"
	gatewayIndex        = "gatewayIndex"
)

func hasMatchingController(ctx context.Context, c client.Client, controllerName string) func(object client.Object) bool {
	return func(obj client.Object) bool {
		scopedLog := log.WithFields(logrus.Fields{
			logfields.Controller: gateway,
			logfields.Resource:   obj.GetName(),
		})
		gw, ok := obj.(*gwapiv1b1.Gateway)
		if !ok {
			return false
		}

		gwc := &gwapiv1b1.GatewayClass{}
		key := types.NamespacedName{Name: string(gw.Spec.GatewayClassName)}
		if err := c.Get(ctx, key, gwc); err != nil {
			scopedLog.WithError(err).Error("Unable to get GatewayClass")
			return false
		}

		return string(gwc.Spec.ControllerName) == controllerName
	}
}

func getGatewaysForSecret(ctx context.Context, c client.Client, obj client.Object) []*gwapiv1b1.Gateway {
	scopedLog := log.WithFields(logrus.Fields{
		logfields.Controller: gateway,
		logfields.Resource:   obj.GetName(),
	})

	gwList := &gwapiv1b1.GatewayList{}
	if err := c.List(ctx, gwList); err != nil {
		scopedLog.WithError(err).Warn("Unable to list Gateways")
		return nil
	}

	var gateways []*gwapiv1b1.Gateway
	for _, gw := range gwList.Items {
		gwCopy := gw
		for _, l := range gw.Spec.Listeners {
			if l.TLS == nil {
				continue
			}

			for _, cert := range l.TLS.CertificateRefs {
				if !helpers.IsSecret(cert) {
					continue
				}
				ns := helpers.NamespaceDerefOr(cert.Namespace, gw.GetNamespace())
				if string(cert.Name) == obj.GetName() && ns == obj.GetNamespace() {
					gateways = append(gateways, &gwCopy)
				}
			}
		}
	}
	return gateways
}

func getGatewaysForNamespace(ctx context.Context, c client.Client, ns client.Object) []types.NamespacedName {
	scopedLog := log.WithFields(logrus.Fields{
		logfields.Controller:   gateway,
		logfields.K8sNamespace: ns.GetName(),
	})

	gwList := &gwapiv1b1.GatewayList{}
	if err := c.List(ctx, gwList); err != nil {
		scopedLog.WithError(err).Warn("Unable to list Gateways")
		return nil
	}

	var gateways []types.NamespacedName
	for _, gw := range gwList.Items {
		for _, l := range gw.Spec.Listeners {
			if l.AllowedRoutes == nil || l.AllowedRoutes.Namespaces == nil {
				continue
			}

			switch *l.AllowedRoutes.Namespaces.From {
			case gwapiv1b1.NamespacesFromAll:
				gateways = append(gateways, client.ObjectKey{
					Namespace: gw.GetNamespace(),
					Name:      gw.GetName(),
				})
			case gwapiv1b1.NamespacesFromSame:
				if ns.GetName() == gw.GetNamespace() {
					gateways = append(gateways, client.ObjectKey{
						Namespace: gw.GetNamespace(),
						Name:      gw.GetName(),
					})
				}
			case gwapiv1b1.NamespacesFromSelector:
				nsList := &corev1.NamespaceList{}
				err := c.List(ctx, nsList, client.MatchingLabels(l.AllowedRoutes.Namespaces.Selector.MatchLabels))
				if err != nil {
					scopedLog.WithError(err).Warn("Unable to list Namespaces")
					return nil
				}
				for _, item := range nsList.Items {
					if item.GetName() == ns.GetName() {
						gateways = append(gateways, client.ObjectKey{
							Namespace: gw.GetNamespace(),
							Name:      gw.GetName(),
						})
					}
				}
			}
		}
	}
	return gateways
}

// onlyStatusChanged returns true if and only if there is status change for underlying objects.
// Supported objects are GatewayClass, Gateway, HTTPRoute and GRPCRoute
func onlyStatusChanged() predicate.Predicate {
	option := cmpopts.IgnoreFields(metav1.Condition{}, "LastTransitionTime")
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			switch e.ObjectOld.(type) {
			case *gwapiv1b1.GatewayClass:
				o, _ := e.ObjectOld.(*gwapiv1b1.GatewayClass)
				n, ok := e.ObjectNew.(*gwapiv1b1.GatewayClass)
				if !ok {
					return false
				}
				return !cmp.Equal(o.Status, n.Status, option)
			case *gwapiv1b1.Gateway:
				o, _ := e.ObjectOld.(*gwapiv1b1.Gateway)
				n, ok := e.ObjectNew.(*gwapiv1b1.Gateway)
				if !ok {
					return false
				}
				return !cmp.Equal(o.Status, n.Status, option)
			case *gwapiv1b1.HTTPRoute:
				o, _ := e.ObjectOld.(*gwapiv1b1.HTTPRoute)
				n, ok := e.ObjectNew.(*gwapiv1b1.HTTPRoute)
				if !ok {
					return false
				}
				return !cmp.Equal(o.Status, n.Status, option)
			case *gwapiv1a2.TLSRoute:
				o, _ := e.ObjectOld.(*gwapiv1a2.TLSRoute)
				n, ok := e.ObjectNew.(*gwapiv1a2.TLSRoute)
				if !ok {
					return false
				}
				return !cmp.Equal(o.Status, n.Status, option)
			case *gwapiv1a2.GRPCRoute:
				o, _ := e.ObjectOld.(*gwapiv1a2.GRPCRoute)
				n, ok := e.ObjectNew.(*gwapiv1a2.GRPCRoute)
				if !ok {
					return false
				}
				return !cmp.Equal(o.Status, n.Status, option)
			default:
				return false
			}
		},
	}
}