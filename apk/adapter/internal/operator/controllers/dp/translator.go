package dp

import (
	"fmt"

	envoy_config_route_v3 "github.com/cilium/proxy/go/envoy/config/route/v3"
	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	dpv1alpha1 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	"google.golang.org/protobuf/types/known/wrapperspb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	wildCard  = "*"
	pathMatch = "/"
)

type Service struct {
	Name      string
	Namespace string
	Port      uint32
}

func translateService(services []dpv1alpha1.Service) []Service {
	tranlatedServices := []Service{}
	for _, s := range services {
		tranlatedServices = append(tranlatedServices, Service{
			Name:      splitHostNamspace(s.Host)[0],
			Port:      s.Port,
			Namespace: splitHostNamspace(s.Host)[1],
		})
	}
	return tranlatedServices
}
func generateJWTEnvoyConfig(services []Service, apiState *synchronizer.APIState) *ciliumv2.CiliumEnvoyConfig {
	cec := &ciliumv2.CiliumEnvoyConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiState.APIDefinition.Name + " zz-generated-envoy-config",
			Namespace: apiState.APIDefinition.Namespace,
			Labels: map[string]string{
				"wso2-auto-generated": "true",
			},
		},
	}

	cec.Spec.Services = getCiliumServices(services, apiState.APIDefinition.Namespace)
	cec.Spec.Resources = getCiliumResources(services)
	return cec
}

func getCiliumServices(services []Service, namespace string) []*ciliumv2.ServiceListener {

	ciliumServices := []*ciliumv2.ServiceListener{}

	for _, service := range services {
		ciliumServices = append(ciliumServices, &ciliumv2.ServiceListener{
			Name:      service.Name,
			Namespace: service.Namespace,
		})
	}
	return ciliumServices
}

func getCiliumResources(services []Service) []ciliumv2.XDSResource {
	var res []ciliumv2.XDSResource
	res = append(res, getHTTPRouteListener()...)
	res = append(res, getEnvoyHTTPRouteConfiguration(services)...)
	return res
}

func getHTTPRouteListener() []ciliumv2.XDSResource {
	mutatorFuns := []ListenerMutator{}
	l, _ := NewHTTPListenerWithDefaults("listener", mutatorFuns...)
	return []ciliumv2.XDSResource{l}
}

func getEnvoyHTTPRouteConfiguration(services []Service) []ciliumv2.XDSResource {
	var res []ciliumv2.XDSResource

	domains := []string{wildCard}
	weightedClusters := make([]*envoy_config_route_v3.WeightedCluster_ClusterWeight, 0, len(services))

	for _, svc := range services {
		var weight int32 = 1
		weightedClusters = append(weightedClusters, &envoy_config_route_v3.WeightedCluster_ClusterWeight{
			Name:   getClusterName(svc.Namespace, svc.Name),
			Weight: wrapperspb.UInt32(uint32(weight)),
		})
	}
	routeAction := &envoy_config_route_v3.Route_Route{
		Route: &envoy_config_route_v3.RouteAction{
			ClusterSpecifier: &envoy_config_route_v3.RouteAction_WeightedClusters{
				WeightedClusters: &envoy_config_route_v3.WeightedCluster{
					Clusters: weightedClusters,
				},
			},
		},
	}

	routes := []*envoy_config_route_v3.Route{
		{
			Match: &envoy_config_route_v3.RouteMatch{
				PathSpecifier: &envoy_config_route_v3.RouteMatch_Prefix{
					Prefix: pathMatch,
				},
			},
			Action: routeAction,
		},
	}

	virtualHosts := []*envoy_config_route_v3.VirtualHost{
		{
			Name:    domains[0],
			Domains: domains,
			Routes:  routes,
		},
	}
	// the route name should match the value in http connection manager
	// otherwise the request will be dropped by envoy
	routeName := fmt.Sprintf("listener-insecure")
	rc, _ := NewRouteConfiguration(routeName, virtualHosts)
	res = append(res, rc)
	return res
}

// func getClusters(services []Service) {
// 	envoyClusters := map[string]ciliumv2.XDSResource{}

// }
