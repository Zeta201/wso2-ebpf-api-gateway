package helpers

import (
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

func NamespaceDerefOr(namespace *gatewayv1.Namespace, defaultNamespace string) string {
	if namespace != nil && *namespace != "" {
		return string(*namespace)
	}
	return defaultNamespace
}
