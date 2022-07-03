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

func newProductionEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.TimeKey = "time"
	cfg.LevelKey = "severity"
	cfg.MessageKey = "message"
	cfg.EncodeLevel = EncodeLevel
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

func newDevelopmentConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.Level.SetLevel(zap.InfoLevel)

	return cfg
}

func newProductConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.ErrorLevel)
	cfg.EncoderConfig = newProductionEncoderConfig()

	return cfg
}

func New(ctx context.Context) *zap.Logger {
	if os.Getenv("APP_ENV") == "local" {
		cfg := newDevelopmentConfig()
		logger, _ := cfg.Build()

		return logger
	}

	cfg := newProductConfig()
	logger, _ := cfg.Build()

	trace := ForContext(ctx)
	if trace != nil {
		fields := zapdriver.TraceContext(trace.TraceID, trace.SpanID, trace.Sampled, os.Getenv("GCP_PROJECT_ID"))
		logger = logger.With(fields...)
	}

	return logger
}
