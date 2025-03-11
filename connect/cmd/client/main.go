package main

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	elizav1 "github.com/junsazanami430u/examples-go/pkg/eliza/buf/v1"
	elizav1connect "github.com/junsazanami430u/examples-go/pkg/eliza/buf/v1/bufv1connect"
	"github.com/junsazanami430u/go-grpc-to-connect/connect/pkg/logger"
	greetingsv1 "github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"
)

func main() {
	ctx := logger.InitLogger(context.Background(), slog.LevelInfo)

	log := logger.FromContext(ctx)

	client := greetingsv1connect.NewGreetingsServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.GetGreetings(
		ctx,
		connect.NewRequest(&greetingsv1.GetGreetingsRequest{Name: "Jane", Greetings: "Hello"}),
	)
	if err != nil {
		log.Error("greetings.client", "error", err.Error())
		return
	}
	log.Info("greetings.client", "greetings", res.Msg.Greetings)

	elizaClient := elizav1connect.NewElizaServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	elizaRes, err := elizaClient.GoodBye(
		ctx,
		connect.NewRequest(&elizav1.GoodByeRequest{Sentence: "Goodbye"}),
	)
	if err != nil {
		log.Error("eliza.client", "error", err.Error())
		return
	}
	log.Info("eliza.client", "sentence", elizaRes.Msg.Sentence)

	elizaResE, err := elizaClient.GoodBye(
		ctx,
		connect.NewRequest(&elizav1.GoodByeRequest{Sentence: "Goodbye ggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"}),
	)
	if err != nil {
		log.Error("eliza.client", "error", err.Error())
		return
	}
	log.Info("eliza.client", "sentence", elizaResE.Msg.Sentence)
}
