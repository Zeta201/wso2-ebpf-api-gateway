package dp

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha1"
)

func Test_tranlate_to_service_Test(t *testing.T) {
	type args struct {
		services []v1alpha1.Service
	}

	tests := []struct {
		name    string
		args    args
		want    []ServiceState
		wantErr bool
	}{
		{
			name: "dv1alpha1 service to service",
			args: args{
				services: []v1alpha1.Service{
					{
						Host: "echo-service.default",
						Port: uint32(8080),
					},
				},
			},
			want: []ServiceState{
				{
					Name:      "echo-service",
					Port:      uint32(8080),
					Namespace: "default",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translatedSvc := translateService(tt.args.services)
			require.Equal(t, tt.want, translatedSvc, "Translation failed")
		})
	}
}
