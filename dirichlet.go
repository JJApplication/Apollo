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

	// init database
	initMongo()
	// init app manager
	initAPPManager()
	// init background ticker
	initBackgroundJobs()
	// init background cron job

	// only all manager init, can uds server be active
	initUDS()
	// init engine
	initEngine()
}
