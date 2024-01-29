package dp

import (
	"testing"

	"github.com/stretchr/testify/require"
	ciliumv2 "github.com/wso2/apk/common-go-libs/apis/cilium.io/v2"

	"github.com/wso2/apk/adapter/internal/operator/synchronizer"
)

func Test_translator_Test(t *testing.T) {
	type args struct {
		apiState *synchronizer.APIState
		services []ServiceState
	}

	tests := []struct {
		name    string
		args    args
		want    *ciliumv2.CiliumEnvoyConfig
		wantErr bool
	}{
		{
			name: "Basic API State",
			args: args{
				apiState: &basicAPIState,
				services: services,
			},
			want: &ciliumv2.CiliumEnvoyConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cec := generateJWTEnvoyConfig(tt.args.services, tt.args.apiState)
			// require.Equal(t,tt.wantErr,err!=nil,"")
			require.NotEqual(t, tt.want, cec, "Translation failed")
		})

	}
}
