package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func NewRouterService(router *router.Router) *RouterService {
	return &RouterService{
		r: router,
	}
}

func NewRequestHandlerService(routerService *RouterService) *RequestHandlerService {
	return &RequestHandlerService{
		rs: routerService,
	}
}

type RequestHandlerService struct {
	rs *RouterService
}

func (rhs RequestHandlerService) InitRequestHandler(pi ...PathInitializer) func(ctx *fasthttp.RequestCtx) {
	for _, route := range pi {
		route.InitPaths(rhs.rs)
	}
	return rhs.rs.Handler()
}

type RouterService struct {
	r *router.Router
}

func (rs RouterService) GET(path string, handler func(*fasthttp.RequestCtx)) {
	rs.r.GET(path, handler)
}

func (rs RouterService) POST(path string, handler func(*fasthttp.RequestCtx)) {
	rs.r.POST(path, handler)
}

func (rs RouterService) PUT(path string, handler func(*fasthttp.RequestCtx)) {
	rs.r.PUT(path, handler)
}

func (rs RouterService) PATCH(path string, handler func(*fasthttp.RequestCtx)) {
	rs.r.PATCH(path, handler)
}

func (rs RouterService) DELETE(path string, handler func(*fasthttp.RequestCtx)) {
	rs.r.DELETE(path, handler)
}

func (rs RouterService) Handler() func(ctx *fasthttp.RequestCtx){
	return rs.r.Handler
}

func (rs RouterService) Group(path string) *router.Group {
	return rs.r.Group(path)
}
