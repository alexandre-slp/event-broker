package infra

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"os"
	"path"
	"strings"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05-0700"
)

////UnaryZerologOption todo
//func UnaryZerologOption() grpc.ServerOption {
//	return grpc.UnaryInterceptor(UnaryZerologInterceptor())
//}

//UnaryZerologInterceptor todo
func UnaryZerologInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// Before processing
		var requestIdSlice []string
		requestId := uuid.New().String()
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			requestIdSlice = md.Get("x-request-id")
		}

		if len(requestIdSlice) > 0 {
			requestId = requestIdSlice[0]
		}

		contextAppLogger := setupAppLogger().
			With().
			Str("requestId", requestId).
			Logger()

		ctxWithLogger := context.WithValue(ctx, "logger", contextAppLogger)
		startRequestTime := time.Now()

		// Process
		resp, err := handler(ctxWithLogger, req)

		// After process
		endRequestTime := time.Now()
		p, ok := peer.FromContext(ctx)
		if !ok {
			// todo error handling
		}
		requestDuration := fmt.Sprintf("%vms", endRequestTime.Sub(startRequestTime).Milliseconds())
		contextRequestLogger := setupRequestLogger().
			With().
			Str("requestId", requestId).
			Str("method", path.Base(info.FullMethod)).
			Str("requestDuration", requestDuration).
			Str("statusCode", status.Code(err).String()).
			Str("ip", p.Addr.String()).
			Logger()

		contextRequestLogger.Info().Msg("")

		return resp, err
	}
}

func setupRequestLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: timeFormat,
	}
	output.FormatFieldName = func(i interface{}) string {
		return ""
	}

	output.PartsExclude = []string{
		zerolog.LevelFieldName,
		zerolog.MessageFieldName,
	}

	return zerolog.New(output).
		With().
		Timestamp().
		Logger()
}

func setupAppLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: timeFormat,
	}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("msg=%v", i)
	}

	return zerolog.New(output).
		With().
		Timestamp().
		Logger()
}
