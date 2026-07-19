package logger

import (
	"fmt"

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
	l.log.Errorf(getLogContent(prefix, message))
}

func (l *AppLogger) Debug(prefix string, message string) {
	l.log.Debugf(getLogContent(prefix, message))
}

func (l *AppLogger) Warn(prefix string, message string) {
	l.log.Warnf(getLogContent(prefix, message))
}

func getLogContent(prefix string, message string) string {
	return fmt.Sprintf("%s: %s", prefix, message)
}
