package domain

import (
	"context"
	"github.com/alexandre-slp/event-broker/app/api/gRPC"
	"github.com/alexandre-slp/event-broker/app/api/serializer"
)

// EventServer is used to implement events service
type EventServer struct {
	gRPC.UnimplementedEventServer
}

// ListEvents implements list events
func (s *EventServer) ListEvents(ctx context.Context, in *serializer.ListEventRequest) (*serializer.ListEventResponse, error) {
	return &serializer.ListEventResponse{Events: "events"}, nil
}
