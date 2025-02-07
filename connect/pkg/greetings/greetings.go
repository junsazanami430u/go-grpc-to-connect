package greetings

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"
)

type GreetingsServer struct {
	greetingsv1connect.UnimplementedGreetingsServiceHandler
}

func (s *GreetingsServer) GetGreetings(ctx context.Context, req *connect.Request[greetingsv1.GetGreetingsRequest]) (*connect.Response[greetingsv1.GetGreetingsResponse], error) {
	log := logger.FromContext(ctx)
	if req.Msg.Greetings == "" {
		return nil, error_InvalidArgment("invalid greetings")
	}
	if req.Msg.Name == "" {
		return nil, error_InvalidArgment("invalid name")
	}
	log.Debug("greetings.greetings", "Name", req.Msg.Name, "greetings", req.Msg.Greetings)
	return connect.NewResponse(&greetingsv1.GetGreetingsResponse{Greetings: fmt.Sprintf("Hello %s", req.Msg.Name)}), nil
}
