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
	fw, err := os.Create(config.LogFile)
	if err != nil {
		panic(fmt.Errorf("cannot create log file (%s): %w", config.LogFile, err))
	}

	encoder := getDefaultEncoder()
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
