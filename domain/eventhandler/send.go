package eventhandler

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/valyala/fasthttp"
)

// SendEvent : Send an event
func SendEvent(ctx *fasthttp.RequestCtx) {
	body := domain.SentEvent{}
	err := infra.Json.Unmarshal(ctx.PostBody(), &body)
	if err != nil {
		return
	}

	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
