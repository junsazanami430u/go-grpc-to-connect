package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestError_InvalidArgument(t *testing.T) {
	expected := status.New(codes.InvalidArgument, "test")
	status := error_InvalidArgument("test")
	assert.Equal(t, status, expected)
}
func TestError_Unknown(t *testing.T) {
	expected := status.New(codes.Unknown, "test")
	status := error_Unknown("test")
	assert.Equal(t, status, expected)
}
