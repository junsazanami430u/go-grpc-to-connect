package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/baleen-dyamaguchi/go-grpc-to-connect/grpc/pkg/greetings"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/grpc/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"google.golang.org/grpc"
)

func main() {
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	log := logger.FromContext(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Error("failed to listen: %v", "error", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	greetingsv1.RegisterGreetingsServiceServer(grpcServer, &greetings.GreetingsServer{})
	go func() {
		log.Info("greetings.server is running", "Addr", lis.Addr())
	}()
	grpcServer.Serve(lis)

}
