package translation

import (
	v1 "k8s.io/api/core/v1"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
)

const ClusterTypeURL = "type.googleapis.com/envoy.config.cluster.v3.Cluster"

// Translator is the interface to take the model and generate required CiliumEnvoyConfig,
// LoadBalancer Service, Endpoint, etc.
//
// Different use cases (e.g. Ingress, Gateway API) can provide its own generation logic.
type Translator interface {
	Translate(gatewayState *synchronizer.APIState) (*ciliumv2.CiliumEnvoyConfig, *v1.Service, *v1.Endpoints, error)
}
