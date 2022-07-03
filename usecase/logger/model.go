package logger

import (
	"context"
	"os"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func newEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.TimeKey = "time"
	cfg.LevelKey = "severity"
	cfg.MessageKey = "message"
	cfg.EncodeLevel = EncodeLevel
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

func New(ctx context.Context) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = newEncoderConfig()
	logger, _ := cfg.Build()

	trace := ForContext(ctx)
	if trace != nil {
		fields := zapdriver.TraceContext(trace.TraceID, trace.SpanID, trace.Sampled, os.Getenv("GCP_PROJECT_ID"))
		logger = logger.With(fields...)
	}

	return logger
}
