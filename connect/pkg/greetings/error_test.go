package greetings

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestError_InvalidArgument(t *testing.T) {
	expected := connect.NewError(connect.CodeInvalidArgument, errors.New("test"))
	err := error_InvalidArgument("test")
	if result := new(connect.Error); errors.As(err, &result) {
		if err != nil {
		}
		assert.Equal(t, result.Code(), expected.Code())
		assert.Equal(t, result.Message(), expected.Message())
	}
}
func TestError_Unknown(t *testing.T) {
	expected := connect.NewError(connect.CodeUnknown, errors.New("test"))
	err := error_Unknown("test")
	if result := new(connect.Error); errors.As(err, &result) {
		if err != nil {
		}
		assert.Equal(t, result.Code(), expected.Code())
		assert.Equal(t, result.Message(), expected.Message())
	}
}
