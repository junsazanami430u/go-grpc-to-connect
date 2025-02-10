package greetings

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func error_InvalidArgument(msg string) *status.Status {
	return status.New(codes.InvalidArgument, msg)
}

func error_Unknown(msg string) *status.Status {
	return status.New(codes.Unknown, msg)
}
