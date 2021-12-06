package infra

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"os"
	"strings"
	"time"
)

var (
	RequestLogger = SetupRequestLogger()
	AppLogger = SetupAppLogger()
)

const (
	timeFormat = "2006-01-02T15:04:05-0700"
)

func SetupRequestLogger() zerolog.Logger {
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

func SetupAppLogger() zerolog.Logger {
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

func LoggerMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		requestId := string(ctx.Request.Header.Peek("x-request-id"))

		if requestId == "" {
			requestId = uuid.New().String()
		}

		contextAppLogger := AppLogger.
			With().
			Str("requestId", requestId).
			Logger()

		startRequestTime := time.Now()
		ctx.SetUserValue("logger", contextAppLogger)
		// Before request
		next(ctx)
		// After request
		endRequestTime := time.Now()
		requestDuration := fmt.Sprintf("%vms", endRequestTime.Sub(startRequestTime).Milliseconds())
		contextRequestLogger := RequestLogger.
			With().
			Str("requestId", requestId).
			Str("path", string(ctx.Path())).
			Str("method", string(ctx.Method())).
			Str("requestDuration", requestDuration).
			Str("statusCode", fmt.Sprintf("%v", ctx.Response.StatusCode())).
			Logger()

		contextRequestLogger.Info().Msg("")
	}
}
