package loggerpkg

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfig struct {
	LogFile  string
	LogLevel zapcore.Level
}

func NewZapLogger(config *ZapConfig) *zap.Logger {
	encoder := getDefaultEncoder()
	fw := mustCreateFile(config.LogFile)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(fw),
		config.LogLevel,
	)
	return zap.New(core, zap.AddCaller())
}

func getDefaultEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	encoderConfig.ConsoleSeparator = "|"
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return encoder
}

func mustCreateFile(logFile string) *os.File {
	fw, err := os.Create(logFile)
	if err != nil {
		panic(fmt.Errorf("cannot create log file (%s): %w", logFile, err))
	}
	return fw
}
