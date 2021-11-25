package healthcheck_test

import (
	"github.com/alexandre-slp/event-broker/app/api/healthcheck"
	"github.com/alexandre-slp/event-broker/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestInitializeHealthCheckService_InitPaths(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ihcs := healthcheck.NewService()
	mrs := mocks.NewMockRouter(ctrl)

	mrs.
		EXPECT().
		GET("/healthcheck", healthcheck.HealthCheck)


	ihcs.InitPaths(mrs)
}

func TestNewService(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want *healthcheck.InitializeHealthCheckService
	}{
		{
			name: "Test new service",
			want: &healthcheck.InitializeHealthCheckService{},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := healthcheck.NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
