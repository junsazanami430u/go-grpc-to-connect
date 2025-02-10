package interceptor

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// NewUnaryValidationInterceptor は、gRPCのユニアリインターセプターを作成する
func NewUnaryValidationInterceptor() grpc.UnaryServerInterceptor {
	validator, err := protovalidate.New()
	if err != nil {
		panic("failed to create protovalidate validator: " + err.Error())
	}

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// リクエストが proto.Message であるかチェック
		msg, ok := req.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "request is not a proto message")
		}

		// バリデーション実行
		if err := validator.Validate(msg); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
		}

		// ハンドラーの実行
		return handler(ctx, req)
	}
}

// NewStreamValidationInterceptor は、gRPCのストリームインターセプターを作成する
func NewStreamValidationInterceptor() grpc.StreamServerInterceptor {
	validator, err := protovalidate.New()
	if err != nil {
		panic("failed to create protovalidate validator: " + err.Error())
	}

	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		wrappedStream := &validatingStream{
			ServerStream: ss,
			validator:    validator, // 修正: validator をそのまま格納
		}

		// ハンドラーの実行
		return handler(srv, wrappedStream)
	}
}

// validatingStream は、受信したメッセージを検証するためのラッパー
type validatingStream struct {
	grpc.ServerStream
	validator protovalidate.Validator // 修正: インターフェース型として保持
}

// RecvMsg は、受信したメッセージをバリデーションする
func (s *validatingStream) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}

	// リクエストが proto.Message であるかチェック
	msg, ok := m.(proto.Message)
	if !ok {
		return status.Errorf(codes.Internal, "request is not a proto message")
	}

	// バリデーション実行
	if err := s.validator.Validate(msg); err != nil {
		return status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}

	return nil
}
