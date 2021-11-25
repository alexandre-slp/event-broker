package eventhandler

import (
	"github.com/alexandre-slp/event-broker/domain"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// CreateEvent : Create new event
func CreateEvent(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b := domain.SavedEvent{}
	err := json.Unmarshal(ctx.PostBody(), &b)
	if err != nil {
		return
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
