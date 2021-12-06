package eventhandler

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/valyala/fasthttp"
)

// ListEvent : List saved events
func ListEvent(ctx *fasthttp.RequestCtx) {
	body := struct {
		Name string `json:"route_name"`
	}{
		Name: "list",
	}

	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
