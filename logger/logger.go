package logger

import (
	"log"
	"strings"

	"github.com/harikrishnanum/go-log-interface/config"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	// Add more methods as needed
	OpenLogFile(file string)
	CloseLogFile()
}

func isProduction(env string) bool {
	switch strings.ToLower(env) {
	case "dev", "development", "test", "stage", "staging":
		return false
	case "prod", "production":
		return true
	default:
		log.Println("logger environment not supported, defaulting to production")
		return true
	}
}

var Log Logger // Global logger

func InitLogger(logger, env string) {
	switch logger {
	case "zapsugar":
		Log = NewZapSugarLogger(env)
	case "logrus":
		Log = NewLogrusLogger(env)
	default:
		Log = NewLogrusLogger(env)
	}
	if Log == nil {
		log.Printf("Failed to initialize %s logger", config.Conf.Logger)
		return
	}
	log.Printf("Initialized %s logger", config.Conf.Logger)
}
