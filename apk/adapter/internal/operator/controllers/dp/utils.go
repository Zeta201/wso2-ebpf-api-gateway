package dp

import (
	"context"
	"errors"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	"github.com/wso2/apk/adapter/internal/operator/utils"
	dpv1alpha2 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"
	"k8s.io/apimachinery/pkg/types"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func (apiReconciler *APIReconciler) getServicesForAPI(ctx context.Context, apiState *synchronizer.APIState, httpRoutes map[string]*gwapiv1b1.HTTPRoute, api dpv1alpha2.API) ([]ServiceState, error) {

	for _, httpRoute := range httpRoutes {
		for _, rule := range httpRoute.Spec.Rules {
			for _, backendRef := range rule.BackendRefs {
				backendNamespacedName := types.NamespacedName{
					Name:      string(backendRef.Name),
					Namespace: utils.GetNamespace(backendRef.Namespace, httpRoute.Namespace),
				}
				resolvedBackend := utils.GetResolvedBackend(ctx, apiReconciler.client, backendNamespacedName, &api)
				if len(resolvedBackend.Backend.Spec.Services) > 0 {

					services := translateService(resolvedBackend.Services)
					return services, nil
				}
			}
		}
	}
	return []ServiceState{}, errors.New("no services found for the api")
}
