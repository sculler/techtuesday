package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.Logger
}

type ILogger interface {
	Debug(ctx context.Context, msg string, f ...zap.Field)
	Error(ctx context.Context, msg string, f ...zap.Field)
	Info(ctx context.Context, msg string, f ...zap.Field)
	Fatal(msg string, f ...zap.Field)
}

func NewLogger() (*Logger, error) {
	cfg := getLoggerConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	l := &Logger{logger}

	l.Zap = l.Zap.WithOptions(
		zap.AddCallerSkip(1),
	)

	return l, nil
}

func getLoggerConfig() zap.Config {
	const local = "local"

	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	encoding := "console"

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zap.Config{
		Level:             level,
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         encoding,
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func (l *Logger) Debug(ctx context.Context, msg string, f ...zap.Field) {
	fields := appendContextFields(ctx, f...)
	l.Zap.Debug(msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, f ...zap.Field) {
	fields := appendContextFields(ctx, f...)
	l.Zap.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, f ...zap.Field) {
	fields := appendContextFields(ctx, f...)
	l.Zap.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, f ...zap.Field) {
	l.Zap.Fatal(msg, f...)
}

func appendContextFields(ctx context.Context, f ...zap.Field) []zap.Field {
	fields := append(
		f,
		zap.String("ip", ctx.Value("ClientIP").(string)),
		zap.String("req-id", ctx.Value("RequestId").(string)),
	)
	return fields
}



