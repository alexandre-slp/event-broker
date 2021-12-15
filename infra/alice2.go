package infra

import (
	"github.com/valyala/fasthttp"
)

type middleware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Alice(middlewares ...middleware) fasthttp.RequestHandler {
	var allMiddlewares fasthttp.RequestHandler
	for i, m := range middlewares {
		if i == 0 {
			allMiddlewares = m
			continue
		}
		allMiddlewares = chain(m)
	}
	return allMiddlewares
}

func chain(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		next(ctx)
	}
}
