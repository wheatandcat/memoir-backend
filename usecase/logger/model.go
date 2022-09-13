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

func New(ctx context.Context) *zap.Logger {
	if os.Getenv("APP_ENV") == "local" {
		cfg := newDevelopmentConfig()
		logger, _ := cfg.Build()

		return logger
	}

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	stdoutSink := zapcore.Lock(os.Stdout)
	stderrSink := zapcore.Lock(os.Stderr)

	enc := zapcore.NewJSONEncoder(newProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(enc, stderrSink, highPriority),
		zapcore.NewCore(enc, stdoutSink, lowPriority),
	)

	logger := zap.New(core)

	trace := ForContext(ctx)
	if trace != nil {
		fields := zapdriver.TraceContext(trace.TraceID, trace.SpanID, trace.Sampled, os.Getenv("GCP_PROJECT_ID"))
		logger = logger.With(fields...)
	}

	return logger
}
