package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() Logger {
	logger := logrus.New()
	return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}
func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
func (l *LogrusLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *LogrusLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Add other log levels or methods as needed
