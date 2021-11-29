package main

import (
	"github.com/landers1037/dirichlet/logger"
)

func main() {
	initGlobalConfig()
	logger.Check()
	err := logger.InitLogger()
	if err != nil {
		logger.Logger.Error(logger.LoggerInitFailed)
		return
	}

	// init app manager
	initAPPManager()
	// init engine
	initEngine()
}
