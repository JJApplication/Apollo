package main

import (
	"github.com/JJApplication/Apollo/logger"
)

// @title Apollo ServiceGroup
// @version 1.0
// @description Apollo服务接口文档
// @termsOfService http://service.renj.io/terms
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath http://service.renj.io/api
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

	// init docker
	initDockerClient()
	// only all manager init, can uds server be active
	initUDS()
	// init engine
	initEngine()
}
