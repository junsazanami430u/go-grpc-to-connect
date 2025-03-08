package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"

	"github.com/junsazanami430u/go-grpc-to-connect/connect/pkg/greetings"
	"github.com/junsazanami430u/go-grpc-to-connect/connect/pkg/interceptor"
	"github.com/junsazanami430u/go-grpc-to-connect/connect/pkg/logger"
	"github.com/junsazanami430u/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"

	"github.com/junsazanami430u/examples-go/pkg/eliza"
	elizav1connect "github.com/junsazanami430u/examples-go/pkg/eliza/gen/connectrpc/eliza/v1/elizav1connect"

	"github.com/spf13/pflag"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {

	helpArg := pflag.BoolP("help", "h", false, "")
	streamDelayArg := pflag.DurationP(
		"server-stream-delay",
		"d",
		0,
		"The duration to delay sending responses on the server stream.",
	)
	pflag.Parse()

	if *helpArg {
		pflag.PrintDefaults()
		return
	}

	if *streamDelayArg < 0 {
		log.Printf("Server stream delay cannot be negative.")
		return
	}

	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	log := logger.FromContext(ctx)

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Error("greetings.server", "error", err)
	}
	mux := http.NewServeMux()
	path, handler := greetingsv1connect.NewGreetingsServiceHandler(&greetings.GreetingsServer{}, connect.WithInterceptors(validateInterceptor, interceptor.NewValidateInterceptor()))
	mux.Handle(path, handler)
	mux.Handle(elizav1connect.NewElizaServiceHandler(eliza.NewElizaServer(*streamDelayArg)))

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
