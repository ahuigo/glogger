package glogger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global logger
var Glogger = GetLogger("root", zap.DebugLevel)

// SetGlogger
func SetGlogger(name string, level zapcore.Level) {
	Glogger = GetLogger(name, level)
}

func GetLogger(name string, level zapcore.Level) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		// StacktraceKey: "stacktrace",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		// EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(level)
	config := zap.Config{
		Level:            atom,
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	loggerRaw, _ := config.Build()
	logger := loggerRaw.Named(name).Sugar()
	// logger.core.callerSkip += 2
	return logger
}

func JsonEncode(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return (string(jsonBytes))
}
