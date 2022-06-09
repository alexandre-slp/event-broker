package endpoints

import (
	"context"

	"github.com/alexandre-slp/event-broker/app/api/v1/gRPC"
	"github.com/alexandre-slp/event-broker/app/api/v1/serializer"
)

// HealthCheckServer is used to implement healthcheck service
type HealthCheckServer struct {
	gRPC.UnimplementedHealthcheckServer
}

// GetHealthCheck implements healthcheck
func (s *HealthCheckServer) GetHealthCheck(ctx context.Context, in *serializer.HealthCheckRequest) (*serializer.HealthCheckResponse, error) {
	return &serializer.HealthCheckResponse{Status: "ok"}, nil
}
