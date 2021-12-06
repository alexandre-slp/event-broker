package main

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/app/api/healthcheck"
	"github.com/alexandre-slp/event-broker/app/api/v1"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
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
	

	s := fasthttp.Server{
		Handler:                            rh,
		Name:                               "Test server",
		Concurrency:                        0,
		ReadBufferSize:                     0,
		WriteBufferSize:                    0,
		ReadTimeout:                        0,
		WriteTimeout:                       0,
		IdleTimeout:                        0,
		MaxConnsPerIP:                      0,
		MaxRequestsPerConn:                 0,
	}
	err := s.ListenAndServe(":8080")
	if err != nil {
		return 
	}
}
