package main

import (
	"context"
	"log/slog"

	"github.com/baleen-dyamaguchi/go-grpc-to-connect/grpc/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := logger.InitLogger(context.Background(), slog.LevelInfo)

	log := logger.FromContext(ctx)

	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("fail to dial: %v", "error", err)
	}
	defer conn.Close()

	client := greetingsv1.NewGreetingsServiceClient(conn)

	res, err := client.GetGreetings(
		ctx,
		&greetingsv1.GetGreetingsRequest{Name: "Jane", Greetings: "Hello"},
	)
	if err != nil {
		log.Error("greetings.client", "error", err.Error())
		return
	}
	log.Info("greetings.client", "greetings", res.Greetings)
}
