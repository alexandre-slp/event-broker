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
	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
	logger.Debug().Msg("test")
		//Fields(map[string]string{
		//	"method":       fmt.Sprintf("%v", ctx.Method()),
		//	"elapsed_time": fmt.Sprintf("%vms", time.Now().Sub(ctx.UserValue("startRequest").(time.Time)).Milliseconds()),
		//}).
		//Interface("msg", ctx)
}
