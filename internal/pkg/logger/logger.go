package logger

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/shengchaohua/red-packet-backend/internal/config"
	loggerpkg "github.com/shengchaohua/red-packet-backend/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey string

const (
	ctxKeyTraceId ctxKey = "InternalTraceId"
	ctxKeyLogger  ctxKey = "Logger"
)

var zapLogger *zap.Logger

func InitLogger(serverConfig *config.ServerConfig) {
	logLevel := zapcore.DebugLevel
	if serverConfig.IsLiveEnv() {
		logLevel = zapcore.InfoLevel
	}
	zapConfig := &loggerpkg.ZapConfig{
		LogFile:  serverConfig.LogFile,
		LogLevel: logLevel,
	}
	zapLogger = loggerpkg.NewZapLogger(zapConfig)
}

// NewCtxWithTraceId creates a new context with logger
// Note: this function should be called immediately in the most top layer
func NewCtxWithTraceId(ctx context.Context, prefix string) context.Context {
	traceId := uuid.NewString()
	if prefix != "" {
		traceId = fmt.Sprintf("%s_%s", prefix, traceId)
	}

	if _, ok := ctx.Value(ctxKeyLogger).(*zap.Logger); ok { // avoid setting up logger
		return ctx
	}

	newLogger := zapLogger.With(zap.String(string(ctxKeyTraceId), traceId))
	newCtx := context.WithValue(ctx, ctxKeyLogger, newLogger)
	return newCtx
}

// Logger returns a zap logger in ctx
func Logger(ctx context.Context) *zap.Logger {
	if newLogger, ok := ctx.Value(ctxKeyLogger).(*zap.Logger); ok {
		return newLogger
	}
	panic("logger has not been inited")
}
