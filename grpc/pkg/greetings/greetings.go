package greetings

import (
	"context"
	"fmt"

	"github.com/junsazanami430u/go-grpc-to-connect/grpc/pkg/logger"
	greetingsv1 "github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
)

type GreetingsServer struct {
	greetingsv1.UnimplementedGreetingsServiceServer
}

func (s *GreetingsServer) GetGreetings(ctx context.Context, req *greetingsv1.GetGreetingsRequest) (*greetingsv1.GetGreetingsResponse, error) {
	log := logger.FromContext(ctx)
	if req.Greetings == "" {
		return nil, errorInvalidArgument("invalid greetings").Err()
	}
	if req.Name == "" {
		return nil, errorInvalidArgument("invalid name").Err()
	}
	log.Debug("greetings.greetings", "Name", req.Name, "greetings", req.Greetings)
	return &greetingsv1.GetGreetingsResponse{Greetings: fmt.Sprintf("Hello %s", req.Name)}, nil
}
