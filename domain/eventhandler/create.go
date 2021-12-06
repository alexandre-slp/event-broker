package eventhandler

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/valyala/fasthttp"
)

// CreateEvent : Create new event
func CreateEvent(ctx *fasthttp.RequestCtx) {
	body := domain.SavedEvent{}
	err := infra.Json.Unmarshal(ctx.PostBody(), &body)
	if err != nil {
		return
	}

	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
