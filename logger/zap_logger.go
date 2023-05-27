package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapSugarLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() Logger {
	// customize as needed
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Add color to log level
	logger, err := cfg.Build()
	if err != nil {
		log.Println("Failed to initialize zap logger")
		return nil
	}
	return &ZapSugarLogger{logger: logger.Sugar()}
}

func (l *ZapSugarLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *ZapSugarLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
func (l *ZapSugarLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}
func (l *ZapSugarLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

// Implement other methods as needed
