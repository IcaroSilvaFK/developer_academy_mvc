package utils

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log        *zap.Logger
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()

}

func Info(message string, tags ...zapcore.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zapcore.Field) {
	tags = append(tags, zap.NamedError("error", err))

	log.Error(message, tags...)
	log.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {

	level := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL)))

	switch level {
	case "info":
		{
			return zapcore.InfoLevel
		}
	case "error":
		{
			return zapcore.ErrorLevel
		}
	default:
		{
			return zapcore.InfoLevel
		}
	}

}
