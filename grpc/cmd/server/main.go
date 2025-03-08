package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/junsazanami430u/go-grpc-to-connect/grpc/pkg/greetings"
	"github.com/junsazanami430u/go-grpc-to-connect/grpc/pkg/interceptor"
	"github.com/junsazanami430u/go-grpc-to-connect/grpc/pkg/logger"
	greetingsv1 "github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"google.golang.org/grpc"
)

func main() {
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	log := logger.FromContext(ctx)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Error("failed to listen: %v", "error", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
			interceptor.NewUnaryValidationInterceptor(),
		))
	greetingsv1.RegisterGreetingsServiceServer(grpcServer, &greetings.GreetingsServer{})
	go func() {
		log.Info("greetings.server is running", "Addr", lis.Addr())
	}()
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("failed to serve: %v", "error", err)
	}

}
