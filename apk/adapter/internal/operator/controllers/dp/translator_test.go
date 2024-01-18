package dp

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
)

func Test_translator_Translate(t *testing.T) {
	type args struct {
		m *synchronizer.APIState
	}
	tests := []struct {
		name    string
		args    args
		want    *ciliumv2.CiliumEnvoyConfig
		wantErr bool
	}{
		{
			name: "Basic API",
			args: args{
				m: &synchronizer.APIState{
					APIDefinition: apiDefinition,
					ProdHTTPRoute: &synchronizer.HTTPRouteState{
						BackendMapping: map[string]*v1alpha1.ResolvedBackend{
							"default/http-bin-backend": {
								Services: []v1alpha1.Service{
									{
										Host: "echo-service.default",
										Port: 8080,
									},
								},
							},
						},
					},
					RateLimitPolicies: map[string]v1alpha1.RateLimitPolicy{
						"default/prod-http-bin-ratelimit": {
							TypeMeta: metav1.TypeMeta{
								APIVersion: "dp.wso2.com/v1alpha1",
								Kind:       "RateLimitPolicy",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "prod-http-bin-ratelimit",
								Namespace: "default",
							},
							Spec: v1alpha1.RateLimitPolicySpec{
								Default: &v1alpha1.RateLimitAPIPolicy{
									API: &v1alpha1.APIRateLimitPolicy{
										RequestsPerUnit: 1,
										Unit:            "Minute",
									},
								},
								TargetRef: v1alpha2.PolicyTargetReference{
									Kind:  "API",
									Name:  "http-bin-api",
									Group: "gateway.networking.k8s.io",
								},
							},
						},
					},
				},
			},
			want: basicAPICiliumEnvoyConfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trans := &translator{
				idleTimeoutSeconds: 60,
			}
			cec, err := trans.Translate(tt.args.m)
			log.Printf("%v", cec)
			require.Equal(t, tt.wantErr, err != nil, "Error mismatch")
		})
	}
}
