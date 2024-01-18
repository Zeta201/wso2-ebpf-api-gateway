package dp

import (
	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	"github.com/wso2/apk/adapter/internal/operator/translation"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type translator struct {
	secretsNamespace   string
	idleTimeoutSeconds int
}

func NewTranslator(secretsNamespace string, idleTimeoutSeconds int) *translator {
	return &translator{
		secretsNamespace:   secretsNamespace,
		idleTimeoutSeconds: idleTimeoutSeconds,
	}
}

func (t *translator) Translate(m *synchronizer.APIState) (*ciliumv2.CiliumEnvoyConfig, error) {
	trans := translation.NewTranslator(m.APIDefinition.Name+"-envoy-config", m.APIDefinition.Namespace, t.secretsNamespace, false, false, true, t.idleTimeoutSeconds)
	cec, _, _, err := trans.Translate(m)
	if err != nil {
		return nil, err
	}

	cec.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: m.APIDefinition.APIVersion,
			Kind:       m.APIDefinition.Kind,
			Name:       m.APIDefinition.Name,
			UID:        types.UID(m.APIDefinition.UID),
			Controller: AddressOf(true),
		},
	}
	return cec, nil
}

func (t *translator) TranslateGateway(m *synchronizer.GatewayState) (*ciliumv2.CiliumEnvoyConfig, *corev1.Service, *corev1.Endpoints, error) {
	return nil, nil, nil, nil
}
