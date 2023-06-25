package main

import (
	"github.com/harikrishnanum/go-log-interface/config"
	"github.com/harikrishnanum/go-log-interface/logger"
)

func main() {
	// Log to console
	logger.InitLogger(config.Conf.Logger, config.Conf.Env)
	logger.Log.Info("Hello World")
	logger.Log.Infof("Hello %s", "World")

	// Log to file
	// Initialize logger first, here we already initialised it...
	// logger.InitLogger(config.Conf.Logger, config.Conf.Env)
	logger.Log.OpenLogFile("logs.log")
	defer logger.Log.CloseLogFile()
	logger.Log.Error("Hello World")
	logger.Log.Errorf("Hello %s", "World")
}
