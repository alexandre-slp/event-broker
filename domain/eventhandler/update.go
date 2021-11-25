package eventhandler

import (
	"fmt"
	"github.com/alexandre-slp/event-broker/domain"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// UpdateEvent : Update event
func UpdateEvent(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b := domain.SavedEvent{}
	err := json.Unmarshal(ctx.PostBody(), &b)
	if err != nil {
		return
	}

	b.Topics = append(b.Topics, fmt.Sprint(ctx.UserValue("id")))
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
