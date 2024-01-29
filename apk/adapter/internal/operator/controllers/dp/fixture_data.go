package dp

import (
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
	"github.com/wso2/apk/adapter/internal/operator/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

var backend = &v1alpha1.ResolvedBackend{
	Backend: v1alpha1.Backend{

		TypeMeta: metav1.TypeMeta{
			Kind:       "Backend",
			APIVersion: "dp.wso2.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "details-backend",
			Namespace: "default",
		},
		Spec: v1alpha1.BackendSpec{
			Services: []v1alpha1.Service{
				{
					Host: "details.default",
					Port: uint32(9090),
				},
			},
		},
	},
}

var httpRoute = &gwapiv1b1.HTTPRoute{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "details-route",
		Namespace: "default",
	},
	Spec: gwapiv1b1.HTTPRouteSpec{
		CommonRouteSpec: gwapiv1b1.CommonRouteSpec{
			ParentRefs: []gwapiv1b1.ParentReference{
				{
					Name:        "valid-gateway",
					SectionName: nil,
				},
			},
		},
		Rules: []gwapiv1b1.HTTPRouteRule{
			{
				BackendRefs: []gwapiv1b1.HTTPBackendRef{
					{
						BackendRef: gwapiv1b1.BackendRef{
							BackendObjectReference: gwapiv1b1.BackendObjectReference{
								Name:  "details-backend",
								Port:  addressOf[gwapiv1b1.PortNumber](9090),
								Group: utils.GroupPtr("dp.wso2.com"),
								Kind:  utils.KindPtr("Backend"),
							},
						},
					},
				},
				Matches: []gwapiv1b1.HTTPRouteMatch{
					{
						Path: &gwapiv1b1.HTTPPathMatch{
							Type:  utils.PathMatchTypePtr(gwapiv1b1.PathMatchPathPrefix),
							Value: utils.StringPtr("/details"),
						},
					},
				},
			},
		},
	},
}

var basicAPIState = synchronizer.APIState{
	APIDefinition: &v1alpha2.API{
		TypeMeta: metav1.TypeMeta{
			Kind:       "API",
			APIVersion: "dp.wso2.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "details-api",
			Namespace: "default",
		},
		Spec: v1alpha2.APISpec{
			APIName:    "details-api",
			APIVersion: "dp.wso2.com/v1alpha1",
			BasePath:   "/details-api",
		},
	},
	ProdHTTPRoute: &synchronizer.HTTPRouteState{
		HTTPRoutePartitions: map[string]*gwapiv1b1.HTTPRoute{
			"default/details-route": httpRoute,
		},
	},
}

var services = []ServiceState{{
	Name:      "details",
	Port:      uint32(9090),
	Namespace: "default",
}}

var translatedHTTPRoute = &gwapiv1b1.HTTPRoute{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "details-route-zz-generated",
		Namespace: "default",
	},
	Spec: gwapiv1b1.HTTPRouteSpec{
		CommonRouteSpec: gwapiv1b1.CommonRouteSpec{
			ParentRefs: []gwapiv1b1.ParentReference{
				{
					Name: "valid-gateway",
				},
			},
		},
		Rules: []gwapiv1b1.HTTPRouteRule{
			{
				BackendRefs: []gwapiv1b1.HTTPBackendRef{
					{
						BackendRef: gwapiv1b1.BackendRef{
							BackendObjectReference: gwapiv1b1.BackendObjectReference{
								Name:  "details",
								Port:  addressOf[gwapiv1b1.PortNumber](9090),
								Group: utils.GroupPtr("gateway.networking.k8s.io"),
								Kind:  utils.KindPtr("Service"),
							},
						},
					},
				},
				Matches: []gwapiv1b1.HTTPRouteMatch{
					{
						Path: &gwapiv1b1.HTTPPathMatch{
							Type:  utils.PathMatchTypePtr(gwapiv1b1.PathMatchPathPrefix),
							Value: utils.StringPtr("/details-api/details"),
						},
					},
				},
			},
		},
	},
}

var apiDefinition = &v1alpha1.API{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "http-bin-api",
		Namespace: "default",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "dp.wso2.com/v1alpha1",
		Kind:       "API",
	},
	Spec: v1alpha1.APISpec{
		APIName:          "HTTP Bin API",
		APIType:          "REST",
		APIVersion:       "1.0.0",
		BasePath:         "/http-bin-api/1.0.0",
		DefinitionPath:   "/api-definition",
		IsDefaultVersion: false,
		Organization:     "apk-system",
		Production: []v1alpha1.EnvConfig{
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
				Name:       "details-api",
			},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "cilium.io/v2",
		Kind:       "CiliumEnvoyConfig",
	},
	Spec: ciliumv2.CiliumEnvoyConfigSpec{
		Services: []*ciliumv2.ServiceListener{
			{
				Name:      "details",
				Namespace: "default",
			},
		},
		Resources: []ciliumv2.XDSResource{},
	},
}
