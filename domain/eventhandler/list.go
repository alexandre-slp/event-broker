package eventhandler

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// ListEvent : List saved events
func ListEvent(ctx *fasthttp.RequestCtx) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	r := struct {
		Name string `json:"route_name"`
	}{
		Name: "list",
	}
	response, err := json.Marshal(&r)

	ctx.Response.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}
	_, err = ctx.Write(response)
	if err != nil {
		return
	}
}
