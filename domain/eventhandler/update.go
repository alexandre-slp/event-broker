package eventhandler

import (
	"fmt"
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/domain"
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/valyala/fasthttp"
)

// UpdateEvent : Update event
func UpdateEvent(ctx *fasthttp.RequestCtx) {
	body := domain.SavedEvent{}
	err := infra.Json.Unmarshal(ctx.PostBody(), &body)
	if err != nil {
		return
	}

	body.Topics = append(body.Topics, fmt.Sprint(ctx.UserValue("id")))

	api.NewResponse(ctx, body, fasthttp.StatusOK).WriteResponse()
}
