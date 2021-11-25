package v1

import (
	"github.com/alexandre-slp/event-broker/app/api"
	"github.com/alexandre-slp/event-broker/domain/eventhandler"
)

func NewService() *InitializeEventV1Service {
	return &InitializeEventV1Service{}
}

type InitializeEventV1Service struct{}

func (InitializeEventV1Service) InitPaths(routerService api.Router) {
	v1 := routerService.Group("/api/v1")

	event := v1.Group("/event")
	event.GET("", eventhandler.ListEvent)
	event.POST("", eventhandler.CreateEvent)
	event.PUT("/{id}", eventhandler.UpdateEvent)
	event.DELETE("/{id}", eventhandler.DeleteEvent)
	event.POST("/send", eventhandler.SendEvent)
}