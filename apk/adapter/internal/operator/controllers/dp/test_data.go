package dp

import (
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var apiDefinition = &v1alpha2.API{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "http-bin-api",
		Namespace: "default",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "dp.wso2.com/v1alpha1",
		Kind:       "API",
	},
	Spec: v1alpha2.APISpec{
		APIName:          "HTTP Bin API",
		APIType:          "REST",
		APIVersion:       "1.0.0",
		BasePath:         "/http-bin-api/1.0.0",
		DefinitionPath:   "/api-definition",
		IsDefaultVersion: false,
		Organization:     "apk-system",
		Production: []v1alpha2.EnvConfig{
			{
				HTTPRouteRefs: []string{"prod-http-route-http-bin-api"},
			},
		},
	},
}

var basicAPICiliumEnvoyConfig = &ciliumv2.CiliumEnvoyConfig{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "http-bin-api-envoy-config",
		Namespace: "default",
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: "dp.wso2.com/v1alpha1",
				Kind:       "API",
				Name:       "http-bin-api",
				Controller: addressOf(true),
			},
		},
	},
}
