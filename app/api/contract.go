package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_requesthandler.go -package=mocks . RequestHandler
type RequestHandler interface {
	InitRequestHandler(...PathInitializer) fasthttp.RequestHandler
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_responsebuilder.go -package=mocks . ResponseBuilder
type ResponseBuilder interface {
	WriteResponse()
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_pathinitializer.go -package=mocks . PathInitializer
type PathInitializer interface {
	InitPaths(Router)
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_router.go -package=mocks . Router
type Router interface {
	GET(string, fasthttp.RequestHandler)
	POST(string, fasthttp.RequestHandler)
	PUT(string, fasthttp.RequestHandler)
	PATCH(string, fasthttp.RequestHandler)
	DELETE(string, fasthttp.RequestHandler)
	Handler() fasthttp.RequestHandler
	Group(string) *router.Group
}

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/mock_customerror.go -package=mocks . CustomError
type CustomError interface {
	ReturnedMsg() string
	ReturnedCode() int
}
