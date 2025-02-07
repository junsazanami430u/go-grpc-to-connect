package main

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"
)

func main() {
	client := greetingsv1connect.NewGreetingsServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	ctx := logger.InitLogger(context.Background(), slog.LevelInfo)

	log := logger.FromContext(ctx)
	res, err := client.GetGreetings(
		ctx,
		connect.NewRequest(&greetingsv1.GetGreetingsRequest{Name: "Jane", Greetings: "Hello"}),
	)
	if err != nil {
		log.Error("greetings.client", "error", err.Error())
		return
	}
	log.Info("greetings.client", "greetings", res.Msg.Greetings)
}
