package healthcheck

import (
	"github.com/alexandre-slp/event-broker/app/api"
)

func NewService() *InitializeHealthCheckService {
	return &InitializeHealthCheckService{}
}

type InitializeHealthCheckService struct{}

func (InitializeHealthCheckService) InitPaths(router api.Router) {
	router.GET("/healthcheck", HealthCheck)
}
