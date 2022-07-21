package main

import (
	"log"
	"net"

	"github.com/alexandre-slp/event-broker/infra/api/v1/endpoints"
	gRPC "github.com/alexandre-slp/event-broker/infra/api/v1/gRPC"

	"github.com/alexandre-slp/event-broker/infra"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := infra.NewConfig()
	if err != nil {
		log.Fatalf("error while loading cfg. err: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opt := grpcRecovery.WithRecoveryHandlerContext(infra.CustomPanicHandler)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			infra.UnaryZerologInterceptor(cfg),
			grpcRecovery.UnaryServerInterceptor(opt),
		)),
	)

	gRPC.RegisterHealthcheckServer(s, &endpoints.HealthCheckServer{})
	gRPC.RegisterEventServer(s, &endpoints.EventServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
