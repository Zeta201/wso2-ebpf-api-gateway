package gateway_api

import (
	"context"

	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	kindGateway   = "Gateway"
	kindHTTPRoute = "HTTPRoute"
	kindTLSRoute  = "TLSRoute"
	kindUDPRoute  = "UDPRoute"
	kindTCPRoute  = "TCPRoute"
	kindService   = "Service"
	kindSecret    = "Secret"
)

func GatewayAddressTypePtr(addr gwapiv1b1.AddressType) *gwapiv1b1.AddressType {
	return &addr
}

func GroupPtr(name string) *gwapiv1b1.Group {
	group := gwapiv1b1.Group(name)
	return &group
}

func KindPtr(name string) *gwapiv1b1.Kind {
	kind := gwapiv1b1.Kind(name)
	return &kind
}

func ObjectNamePtr(name string) *gwapiv1b1.ObjectName {
	objectName := gwapiv1b1.ObjectName(name)
	return &objectName
}

func groupDerefOr(group *gwapiv1b1.Group, defaultGroup string) string {
	if group != nil && *group != "" {
		return string(*group)
	}
	return defaultGroup
}

// isAllowed returns true if the provided Route is allowed to attach to given gateway
func isAllowed(ctx context.Context, c client.Client, gw *gwapiv1b1.Gateway, route metav1.Object) bool {
	for _, listener := range gw.Spec.Listeners {
		// all routes in the same namespace are allowed for this listener
		if listener.AllowedRoutes == nil || listener.AllowedRoutes.Namespaces == nil {
			return route.GetNamespace() == gw.GetNamespace()
		}

		// check if route is kind-allowed
		if !isKindAllowed(listener, route) {
			continue
		}

		// check if route is namespace-allowed
		switch *listener.AllowedRoutes.Namespaces.From {
		case gwapiv1b1.NamespacesFromAll:
			return true
		case gwapiv1b1.NamespacesFromSame:
			if route.GetNamespace() == gw.GetNamespace() {
				return true
			}
		case gwapiv1b1.NamespacesFromSelector:
			nsList := &corev1.NamespaceList{}
			selector, _ := metav1.LabelSelectorAsSelector(listener.AllowedRoutes.Namespaces.Selector)
			if err := c.List(ctx, nsList, client.MatchingLabelsSelector{Selector: selector}); err != nil {
				log.WithError(err).Error("Unable to list namespaces")
				return false
			}

			for _, ns := range nsList.Items {
				if ns.Name == route.GetNamespace() {
					return true
				}
			}
		}
	}
	return false
}

func isKindAllowed(listener gwapiv1b1.Listener, route metav1.Object) bool {
	if listener.AllowedRoutes.Kinds == nil {
		return true
	}

	routeKind := getGatewayKindForObject(route)

	for _, kind := range listener.AllowedRoutes.Kinds {
		if (kind.Group == nil || string(*kind.Group) == gwapiv1b1.GroupName) &&
			kind.Kind == kindHTTPRoute && routeKind == kindHTTPRoute {
			return true
		} else if (kind.Group == nil || string(*kind.Group) == gwapiv1a2.GroupName) &&
			kind.Kind == kindTLSRoute && routeKind == kindTLSRoute {
			return true
		}
	}
	return false
}

func computeHosts[T ~string](gw *gwapiv1b1.Gateway, hostnames []T) []string {
	hosts := make([]string, 0, len(hostnames))
	for _, listener := range gw.Spec.Listeners {
		hosts = append(hosts, computeHostsForListener(&listener, hostnames)...)
	}

	return hosts
}

func computeHostsForListener[T ~string](listener *gwapiv1b1.Listener, hostnames []T) []string {
	return model.ComputeHosts(toStringSlice(hostnames), (*string)(listener.Hostname))
}

func toStringSlice[T ~string](s []T) []string {
	res := make([]string, 0, len(s))
	for _, h := range s {
		res = append(res, string(h))
	}
	return res
}

func getSupportedGroupKind(protocol gwapiv1b1.ProtocolType) (*gwapiv1b1.Group, gwapiv1b1.Kind) {
	switch protocol {
	case gwapiv1b1.TLSProtocolType:
		return GroupPtr(gwapiv1a2.GroupName), kindTLSRoute
	case gwapiv1b1.HTTPSProtocolType:
		return GroupPtr(gwapiv1b1.GroupName), kindHTTPRoute
	case gwapiv1b1.HTTPProtocolType:
		return GroupPtr(gwapiv1b1.GroupName), kindHTTPRoute
	case gwapiv1b1.TCPProtocolType:
		return GroupPtr(gwapiv1a2.GroupName), kindTCPRoute
	case gwapiv1b1.UDPProtocolType:
		return GroupPtr(gwapiv1a2.GroupName), kindUDPRoute
	default:
		return GroupPtr("Unknown"), "Unknown"
	}
}
func getGatewayKindForObject(obj metav1.Object) gwapiv1b1.Kind {
	switch obj.(type) {
	case *gwapiv1b1.HTTPRoute:
		return kindHTTPRoute
	case *gwapiv1a2.TLSRoute:
		return kindTLSRoute
	case *gwapiv1a2.UDPRoute:
		return kindUDPRoute
	case *gwapiv1a2.TCPRoute:
		return kindTCPRoute
	default:
		return "Unknown"
	}
}

func mergeMap(left, right map[string]string) map[string]string {
	if left == nil {
		return right
	} else {
		for key, value := range right {
			left[key] = value
		}
	}
	return left
}

func setMergedLabelsAndAnnotations(temp, desired client.Object) {
	temp.SetAnnotations(mergeMap(temp.GetAnnotations(), desired.GetAnnotations()))
	temp.SetLabels(mergeMap(temp.GetLabels(), desired.GetLabels()))
}
