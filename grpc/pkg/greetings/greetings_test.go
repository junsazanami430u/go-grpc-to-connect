package greetings

import (
	"context"
	"log"
	"log/slog"
	"net"
	"testing"

	"github.com/baleen-dyamaguchi/go-grpc-to-connect/grpc/pkg/logger"
	greetingsv1 "github.com/baleen-dyamaguchi/go-grpc-to-connect/pkg/gen/proto/greetings/v1"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/test/bufconn"
)

func TestGetGreetings(t *testing.T) {
	expected := &greetingsv1.GetGreetingsResponse{
		Greetings: "Hello John",
	}
	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	g := &GreetingsServer{}
	result, err := g.GetGreetings(ctx, &greetingsv1.GetGreetingsRequest{
		Name:      "John",
		Greetings: "Hello",
	})

	assert.NoError(t, err)
	assert.Equal(t, result, expected)

	result, err = g.GetGreetings(ctx, &greetingsv1.GetGreetingsRequest{
		Name:      "",
		Greetings: "Hello",
	})

	assert.Error(t, err)
	result, err = g.GetGreetings(ctx, &greetingsv1.GetGreetingsRequest{
		Name:      "",
		Greetings: "",
	})

	assert.Error(t, err)
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type down func()

func setup() (context.Context, down) {

	ctx := logger.InitLogger(context.Background(), slog.LevelDebug)
	greeter := &GreetingsServer{}

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_validator.UnaryServerInterceptor(),
		),
	)

	greetingsv1.RegisterGreetingsServiceServer(s, greeter)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	return ctx, func() {
		s.GracefulStop()
	}
}
func bufDialer(ctx context.Context, _ string) (net.Conn, error) {
	return lis.Dial()
}

func TestApiGreetings(t *testing.T) {
	ctx, down := setup()
	defer down()

	expected := &greetingsv1.GetGreetingsResponse{
		Greetings: "Hello John",
	}
	resolver.SetDefaultScheme("passthrough")
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(bufDialer), opt)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()
	client := greetingsv1.NewGreetingsServiceClient(conn)
	resp, err := client.GetGreetings(ctx, &greetingsv1.GetGreetingsRequest{
		Name:      "John",
		Greetings: "Hello",
	})
	assert.NoError(t, err)
	assert.Equal(t, resp.Greetings, expected.Greetings)
}
