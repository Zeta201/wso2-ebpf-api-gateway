package helpers

import (
	corev1 "k8s.io/api/core/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// IsBackendReferenceAllowed returns true if the backend reference is allowed by the reference grant.
func IsBackendReferenceAllowed(originatingNamespace string, be gwapiv1b1.BackendRef, gvk schema.GroupVersionKind, grants []gwapiv1b1.ReferenceGrant) bool {
	return isReferenceAllowed(originatingNamespace, string(be.Name), be.Namespace, gvk, corev1.SchemeGroupVersion.WithKind("Service"), grants)
}

// IsSecretReferenceAllowed returns true if the secret reference is allowed by the reference grant.
func IsSecretReferenceAllowed(originatingNamespace string, sr gwapiv1b1.SecretObjectReference, gvk schema.GroupVersionKind, grants []gwapiv1b1.ReferenceGrant) bool {
	return isReferenceAllowed(originatingNamespace, string(sr.Name), sr.Namespace, gvk, corev1.SchemeGroupVersion.WithKind("Secret"), grants)
}

func isReferenceAllowed(originatingNamespace, name string, namespace *gwapiv1b1.Namespace, fromGVK, toGVK schema.GroupVersionKind, grants []gwapiv1b1.ReferenceGrant) bool {
	ns := NamespaceDerefOr(namespace, originatingNamespace)
	if originatingNamespace == ns {
		return true // same namespace is always allowed
	}

	for _, g := range grants {
		if g.Namespace != ns {
			continue
		}
		for _, from := range g.Spec.From {
			if (from.Group == gwapiv1b1.Group(fromGVK.Group) && from.Kind == gwapiv1b1.Kind(fromGVK.Kind)) &&
				(string)(from.Namespace) == originatingNamespace {
				for _, to := range g.Spec.To {
					if to.Group == gwapiv1b1.Group(toGVK.Group) && to.Kind == gwapiv1b1.Kind(toGVK.Kind) &&
						(to.Name == nil || string(*to.Name) == name) {
						return true
					}
				}
			}
		}
	}
	return false
}
