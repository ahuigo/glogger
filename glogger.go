package glogger

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
)

// Refer: go.uber.org/zap@v1.16.0/example_test.go
// zapcore.NewTee( zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),

// Global logger
var glogger *zap.SugaredLogger

func init(){
    if isInDev(){
        glogger = GetLogger("root", zap.DebugLevel)
    }else{
        glogger = GetLogger("root", zap.InfoLevel)
    }
}

func Error(args ...interface{}) {
	glogger.Error(args...)
}
func Warn(args ...interface{}) {
	glogger.Warn(args...)
}
func Info(args ...interface{}) {
	glogger.Info(args...)
}
func Debug(args ...interface{}) {
	glogger.Debug(args...)
}

// SetGlogger
func SetGlogger(name string, level zapcore.Level) {
	glogger = GetLogger(name, level)
}

func isInDev() bool{
    return os.Getenv("APP_ENV") == "dev"
}

/**
console and json ok
trace with error ok
level ok

output path:
	rotate with date
*/
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
	if !isInDev() {
		encoding = "json"
	}
	atom := zap.NewAtomicLevelAt(level)
	// 1. with zapCore config
	config := zap.Config{
		Level:            atom,
		Development:      false,
		Encoding:         encoding,
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	loggerRaw, _ := config.Build()
	// 2. zap option
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1), // logger.core.callerSkip += 2
		zap.AddStacktrace(levelEnablerFunc(zapcore.ErrorLevel)),
		// zap.Development(),
	}
	logger := loggerRaw.WithOptions(options...).Named(name).Sugar()
	return logger
}

/* The inner builder: Config.Build():
var zapCores []zapcore.Core // zapcore.WriteSyncer, encoder, ...
	writeSyncer:= zap.Open(outputPath.Path)
	writeSyncer:= zapcore.AddSync(&lumberjackWrite)
	zap.CombineWriteSyncers(writeSyncers...)
core := zapcore.NewTee(zapCores...)
var options []zap.Option // caller, zap.Fileds, zap.ErrorOutput(writeSyncer) ...
zap.New(core, options...).Sugar()
*/

func levelEnablerFunc(level zapcore.Level) zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level //zapcore.ErrorLevel
	})
}

func JsonEncode(data interface{}) string {
	jsonBytes, _ := json.Marshal(data)
	return (string(jsonBytes))
}
