package endpoints

import (
	"context"

	gRPC "github.com/alexandre-slp/event-broker/infra/api/v1/gRPC"
	serializer "github.com/alexandre-slp/event-broker/infra/api/v1/serializer"
)

// HealthCheckServer is used to implement healthcheck service
type HealthCheckServer struct {
	gRPC.UnimplementedHealthcheckServer
}

// GetHealthCheck implements healthcheck
func (s *HealthCheckServer) GetHealthCheck(ctx context.Context, in *serializer.HealthCheckRequest) (*serializer.HealthCheckResponse, error) {
	// TODO: import logic
	return &serializer.HealthCheckResponse{Status: "ok"}, nil
}
