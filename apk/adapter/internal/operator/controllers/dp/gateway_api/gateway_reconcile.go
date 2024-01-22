package gateway_api

import (
	"context"
	"encoding/pem"
	"fmt"

	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"
	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/helpers"
	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/model"
	ingestion "github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/model/igestion"
	translation "github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/model/translation/gateway-api"
	controllerruntime "github.com/wso2/apk/adapter/pkg/controller-runtime"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	dpv1alpha1 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func (r *gatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	scopedLog := log.WithContext(ctx).WithFields(logrus.Fields{
		logfields.Controller: gateway,
		logfields.Resource:   req.NamespacedName,
	})
	scopedLog.Info("Reconciling Gateway")

	original := &gwapiv1b1.Gateway{}

	if err := r.Client.Get(ctx, req.NamespacedName, original); err != nil {
		if k8serrors.IsNotFound(err) {
			return controllerruntime.Success()
		}
		scopedLog.WithError(err).Error("Unable to get Gateway")
		return controllerruntime.Fail(err)
	}
	// Ignore deleting Gateway, this can happen when foregroundDeletion is enabled
	// The reconciliation loop will automatically kick off for related Gateway resources.
	if original.GetDeletionTimestamp() != nil {
		scopedLog.Info("Gateway is being deleted, doing nothing")
		return controllerruntime.Success()
	}

	gw := original.DeepCopy()

	// Step 2: Gather all required information for the ingestion model
	gwc := &gwapiv1b1.GatewayClass{}
	if err := r.Client.Get(ctx, client.ObjectKey{Name: string(gw.Spec.GatewayClassName)}, gwc); err != nil {
		scopedLog.WithField(gatewayClass, gw.Spec.GatewayClassName).
			WithError(err).
			Error("Unable to get GatewayClass")
		if k8serrors.IsNotFound(err) {
			setGatewayAccepted(gw, false, "GatewayClass does not exist")
			return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
		}
		setGatewayAccepted(gw, false, "Unable to get GatewayClass")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	backendsList := &dpv1alpha1.BackendList{}
	if err := r.Client.List(ctx, backendsList); err != nil {
		scopedLog.WithError(err).Error("Unable to list Backends")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	httpRouteList := &gwapiv1b1.HTTPRouteList{}
	if err := r.Client.List(ctx, httpRouteList); err != nil {
		scopedLog.WithError(err).Error("Unable to list HTTPRoutes")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	grpcRouteList := &gwapiv1a2.GRPCRouteList{}
	if err := r.Client.List(ctx, grpcRouteList); err != nil {
		scopedLog.WithError(err).Error("Unable to list GRPCRoutes")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	tlsRouteList := &gwapiv1a2.TLSRouteList{}
	if err := r.Client.List(ctx, tlsRouteList); err != nil {
		scopedLog.WithError(err).Error("Unable to list TLSRoutes")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	servicesList := &corev1.ServiceList{}
	if err := r.Client.List(ctx, servicesList); err != nil {
		scopedLog.WithError(err).Error("Unable to list Services")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	grants := &gwapiv1b1.ReferenceGrantList{}
	if err := r.Client.List(ctx, grants); err != nil {
		scopedLog.WithError(err).Error("Unable to list ReferenceGrants")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	httpListeners, tlsListeners := ingestion.GatewayAPI(ingestion.Input{
		GatewayClass:    *gwc,
		Gateway:         *gw,
		HTTPRoutes:      r.filterHTTPRoutesByGateway(ctx, gw, httpRouteList.Items),
		TLSRoutes:       r.filterTLSRoutesByGateway(ctx, gw, tlsRouteList.Items),
		GRPCRoutes:      r.filterGRPCRoutesByGateway(ctx, gw, grpcRouteList.Items),
		Services:        servicesList.Items,
		ReferenceGrants: grants.Items,
	})

	if err := r.setListenerStatus(ctx, gw, httpRouteList, tlsRouteList); err != nil {
		scopedLog.WithError(err).Error("Unable to set listener status")
		setGatewayAccepted(gw, false, "Unable to set listener status")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}
	setGatewayAccepted(gw, true, "Gateway successfully scheduled")

	// Step 3: Translate the listeners into Cilium model
	cec, svc, ep, err := translation.NewTranslator(r.SecretsNamespace, r.IdleTimeoutSeconds).Translate(&model.Model{HTTP: httpListeners, TLS: tlsListeners})

	if err != nil {
		scopedLog.WithError(err).Error("Unable to translate resources")
		setGatewayAccepted(gw, false, "Unable to translate resources")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	if err = r.ensureService(ctx, svc); err != nil {
		scopedLog.WithError(err).Error("Unable to create Service")
		setGatewayAccepted(gw, false, "Unable to create Service resource")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	if err = r.ensureEndpoints(ctx, ep); err != nil {
		scopedLog.WithError(err).Error("Unable to ensure Endpoints")
		setGatewayAccepted(gw, false, "Unable to ensure Endpoints resource")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	if err = r.ensureEnvoyConfig(ctx, cec); err != nil {
		scopedLog.WithError(err).Error("Unable to ensure CiliumEnvoyConfig")
		setGatewayAccepted(gw, false, "Unable to ensure CEC resource")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	// Step 4: Update the status of the Gateway
	if err = r.setAddressStatus(ctx, gw); err != nil {
		scopedLog.WithError(err).Error("Address is not ready")
		setGatewayProgrammed(gw, false, "Address is not ready")
		return r.handleReconcileErrorWithStatus(ctx, err, original, gw)
	}

	setGatewayProgrammed(gw, true, "Gateway successfully reconciled")
	if err := r.updateStatus(ctx, original, gw); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to update Gateway status: %w", err)
	}

	scopedLog.Info("Successfully reconciled Gateway")
	return controllerruntime.Success()

}

func (r *gatewayReconciler) handleReconcileErrorWithStatus(ctx context.Context, reconcileErr error, original *gwapiv1b1.Gateway, modified *gwapiv1b1.Gateway) (ctrl.Result, error) {
	if err := r.updateStatus(ctx, original, modified); err != nil {
		return controllerruntime.Fail(fmt.Errorf("failed to update Gateway status while handling the reconcile error %w: %w", reconcileErr, err))
	}

	return controllerruntime.Fail(reconcileErr)
}

func (r *gatewayReconciler) updateStatus(ctx context.Context, original *gwapiv1b1.Gateway, new *gwapiv1b1.Gateway) error {
	oldStatus := original.Status.DeepCopy()
	newStatus := new.Status.DeepCopy()

	if cmp.Equal(oldStatus, newStatus, cmpopts.IgnoreFields(metav1.Condition{}, lastTransitionTime)) {
		return nil
	}
	return r.Client.Status().Update(ctx, new)
}

func (r *gatewayReconciler) filterHTTPRoutesByGateway(ctx context.Context, gw *gwapiv1b1.Gateway, routes []gwapiv1b1.HTTPRoute) []gwapiv1b1.HTTPRoute {
	var filtered []gwapiv1b1.HTTPRoute
	for _, route := range routes {
		if isAttachable(ctx, gw, &route, route.Status.Parents) && isAllowed(ctx, r.Client, gw, &route) && len(computeHosts(gw, route.Spec.Hostnames)) > 0 {
			filtered = append(filtered, route)
		}
	}
	return filtered
}

func isAttachable(_ context.Context, gw *gwapiv1b1.Gateway, route metav1.Object, parents []gwapiv1b1.RouteParentStatus) bool {
	for _, rps := range parents {
		if helpers.NamespaceDerefOr(rps.ParentRef.Namespace, route.GetNamespace()) != gw.GetNamespace() ||
			string(rps.ParentRef.Name) != gw.GetName() {
			continue
		}

		for _, cond := range rps.Conditions {
			if cond.Type == string(gwapiv1b1.RouteConditionAccepted) && cond.Status == metav1.ConditionTrue {
				return true
			}

			if cond.Type == string(gwapiv1b1.RouteConditionResolvedRefs) && cond.Status == metav1.ConditionFalse {
				return true
			}
		}
	}
	return false
}

func (r *gatewayReconciler) filterTLSRoutesByGateway(ctx context.Context, gw *gwapiv1b1.Gateway, routes []gwapiv1a2.TLSRoute) []gwapiv1a2.TLSRoute {
	var filtered []gwapiv1a2.TLSRoute
	for _, route := range routes {
		if isAttachable(ctx, gw, &route, route.Status.Parents) && isAllowed(ctx, r.Client, gw, &route) && len(computeHosts(gw, route.Spec.Hostnames)) > 0 {
			filtered = append(filtered, route)
		}
	}
	return filtered
}

func (r *gatewayReconciler) filterTLSRoutesByListener(ctx context.Context, gw *gwapiv1b1.Gateway, listener *gwapiv1b1.Listener, routes []gwapiv1a2.TLSRoute) []gwapiv1a2.TLSRoute {
	var filtered []gwapiv1a2.TLSRoute
	for _, route := range routes {
		if isAttachable(ctx, gw, &route, route.Status.Parents) &&
			isAllowed(ctx, r.Client, gw, &route) &&
			len(computeHostsForListener(listener, route.Spec.Hostnames)) > 0 &&
			parentRefMatched(gw, listener, route.GetNamespace(), route.Spec.ParentRefs) {
			filtered = append(filtered, route)
		}
	}
	return filtered
}

func parentRefMatched(gw *gwapiv1b1.Gateway, listener *gwapiv1b1.Listener, routeNamespace string, refs []gwapiv1b1.ParentReference) bool {
	for _, ref := range refs {
		if string(ref.Name) == gw.GetName() && gw.GetNamespace() == helpers.NamespaceDerefOr(ref.Namespace, routeNamespace) {
			if ref.SectionName == nil && ref.Port == nil {
				return true
			}
			sectionNameCheck := ref.SectionName == nil || *ref.SectionName == listener.Name
			portCheck := ref.Port == nil || *ref.Port == listener.Port
			if sectionNameCheck && portCheck {
				return true
			}
		}
	}
	return false
}
func (r *gatewayReconciler) filterGRPCRoutesByGateway(ctx context.Context, gw *gwapiv1b1.Gateway, routes []gwapiv1a2.GRPCRoute) []gwapiv1a2.GRPCRoute {
	var filtered []gwapiv1a2.GRPCRoute
	for _, route := range routes {
		if isAttachable(ctx, gw, &route, route.Status.Parents) && isAllowed(ctx, r.Client, gw, &route) && len(computeHosts(gw, route.Spec.Hostnames)) > 0 {
			filtered = append(filtered, route)
		}
	}
	return filtered
}

func (r *gatewayReconciler) setListenerStatus(ctx context.Context, gw *gwapiv1b1.Gateway, httpRoutes *gwapiv1b1.HTTPRouteList, tlsRoutes *gwapiv1a2.TLSRouteList) error {
	grants := &gwapiv1b1.ReferenceGrantList{}
	if err := r.Client.List(ctx, grants); err != nil {
		return fmt.Errorf("failed to retrieve reference grants: %w", err)
	}

	for _, l := range gw.Spec.Listeners {
		isValid := true

		// SupportedKinds is a required field, so we can't declare it as nil.
		supportedKinds := []gwapiv1b1.RouteGroupKind{}
		invalidRouteKinds := false
		protoGroup, protoKind := getSupportedGroupKind(l.Protocol)

		if l.AllowedRoutes != nil && len(l.AllowedRoutes.Kinds) != 0 {
			for _, k := range l.AllowedRoutes.Kinds {
				if groupDerefOr(k.Group, gwapiv1b1.GroupName) == string(*protoGroup) &&
					k.Kind == protoKind {
					supportedKinds = append(supportedKinds, k)
				} else {
					invalidRouteKinds = true
				}
			}
		} else {
			g, k := getSupportedGroupKind(l.Protocol)
			supportedKinds = []gwapiv1b1.RouteGroupKind{
				{
					Group: g,
					Kind:  k,
				},
			}
		}
		var conds []metav1.Condition
		if invalidRouteKinds {
			conds = append(conds, gatewayListenerInvalidRouteKinds(gw, "Invalid Route Kinds"))
			isValid = false
		} else {
			conds = append(conds, gatewayListenerProgrammedCondition(gw, true, "Listener Programmed"))
			conds = append(conds, gatewayListenerAcceptedCondition(gw, true, "Listener Accepted"))
			conds = append(conds, metav1.Condition{
				Type:               string(gwapiv1b1.ListenerConditionResolvedRefs),
				Status:             metav1.ConditionTrue,
				Reason:             string(gwapiv1b1.ListenerReasonResolvedRefs),
				Message:            "Resolved Refs",
				LastTransitionTime: metav1.Now(),
			})
		}

		if l.TLS != nil {
			for _, cert := range l.TLS.CertificateRefs {
				if !helpers.IsSecret(cert) {
					conds = merge(conds, metav1.Condition{
						Type:               string(gwapiv1b1.ListenerConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						Reason:             string(gwapiv1b1.ListenerReasonInvalidCertificateRef),
						Message:            "Invalid CertificateRef",
						LastTransitionTime: metav1.Now(),
					})
					isValid = false
					break
				}

				if !helpers.IsSecretReferenceAllowed(gw.Namespace, cert, gwapiv1b1.SchemeGroupVersion.WithKind("Gateway"), grants.Items) {
					conds = merge(conds, metav1.Condition{
						Type:               string(gwapiv1b1.ListenerConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						Reason:             string(gwapiv1b1.ListenerReasonRefNotPermitted),
						Message:            "CertificateRef is not permitted",
						LastTransitionTime: metav1.Now(),
					})
					isValid = false
					break
				}

				if err := validateTLSSecret(ctx, r.Client, helpers.NamespaceDerefOr(cert.Namespace, gw.GetNamespace()), string(cert.Name)); err != nil {
					conds = merge(conds, metav1.Condition{
						Type:               string(gwapiv1b1.ListenerConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						Reason:             string(gwapiv1b1.ListenerReasonInvalidCertificateRef),
						Message:            "Invalid CertificateRef",
						LastTransitionTime: metav1.Now(),
					})
					isValid = false
					break
				}
			}
		}

		if !isValid {
			conds = merge(conds, metav1.Condition{
				Type:               string(gwapiv1b1.ListenerConditionProgrammed),
				Status:             metav1.ConditionFalse,
				Reason:             string(gwapiv1b1.ListenerReasonInvalid),
				Message:            "Invalid CertificateRef",
				LastTransitionTime: metav1.Now(),
			})
		}

		var attachedRoutes int32
		attachedRoutes += int32(len(r.filterHTTPRoutesByListener(ctx, gw, &l, httpRoutes.Items)))
		attachedRoutes += int32(len(r.filterTLSRoutesByListener(ctx, gw, &l, tlsRoutes.Items)))

		found := false
		for i := range gw.Status.Listeners {
			if l.Name == gw.Status.Listeners[i].Name {
				found = true
				gw.Status.Listeners[i].SupportedKinds = supportedKinds
				gw.Status.Listeners[i].Conditions = conds
				gw.Status.Listeners[i].AttachedRoutes = attachedRoutes
				break
			}
		}
		if !found {
			gw.Status.Listeners = append(gw.Status.Listeners, gwapiv1b1.ListenerStatus{
				Name:           l.Name,
				SupportedKinds: supportedKinds,
				Conditions:     conds,
				AttachedRoutes: attachedRoutes,
			})
		}
	}

	// filter listener status to only have active listeners
	var newListenersStatus []gwapiv1b1.ListenerStatus
	for _, ls := range gw.Status.Listeners {
		for _, l := range gw.Spec.Listeners {
			if ls.Name == l.Name {
				newListenersStatus = append(newListenersStatus, ls)
				break
			}
		}
	}
	gw.Status.Listeners = newListenersStatus
	return nil
}
func validateTLSSecret(ctx context.Context, c client.Client, namespace, name string) error {
	secret := &corev1.Secret{}
	if err := c.Get(ctx, client.ObjectKey{
		Namespace: namespace,
		Name:      name,
	}, secret); err != nil {
		return err
	}

	if !isValidPemFormat(secret.Data[corev1.TLSCertKey]) {
		return fmt.Errorf("invalid certificate")
	}

	if !isValidPemFormat(secret.Data[corev1.TLSPrivateKeyKey]) {
		return fmt.Errorf("invalid private key")
	}
	return nil
}

// isValidPemFormat checks if the given byte array is a valid PEM format.
// This function is not intended to be used for validating the actual
// content of the PEM block.
func isValidPemFormat(b []byte) bool {
	if len(b) == 0 {
		return false
	}

	p, rest := pem.Decode(b)
	if p == nil {
		return false
	}
	if len(rest) == 0 {
		return true
	}
	return isValidPemFormat(rest)
}

func (r *gatewayReconciler) filterHTTPRoutesByListener(ctx context.Context, gw *gwapiv1b1.Gateway, listener *gwapiv1b1.Listener, routes []gwapiv1b1.HTTPRoute) []gwapiv1b1.HTTPRoute {
	var filtered []gwapiv1b1.HTTPRoute
	for _, route := range routes {
		if isAttachable(ctx, gw, &route, route.Status.Parents) &&
			isAllowed(ctx, r.Client, gw, &route) &&
			len(computeHostsForListener(listener, route.Spec.Hostnames)) > 0 &&
			parentRefMatched(gw, listener, route.GetNamespace(), route.Spec.ParentRefs) {
			filtered = append(filtered, route)
		}
	}
	return filtered
}

func (r *gatewayReconciler) ensureService(ctx context.Context, desired *corev1.Service) error {
	svc := desired.DeepCopy()
	_, err := controllerutil.CreateOrPatch(ctx, r.Client, svc, func() error {
		// Save and restore loadBalancerClass
		// e.g. if a mutating webhook writes this field
		lbClass := svc.Spec.LoadBalancerClass

		svc.Spec = desired.Spec
		svc.OwnerReferences = desired.OwnerReferences
		setMergedLabelsAndAnnotations(svc, desired)

		// Ignore the loadBalancerClass if it was set by a mutating webhook
		svc.Spec.LoadBalancerClass = lbClass
		return nil
	})
	return err
}

func (r *gatewayReconciler) ensureEndpoints(ctx context.Context, desired *corev1.Endpoints) error {
	ep := desired.DeepCopy()
	_, err := controllerutil.CreateOrPatch(ctx, r.Client, ep, func() error {
		ep.Subsets = desired.Subsets
		ep.OwnerReferences = desired.OwnerReferences
		setMergedLabelsAndAnnotations(ep, desired)
		return nil
	})
	return err
}

func (r *gatewayReconciler) ensureEnvoyConfig(ctx context.Context, desired *ciliumv2.CiliumEnvoyConfig) error {
	cec := desired.DeepCopy()
	_, err := controllerutil.CreateOrPatch(ctx, r.Client, cec, func() error {
		cec.Spec = desired.Spec
		setMergedLabelsAndAnnotations(cec, desired)
		return nil
	})
	return err
}

func (r *gatewayReconciler) setAddressStatus(ctx context.Context, gw *gwapiv1b1.Gateway) error {
	svcList := &corev1.ServiceList{}
	if err := r.Client.List(ctx, svcList, client.MatchingLabels{
		owningGatewayLabel: gw.GetName(),
	}); err != nil {
		return err
	}

	if len(svcList.Items) == 0 {
		return fmt.Errorf("no service found")
	}

	svc := svcList.Items[0]
	if len(svc.Status.LoadBalancer.Ingress) == 0 {
		return fmt.Errorf("load balancer status is not ready")
	}

	var addresses []gwapiv1b1.GatewayAddress
	for _, s := range svc.Status.LoadBalancer.Ingress {
		if len(s.IP) != 0 {
			addresses = append(addresses, gwapiv1b1.GatewayAddress{
				Type:  GatewayAddressTypePtr(gwapiv1b1.IPAddressType),
				Value: s.IP,
			})
		}
		if len(s.Hostname) != 0 {
			addresses = append(addresses, gwapiv1b1.GatewayAddress{
				Type:  GatewayAddressTypePtr(gwapiv1b1.HostnameAddressType),
				Value: s.Hostname,
			})
		}
	}

	gw.Status.Addresses = addresses
	return nil
}
