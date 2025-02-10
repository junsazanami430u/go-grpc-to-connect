package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestError_InvalidArgument(t *testing.T) {
	expected := status.New(codes.InvalidArgument, "test")
	status := errorInvalidArgument("test")
	assert.Equal(t, status, expected)
}
func TestError_Unknown(t *testing.T) {
	expected := status.New(codes.Unknown, "test")
	status := errorUnknown("test")
	assert.Equal(t, status, expected)
}
