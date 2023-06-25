package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

type ZapSugarLogger struct {
	logger  *zap.SugaredLogger
	cfg     zap.Config
	logFile *os.File
}

/******** Methods for ZapSugarLogger ********/

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

func (l *ZapSugarLogger) CloseLogFile() {
	l.logFile.Close()
}

func (l *ZapSugarLogger) OpenLogFile(filename string) {
	// Create or Open file for logging
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	l.cfg.OutputPaths = []string{file.Name()}
	l.cfg.ErrorOutputPaths = []string{file.Name()}
	l.cfg.Encoding = "json"
	logger, err := l.cfg.Build()
	if err != nil {
		log.Println("failed to build logger")
		return
	}
	l.logger = logger.Sugar()
	l.logFile = file
}

/********* End of ZapSugarLogger *********/

func NewZapSugarLogger(env string) Logger {
	var cfg zap.Config
	if isProduction(env) {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.LevelKey = "level"

	logger, err := cfg.Build()
	if err != nil {
		log.Println("failed to build logger")
		return nil
	}

	var zapSugar ZapSugarLogger
	zapSugar.cfg = cfg
	zapSugar.logger = logger.Sugar()
	return &zapSugar
}
