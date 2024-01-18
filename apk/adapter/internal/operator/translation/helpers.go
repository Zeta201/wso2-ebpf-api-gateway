package translation

import (
	"fmt"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func getClusterName(ns, name, port string) string {
	// the name is having the format of "namespace:name:port"
	// -> slash would prevent ParseResources from rewriting with CEC namespace and name!
	return fmt.Sprintf("%s:%s:%s", ns, name, port)
}

func getClusterServiceName(ns, name, port string) string {
	// the name is having the format of "namespace/name:port"
	return fmt.Sprintf("%s/%s:%s", ns, name, port)
}

func mergeMap[T any](left, right map[string]T) map[string]T {
	if left == nil {
		return right
	} else {
		for key, value := range right {
			left[key] = value
		}
	}
	return left
}

func mergeBackendMaps(apiState *synchronizer.APIState) map[string]*v1alpha1.ResolvedBackend {
	var backendMap = make(map[string]*v1alpha1.ResolvedBackend)

	if apiState.ProdHTTPRoute != nil && apiState.SandHTTPRoute != nil {
		backendMap = mergeMap(apiState.ProdHTTPRoute.BackendMapping, apiState.SandHTTPRoute.BackendMapping)
	} else if apiState.SandHTTPRoute != nil {
		backendMap = apiState.SandHTTPRoute.BackendMapping
	} else {
		backendMap = apiState.ProdHTTPRoute.BackendMapping
	}
	return backendMap

}
func SetMergedLabelsAndAnnotations(temp, desired client.Object) {
	temp.SetAnnotations(mergeMap(temp.GetAnnotations(), desired.GetAnnotations()))
	temp.SetLabels(mergeMap(temp.GetLabels(), desired.GetLabels()))
}
