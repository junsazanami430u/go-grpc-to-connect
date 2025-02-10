package main

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/greetings"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/interceptor"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/logger"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	log := logger.FromContext(ctx)

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Error("greetings.server", "error", err)
	}
	mux := http.NewServeMux()
	path, handler := greetingsv1connect.NewGreetingsServiceHandler(&greetings.GreetingsServer{}, connect.WithInterceptors(validateInterceptor, interceptor.NewValidateInterceptor()))
	mux.Handle(path, handler)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		log.Info("greetings.server is running", "Host", server.Addr)
	}()

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Error("Failed to start server", "error", err)
	}
}
