package domain

import (
	"context"

	"github.com/alexandre-slp/event-broker/app/api/v1/gRPC"
	"github.com/alexandre-slp/event-broker/app/api/v1/serializer"
)

// EventServer is used to implement events service
type EventServer struct {
	gRPC.UnimplementedEventServer
}

// ListEvents implements list events
func (s *EventServer) ListEvents(ctx context.Context, in *serializer.ListEventRequest) (*serializer.ListEventResponse, error) {
	// TODO: add logic
	return &serializer.ListEventResponse{Events: "events"}, nil
}
