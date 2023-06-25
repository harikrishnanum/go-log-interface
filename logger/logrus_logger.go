package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger  *logrus.Logger
	logFile *os.File
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

func (l *LogrusLogger) OpenLogFile(filename string) {
	// Create or Open file for logging
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	l.logger.SetFormatter(&logrus.JSONFormatter{})
	l.logger.SetOutput(file)
	l.logFile = file
}

func (l *LogrusLogger) CloseLogFile() {
	l.logFile.Close()
}

func NewLogrusLogger(env string) Logger {
	logger := logrus.New()
	if isProduction(env) {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}
	return &LogrusLogger{logger: logger}
}
