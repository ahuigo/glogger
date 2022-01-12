package glogger

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Refer: go.uber.org/zap@v1.16.0/example_test.go
// zapcore.NewTee( zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),

// Global logger
var Glogger = GetLogger("root", zap.DebugLevel)

// SetGlogger
func SetGlogger(name string, level zapcore.Level) {
	Glogger = GetLogger(name, level)
}

func GetLogger(name string, level zapcore.Level) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		// EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	encoding := "console"
	if os.Getenv("APP_ENV") != "" && os.Getenv("APP_ENV") != "dev" {
		encoding = "json"

	}
	atom := zap.NewAtomicLevelAt(level)
	config := zap.Config{
		Level:            atom,
		Development:      false,
		Encoding:         encoding,
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	// option := make(zap.Option, 0)
	loggerRaw, _ := config.Build()
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		// logger.core.callerSkip += 2
		zap.AddStacktrace(zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})),
		// zap.Development(),
	}
	logger := loggerRaw.WithOptions(options...).Named(name).Sugar()
	return logger
}

func JsonEncode(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return (string(jsonBytes))
}
