package endpoints

import (
	"context"

	gRPC "github.com/alexandre-slp/event-broker/infra/api/v1/gRPC"
	serializer "github.com/alexandre-slp/event-broker/infra/api/v1/serializer"
)

// EventServer is used to implement events service
type EventServer struct {
	gRPC.UnimplementedEventServer
}

// ListEvents implements list events
func (s *EventServer) ListEvents(ctx context.Context, in *serializer.ListEventRequest) (*serializer.ListEventResponse, error) {
	// TODO: import logic
	return &serializer.ListEventResponse{Events: "events"}, nil
}
