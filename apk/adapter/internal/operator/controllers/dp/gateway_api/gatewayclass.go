package gateway_api

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

type gatewayClassReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func newGatewayClassReconciler(mgr ctrl.Manager) *gatewayClassReconciler {
	return &gatewayClassReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}
}

func (r *gatewayClassReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).For(&gwapiv1b1.GatewayClass{}, builder.WithPredicates(predicate.NewPredicateFuncs(matchesControllerName(controllerName)))).Complete(r)

}

func matchesControllerName(controllerName string) func(object client.Object) bool {
	return func(object client.Object) bool {
		gwc, ok := object.(*gwapiv1b1.GatewayClass)
		if !ok {
			return false
		}
		return string(gwc.Spec.ControllerName) == controllerName
	}
}
