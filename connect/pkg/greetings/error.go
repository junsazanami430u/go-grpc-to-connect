package greetings

import (
	"connectrpc.com/connect"
	"github.com/friendsofgo/errors"
)

func error_InvalidArgument(msg string) error {
	return connect.NewError(connect.CodeInvalidArgument, errors.New(msg))
}

func error_Unknown(msg string) error {
	return connect.NewError(connect.CodeUnknown, errors.New(msg))
}
