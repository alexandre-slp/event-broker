package main

import (
	"github.com/alexandre-slp/event-broker/app"
	gRPC2 "github.com/alexandre-slp/event-broker/app/api/v1/gRPC"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatalf("error while loading cfg. err: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opt := grpc_recovery.WithRecoveryHandlerContext(app.CustomPanicHandler)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			infra.UnaryZerologInterceptor(cfg),
			grpc_recovery.UnaryServerInterceptor(opt),
		)),
	)

	gRPC2.RegisterHealthcheckServer(s, &domain.HealthCheckServer{})
	gRPC2.RegisterEventServer(s, &domain.EventServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
