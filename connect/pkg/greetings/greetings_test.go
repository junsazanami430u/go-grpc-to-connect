package greetings

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/interceptor"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/connect/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	"github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1/greetingsv1connect"
	"github.com/stretchr/testify/assert"
)

func TestGetGreetings(t *testing.T) {
	expected := connect.NewResponse(&greetingsv1.GetGreetingsResponse{
		Greetings: "Hello John",
	})
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	g := &GreetingsServer{}
	result, err := g.GetGreetings(ctx, connect.NewRequest(&greetingsv1.GetGreetingsRequest{
		Name:      "John",
		Greetings: "Hello",
	}))

	assert.NoError(t, err)
	assert.Equal(t, result, expected)

	result, err = g.GetGreetings(ctx, connect.NewRequest(&greetingsv1.GetGreetingsRequest{
		Name:      "",
		Greetings: "Hello",
	}))

	assert.Error(t, err)
	result, err = g.GetGreetings(ctx, connect.NewRequest(&greetingsv1.GetGreetingsRequest{
		Name:      "",
		Greetings: "",
	}))

	assert.Error(t, err)
}

type down func()

func setup() (context.Context, down, *httptest.Server) {
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	log := logger.FromContext(ctx)
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Error("greetings.server", "error", err)
	}
	greeter := &GreetingsServer{}
	mux := http.NewServeMux()
	path, handler := greetingsv1connect.NewGreetingsServiceHandler(greeter, connect.WithInterceptors(validateInterceptor, interceptor.NewValidateInterceptor()))
	mux.Handle(path, handler)
	server := httptest.NewServer(mux)
	down := func() {
		server.Close()
	}
	return ctx, down, server
}

func TestApiGreetings(t *testing.T) {
	ctx, down, server := setup()
	defer down()

	expected := connect.NewResponse(&greetingsv1.GetGreetingsResponse{
		Greetings: "Hello John",
	})
	client := greetingsv1connect.NewGreetingsServiceClient(
		server.Client(),
		server.URL,
	)
	resp, err := client.GetGreetings(ctx, connect.NewRequest(&greetingsv1.GetGreetingsRequest{
		Name:      "John",
		Greetings: "Hello",
	}))
	assert.NoError(t, err)
	assert.Equal(t, resp.Msg.Greetings, expected.Msg.Greetings)
}
