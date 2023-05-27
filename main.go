package main

import "github.com/harikrishnanum/go-log-interface/logger"

func main() {
	logger.Log.Info("Hello World")
	logger.Log.Infof("Hello %s", "World")
	logger.Log.Error("Hello World")
	logger.Log.Errorf("Hello %s", "World")
}
