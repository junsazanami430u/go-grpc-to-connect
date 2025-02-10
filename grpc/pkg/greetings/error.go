package greetings

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func errorInvalidArgument(msg string) *status.Status {
	return status.New(codes.InvalidArgument, msg)
}

func errorUnknown(msg string) *status.Status {
	return status.New(codes.Unknown, msg)
}
