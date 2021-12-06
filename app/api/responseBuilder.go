package api

import (
	"github.com/alexandre-slp/event-broker/infra"
	"github.com/valyala/fasthttp"
)

type Response struct {
	ctx  *fasthttp.RequestCtx
	body interface{}
	code int
}

func (r Response) WriteResponse() {
	responseBody, err := infra.Json.Marshal(&r.body)
	if err != nil {
		return
	}

	r.ctx.SetStatusCode(r.code)
	r.ctx.Response.Header.Set("Content-Type", "application/json")

	_, err = r.ctx.Write(responseBody)
	if err != nil {
		return
	}
}

func NewResponse(ctx *fasthttp.RequestCtx, body interface{}, code int) *Response {
	return &Response{
		ctx:  ctx,
		body: body,
		code: code,
	}
}
