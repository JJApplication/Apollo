package main

import (
	"runtime"

	"github.com/JJApplication/Apollo/cmd"

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
	cmd.ApolloCmd()
	initGlobalConfig()
	logger.Check()
	err := logger.InitLogger()
	if err != nil {
		logger.Logger.Error(logger.LoggerInitFailed)
		return
	}
	logger.CoreDump()
	// init recover
	defer func() {
		if r := recover(); r != nil {
			logger.LoggerSugar.Error(r)
			// add stack
			var stackBuf [1 << 16]byte
			stackLen := runtime.Stack(stackBuf[:], false)
			logger.LoggerSugar.Errorf("%s", stackBuf[:stackLen])
		}
	}()

	// init database
	initMongo()
	// init app manager
	initAPPManager()
	// init app discover
	initDiscoverManager()
	// init log manager
	initLogManager()
	// init task manager
	initTaskManager()
	// init background ticker
	initBackgroundJobs()
	// init background cron job
	initCronJobs()
	// init docker
	initDockerClient()
	// init NoEngine
	initNoEngineApps()
	//init env manager etcd client
	initEnvManager()
	// init repo manager
	initRepoManager()
	// only all manager init, can uds server be active
	initUDS()
	// init engine
	initEngine()
}
