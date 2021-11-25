package eventhandler

import (
	"github.com/alexandre-slp/event-broker/domain"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// SendEvent : Send an event
func SendEvent(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b := domain.SentEvent{}
	err := json.Unmarshal(ctx.PostBody(), &b)
	if err != nil {
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	response, err := json.Marshal(&b)
	if err != nil {
		return
	}
	_, err = ctx.Write(response)
	if err != nil {
		return
	}
}
