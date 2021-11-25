package api

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

func PanicHandler(ctx *fasthttp.RequestCtx, error interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	r := struct {
		Error string `json:"Error"`
	}{
		Error: error.(CustomError).Msg(),
	}
	response, _ := json.Marshal(&r)
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(error.(CustomError).Code())
	_, err := ctx.Write(response)
	if err != nil {
		return
	}
}
