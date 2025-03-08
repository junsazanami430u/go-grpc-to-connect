package greetings

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/junsazanami430u/go-grpc-to-connect/connect/pkg/logger"
	greetingsv1 "github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"

	elizav1 "github.com/junsazanami430u/examples-go/pkg/eliza/gen/connectrpc/eliza/v1"
)

type GreetingsServer struct {
	greetingsv1connect.UnimplementedGreetingsServiceHandler
}

func (s *GreetingsServer) GetGreetings(ctx context.Context, req *connect.Request[greetingsv1.GetGreetingsRequest]) (*connect.Response[greetingsv1.GetGreetingsResponse], error) {
	log := logger.FromContext(ctx)
	if req.Msg.Greetings == "" {
		return nil, errorInvalidArgument("invalid greetings")
	}
	if req.Msg.Name == "" {
		return nil, errorInvalidArgument("invalid name")
	}
	log.Debug("greetings.greetings", "Name", req.Msg.Name, "greetings", req.Msg.Greetings)
	return connect.NewResponse(&greetingsv1.GetGreetingsResponse{Greetings: fmt.Sprintf("Hello %s", req.Msg.Name)}), nil
}

func (s *GreetingsServer) GetGoodBye(ctx context.Context, req *connect.Request[elizav1.GoodByeRequest]) (*connect.Response[elizav1.GoodByeResponse], error) {
	log := logger.FromContext(ctx)
	if req.Msg.Sentence == "" {
		return nil, errorInvalidArgument("invalid name")
	}
	log.Debug("greetings.goodbye", "Sentence", req.Msg.Sentence)
	return connect.NewResponse(&elizav1.GoodByeResponse{Sentence: fmt.Sprintf("Goodbye %s", req.Msg.Sentence)}), nil
}
