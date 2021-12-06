package eventhandler

import (
	"fmt"
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/valyala/fasthttp"
)

// DeleteEvent : Delete event
func DeleteEvent(ctx *fasthttp.RequestCtx) {
	body := domain.SavedEvent{
		Name: fmt.Sprint(ctx.UserValue("id")),
	}

	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
