package util

import "go.uber.org/zap"

type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)

	Info(args ...any)
	Infof(format string, args ...any)

	Warn(args ...any)
	Warnf(format string, args ...any)

	Error(args ...any)
	Errorf(format string, args ...any)
}

type zapLogger struct {
	sugar  *zap.SugaredLogger
	logger *zap.Logger
}

func NewZapLogger() Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		stderr("failed to initialize zap logger")
	}
	sugar := logger.Sugar()

	return &zapLogger{
		logger: logger,
		sugar:  sugar,
	}
}

func (z zapLogger) Debug(args ...any) {
	z.sugar.Debug(args...)
}

func (z zapLogger) Debugf(format string, args ...any) {
	z.sugar.Debugf(format, args...)
}

func (z zapLogger) Info(args ...any) {
	z.sugar.Info(args...)
}

func (z zapLogger) Infof(format string, args ...any) {
	z.sugar.Infof(format, args...)
}

func (z zapLogger) Warn(args ...any) {
	z.sugar.Warn(args...)
}

func (z zapLogger) Warnf(format string, args ...any) {
	z.sugar.Warnf(format, args...)
}

func (z zapLogger) Error(args ...any) {
	z.sugar.Error(args...)
}

func (z zapLogger) Errorf(format string, args ...any) {
	z.sugar.Errorf(format, args...)
}

func stderr(msg string) {
	print(msg)
}
