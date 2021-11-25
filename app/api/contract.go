package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_requesthandler.go -package=mocks . RequestHandler
type RequestHandler interface {
	InitRequestHandler(...PathInitializer) func(*fasthttp.RequestCtx)
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_pathinitializer.go -package=mocks . PathInitializer
type PathInitializer interface {
	InitPaths(Router)
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_router.go -package=mocks . Router
type Router interface {
	GET(string, func(*fasthttp.RequestCtx))
	POST(string, func(*fasthttp.RequestCtx))
	PUT(string, func(*fasthttp.RequestCtx))
	PATCH(string, func(*fasthttp.RequestCtx))
	DELETE(string, func(*fasthttp.RequestCtx))
	Handler() func(*fasthttp.RequestCtx)
	Group(string) *router.Group
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_customerror.go -package=mocks . CustomError
type CustomError interface {
	Msg() string
	Code() int
}
