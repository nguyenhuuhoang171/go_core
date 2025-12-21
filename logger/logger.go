package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppLogger struct {
	log *zap.SugaredLogger
}

func NewLogger() AppLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	appLogger := AppLogger{log: logger.Sugar()}
	return appLogger
}

// Sync flushes any buffered log entries.
func (l *AppLogger) Sync() {
	l.log.Sync()
}

func (l *AppLogger) Info(prefix string, message string) {
	l.log.Infof("%s: %s", prefix, message)
}

func (l *AppLogger) Error(prefix string, message string) {
	l.log.Errorf("%s: %s", prefix, message)
}

func (l *AppLogger) Debug(prefix string, message string) {
	l.log.Debugf("%s: %s", prefix, message)
}

func (l *AppLogger) Warn(prefix string, message string) {
	l.log.Warnf("%s: %s", prefix, message)
}
