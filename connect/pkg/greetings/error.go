package greetings

import (
	"connectrpc.com/connect"
	"github.com/friendsofgo/errors"
)

func errorInvalidArgument(msg string) error {
	return connect.NewError(connect.CodeInvalidArgument, errors.New(msg))
}

func errorUnknown(msg string) error {
	return connect.NewError(connect.CodeUnknown, errors.New(msg))
}
