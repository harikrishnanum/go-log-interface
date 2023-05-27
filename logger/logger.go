package logger

import (
	"log"

	"github.com/harikrishnanum/go-log-interface/config"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	// Add more methods as needed
}

var Log Logger // Global logger

func init() {
	switch config.Conf.Logger {
	case "zapsugar":
		Log = NewZapLogger()
	case "logrus":
		Log = NewLogrusLogger()
	default:
		Log = NewZapLogger()
	}
	if Log == nil {
		log.Printf("Failed to initialize %s logger", config.Conf.Logger)
		return
	}
	log.Printf("Initialized %s logger", config.Conf.Logger)
}
