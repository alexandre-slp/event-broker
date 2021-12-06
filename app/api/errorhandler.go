package api

import (
	"github.com/valyala/fasthttp"
)

func PanicHandler(ctx *fasthttp.RequestCtx, error interface{}) {
	NewResponse(ctx, error, error.(CustomError).ReturnedCode()).WriteResponse()
}
