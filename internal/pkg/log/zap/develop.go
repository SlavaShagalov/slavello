package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DevelopConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func NewDev() *zap.Logger {
	consoleCfg := DevelopConfig()

	consoleEncoder := zapcore.NewConsoleEncoder(consoleCfg)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	logger := zap.New(consoleCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
