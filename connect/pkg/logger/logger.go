// pkg/logger/logger.go
package logger

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type contextKey string

const loggerKey contextKey = "logger"

var (
	instance *slog.Logger
	once     sync.Once
)

// InitLogger はロガーを初期化してコンテキストに埋め込む
func InitLogger(ctx context.Context, level slog.Level) context.Context {
	once.Do(func() {
		opts := &slog.HandlerOptions{Level: level}
		instance = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	})
	return context.WithValue(ctx, loggerKey, instance)
}

// FromContext はコンテキストからロガーを取得する
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default() // デフォルトの `slog` を返す
}
