package logger

import (
	"context"
	"fmt"
	"os"

	"github.com/hardcore-os/plato/common/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger = &Logger{}
)

type Logger struct {
	Options

	logger *zap.Logger
}

func NewLogger(opts ...Option) {
	logger.Options = defaultOptions
	for _, o := range opts {
		o.apply(&logger.Options)
	}

	fileWriteSyncer := logger.getFileLogWriter()
	var core zapcore.Core
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if config.IsDebug() {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, fileWriteSyncer, zapcore.InfoLevel),
		)
	}
	logger.logger = zap.New(core, zap.WithCaller(true), zap.AddCallerSkip(logger.callerSkip))
}

func (l *Logger) getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", l.logDir, l.filename),
		MaxSize:    l.maxSize,
		MaxBackups: l.maxBackups,
		MaxAge:     l.maxAge,
		Compress:   l.compress,
	}

	return zapcore.AddSync(lumberJackLogger)
}

// DebugCtx ...
func DebugCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Debug(message, fields...)
}

// InfoCtx ...
func InfoCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Info(message, fields...)
}

// WarnCtx ...
func WarnCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Warn(message, fields...)
}

// ErrorCtx ...
func ErrorCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Error(message, fields...)
}

// DPanicCtx ...
func DPanicCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).DPanic(message, fields...)
}

// PanicCtx ...
func PanicCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Panic(message, fields...)
}

// FatalCtx ...
func FatalCtx(ctx context.Context, message string, fields ...zap.Field) {
	logger.logger.With(zap.String(traceID, GetTraceID(ctx))).Fatal(message, fields...)
}
