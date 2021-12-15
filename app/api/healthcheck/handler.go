package healthcheck

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"time"
)

// HealthCheck : Check api health
func HealthCheck(ctx *fasthttp.RequestCtx) {
	logger := ctx.UserValue("logger").(zerolog.Logger)

	body := struct {
		Name string `json:"route_name"`
	}{
		Name: "health check",
	}

	time.Sleep(3 * time.Millisecond)
	logger.Debug().Msg("test")
	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
