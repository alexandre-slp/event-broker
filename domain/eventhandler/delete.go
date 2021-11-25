package eventhandler

import (
	"fmt"
	"github.com/alexandre-slp/event-broker/domain"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// DeleteEvent : Delete event
func DeleteEvent(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b := domain.SavedEvent{
		Name: fmt.Sprint(ctx.UserValue("id")),
	}
	response, err := json.Marshal(&b)

	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}
	_, err = ctx.Write(response)
	if err != nil {
		return
	}
}
