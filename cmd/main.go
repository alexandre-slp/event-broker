package main

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/app/api/healthcheck"
	"github.com/alexandre-slp/event-broker/app/api/v1"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	r := router.New()
	r.PanicHandler = api.PanicHandler
	routerService := api.NewRouterService(r)
	requestHandlerService := api.NewRequestHandlerService(routerService)
	rh := requestHandlerService.InitRequestHandler(
		healthcheck.NewService(),
		v1.NewService(),
	)

	log.Fatal(fasthttp.ListenAndServe(":8080", rh))
}
