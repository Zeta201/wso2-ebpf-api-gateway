package translation

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cilium/cilium/pkg/slices"
	envoy_config_cluster_v3 "github.com/cilium/proxy/go/envoy/config/cluster/v3"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	secureHost                    = "secure"
	insecureHost                  = "insecure"
	UseOriginalSourceAddressLabel = "cilium.io/use-original-source-address"
)
const (
	httpProtocolOptionsType = "envoy.extensions.upstreams.http.v3.HttpProtocolOptions"
)

type ClusterMutator func(*envoy_config_cluster_v3.Cluster) *envoy_config_cluster_v3.Cluster

var _ Translator = (*defaultTranslator)(nil)

type defaultTranslator struct {
	name                string
	namespace           string
	secretsNamespace    string
	enforceHTTPs        bool
	useProxyProtocol    bool
	hostNameSuffixMatch bool
	idleTimeoutSeconds  int
}

func NewTranslator(name, namespace, secretsNamespace string, enforceHTTPs bool, useProxyProtocol bool, hostNameSuffixMatch bool, idleTimeoutSeconds int) Translator {
	return &defaultTranslator{
		name:                name,
		namespace:           namespace,
		secretsNamespace:    secretsNamespace,
		enforceHTTPs:        enforceHTTPs,
		useProxyProtocol:    useProxyProtocol,
		hostNameSuffixMatch: hostNameSuffixMatch,
		idleTimeoutSeconds:  idleTimeoutSeconds,
	}
}

func (i *defaultTranslator) Translate(apiState *synchronizer.APIState) (*ciliumv2.CiliumEnvoyConfig, *corev1.Service, *corev1.Endpoints, error) {

	cec := &ciliumv2.CiliumEnvoyConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      i.name,
			Namespace: i.namespace,
			// Labels: map[string]string{
			// 	UseOriginalSourceAddressLabel: "false",
			// },
		},
	}
	// cec.Spec.BackendServices = i.getBackendServices(apiState)
	cec.Spec.Services = i.getServices(apiState)
	cec.Spec.Resources = i.getResources(apiState)

	return cec, nil, nil, nil
}

func (i *defaultTranslator) getServices(apiState *synchronizer.APIState) []*ciliumv2.ServiceListener {

	var services []*ciliumv2.ServiceListener

	for name, backend := range mergeBackendMaps(apiState) {
		ns := strings.Split(name, "/")[0]

		for _, svc := range backend.Services {
			host := strings.Split(svc.Host, ".")[0]
			services = append(services, &ciliumv2.ServiceListener{

				Name:      host,
				Namespace: ns,
			})
		}

	}
	return services
}

func getNamespaceNamePortsMapForHTTP(m *synchronizer.APIState) map[string]map[string][]string {

	namespaceNamePortMap := map[string]map[string][]string{}

	for name, backend := range mergeBackendMaps(m) {
		ns := strings.Split(name, "/")[0]

		for _, svc := range backend.Services {

			host := strings.Split(svc.Host, ".")[0]
			namePortMap, exist := namespaceNamePortMap[ns]

			if exist {
				namePortMap[host] = slices.SortedUnique(append(namePortMap[host], fmt.Sprint(svc.Port)))
			} else {
				namePortMap = map[string][]string{
					host: {fmt.Sprint(svc.Port)},
				}
			}
			namespaceNamePortMap[ns] = namePortMap
		}

	}

	return namespaceNamePortMap
}
func (i *defaultTranslator) getResources(m *synchronizer.APIState) []ciliumv2.XDSResource {
	var res []ciliumv2.XDSResource
	res = append(res, i.getClusters(m)...)
	return res
}




func (i *defaultTranslator) getClusters(m *synchronizer.APIState) []ciliumv2.XDSResource {
	envoyClusters := map[string]ciliumv2.XDSResource{}
	rateLimitClusters := map[string]ciliumv2.XDSResource{}

	var sortedClusterNames []string
	// EDS Cluster for services
	for ns, v := range getNamespaceNamePortsMapForHTTP(m) {
		for name, ports := range v {
			for _, port := range ports {
				clusterName := getClusterName(ns, name, port)
				clusterServiceName := getClusterServiceName(ns, name, port)
				sortedClusterNames = append(sortedClusterNames, clusterName)
				mutators := []ClusterMutator{
					WithConnectionTimeout(5),
					WithIdleTimeout(i.idleTimeoutSeconds),
					WithClusterLbPolicy(int32(envoy_config_cluster_v3.Cluster_ROUND_ROBIN)),
					WithOutlierDetection(true),
				}
				envoyClusters[clusterName], _ = NewHTTPCluster(clusterName, clusterServiceName, mutators...)
			}
		}
	}

	// rate limit clusters

	for clusterName := range m.RateLimitPolicies {
		sortedClusterNames = append(sortedClusterNames, clusterName)

		mutators := []ClusterMutator{
			WithConnectionTimeout(5),
			WithIdleTimeout(i.idleTimeoutSeconds),
			WithClusterLbPolicy(int32(envoy_config_cluster_v3.Cluster_ROUND_ROBIN)),
			WithOutlierDetection(true),
		}

		rateLimitClusters[clusterName], _ = NewRateLimitCluster(clusterName, mutators...)
	}
	sort.Strings(sortedClusterNames)

	res := make([]ciliumv2.XDSResource, len(sortedClusterNames))
	for i, name := range sortedClusterNames {

		if _, ok := envoyClusters[name]; !ok {
			res[i] = rateLimitClusters[name]
		} else {

			res[i] = envoyClusters[name]
		}
	}

	return res
}
