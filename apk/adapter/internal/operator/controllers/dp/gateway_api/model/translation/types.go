package translation

import (
	v1 "k8s.io/api/core/v1"

	"github.com/wso2/apk/adapter/internal/operator/controllers/dp/gateway_api/model"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
)

// Translator is the interface to take the model and generate required CiliumEnvoyConfig,
// LoadBalancer Service, Endpoint, etc.
//
// Different use cases (e.g. Ingress, Gateway API) can provide its own generation logic.
type Translator interface {
	Translate(model *model.Model) (*ciliumv2.CiliumEnvoyConfig, *v1.Service, *v1.Endpoints, error)
}
