package logger

import (
	"context"

	"github.com/google/uuid"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	logpkg "github.com/shengchaohua/red-packet-backend/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey string

const (
	ctxKeyTraceId ctxKey = "TraceId"
	ctxKeyLogger  ctxKey = "Logger"
)

var zapLogger *zap.Logger

func InitLogger(serverConfig *config.ServerConfig) {
	logLevel := zapcore.DebugLevel
	if serverConfig.IsLiveEnv() {
		logLevel = zapcore.InfoLevel
	}
	zapConfig := &logpkg.ZapConfig{
		LogFile:  serverConfig.Log,
		LogLevel: logLevel,
	}
	zapLogger = logpkg.NewZapLogger(zapConfig)
}

// NewCtxWithTraceId returns a context which knows its request ID
func NewCtxWithTraceId() context.Context {
	ctx := context.Background()
	traceId := uuid.NewString()
	newLogger := zapLogger.With(zap.String(string(ctxKeyTraceId), traceId))
	newCtx := context.WithValue(ctx, ctxKeyLogger, newLogger)
	return newCtx
}

// Logger returns a zap logger in ctx
func Logger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return zapLogger
	}
	if newLogger, ok := ctx.Value(ctxKeyLogger).(*zap.Logger); ok {
		return newLogger
	}
	return zapLogger
}
