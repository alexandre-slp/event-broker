package main

import (
	"context"
	"github.com/alexandre-slp/event-broker/app"
	"github.com/alexandre-slp/event-broker/app/api/gRPC"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", app.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Define customfunc to handle panic
	customFunc := func(ctx context.Context, p interface{}) (err error) {
		l := ctx.Value("logger").(zerolog.Logger)
		l.Error().Msg("test")
		return status.Errorf(codes.Internal, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opt := grpc_recovery.WithRecoveryHandlerContext(customFunc)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			infra.UnaryZerologInterceptor(),
			grpc_recovery.UnaryServerInterceptor(opt),
		)),
	)

	gRPC.RegisterHealthcheckServer(s, &domain.HealthCheckServer{})
	gRPC.RegisterEventServer(s, &domain.EventServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
