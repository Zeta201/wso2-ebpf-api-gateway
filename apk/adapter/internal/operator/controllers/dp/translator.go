package dp

import (
	"log"
	"sort"

	envoy_config_cluster_v3 "github.com/cilium/proxy/go/envoy/config/cluster/v3"

	envoy_config_route_v3 "github.com/cilium/proxy/go/envoy/config/route/v3"
	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	dpv1alpha1 "github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	"google.golang.org/protobuf/types/known/wrapperspb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	wildCard           = "*"
	pathMatch          = "/"
	idleTimeoutSeconds = 0
)

type ServiceState struct {
	Name      string
	Namespace string
	Port      uint32
}

type AuthProvider struct {
	ProviderName       string
	Issuer             string
	Uri                string
	DefaultClusterName string
	ClusterName        string
}

var authProvider = AuthProvider{
	ProviderName:       "provider_funnel",
	Issuer:             " https://auth.funnel-labs.io/auth/realms/funnel",
	Uri:                "https://auth.funnel-labs.io/auth/realms/funnel/protocol/openid-connect/certs",
	DefaultClusterName: "funnel_auth_cluster",
}

func translateService(services []dpv1alpha1.Service) []ServiceState {
	tranlatedServices := []ServiceState{}
	for _, s := range services {
		tranlatedServices = append(tranlatedServices, ServiceState{
			Name:      splitHostNamspace(s.Host)[0],
			Port:      s.Port,
			Namespace: splitHostNamspace(s.Host)[1],
		})
	}
	return tranlatedServices
}
func generateJWTEnvoyConfig(services []ServiceState, apiState *synchronizer.APIState) *ciliumv2.CiliumEnvoyConfig {

	cec := &ciliumv2.CiliumEnvoyConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      apiState.APIDefinition.Name + "-zz-generated-envoy-config",
			Namespace: apiState.APIDefinition.Namespace,
			Labels: map[string]string{
				"wso2-auto-generated": "true",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "cilium.io/v2",
			Kind:       "CiliumEnvoyConfig",
		},
	}

	authProvider.ClusterName = cec.Namespace + "/" + cec.Name + "/" + authProvider.DefaultClusterName

	cec.Spec.Services = getCiliumServices(services, apiState.APIDefinition.Namespace)
	cec.Spec.Resources = getCiliumResources(services)

	// Set the owner reference to the CEC object.
	cec.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: apiState.APIDefinition.APIVersion,
			Kind:       apiState.APIDefinition.Kind,
			Name:       apiState.APIDefinition.Name,
			UID:        types.UID(apiState.APIDefinition.UID),
		},
	}
	log.Printf("Generated CEC %v", *cec)
	return cec
}

func getCiliumServices(services []ServiceState, namespace string) []*ciliumv2.ServiceListener {

	ciliumServices := []*ciliumv2.ServiceListener{}

	for _, service := range services {
		ciliumServices = append(ciliumServices, &ciliumv2.ServiceListener{
			Name:      service.Name,
			Namespace: service.Namespace,
		})
	}
	return ciliumServices
}

func getCiliumResources(services []ServiceState) []ciliumv2.XDSResource {
	var res []ciliumv2.XDSResource
	res = append(res, getHTTPRouteListener()...)
	res = append(res, getEnvoyHTTPRouteConfiguration(services)...)
	res = append(res, getClusters(services)...)

	return res
}

func getHTTPRouteListener() []ciliumv2.XDSResource {
	mutatorFuns := []ListenerMutator{}

	l, _ := NewHTTPListenerWithDefaults("listener", &authProvider, mutatorFuns...)
	return []ciliumv2.XDSResource{l}
}

func getEnvoyHTTPRouteConfiguration(services []ServiceState) []ciliumv2.XDSResource {
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
	routeName := "listener-insecure"
	rc, _ := NewRouteConfiguration(routeName, virtualHosts)
	res = append(res, rc)
	return res
}

func getClusters(services []ServiceState) []ciliumv2.XDSResource {
	envoyClusters := map[string]ciliumv2.XDSResource{}
	var sortedClusterNames []string

	mutators := []ClusterMutator{
		WithConnectionTimeout(5),
		WithIdleTimeout(idleTimeoutSeconds),
		WithClusterLbPolicy(int32(envoy_config_cluster_v3.Cluster_ROUND_ROBIN)),
		WithOutlierDetection(true),
	}
	for _, svc := range services {
		clusterName := getClusterName(svc.Namespace, svc.Name)
		sortedClusterNames = append(sortedClusterNames, clusterName)

		envoyClusters[clusterName], _ = NewHTTPCluster(clusterName, mutators...)

	}
	envoyClusters[authProvider.DefaultClusterName], _ = NewJWKSCluster(mutators...)
	sortedClusterNames = append(sortedClusterNames, authProvider.DefaultClusterName)

	sort.Strings(sortedClusterNames)
	res := make([]ciliumv2.XDSResource, len(sortedClusterNames))
	for i, name := range sortedClusterNames {
		res[i] = envoyClusters[name]
	}

	return res
}
