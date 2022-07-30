/*
Project: Apollo apollo_init.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/app/docker_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/cron"
	"github.com/JJApplication/Apollo/database"
	"github.com/JJApplication/Apollo/engine"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/router/router_app"
	"github.com/JJApplication/Apollo/router/router_container"
	"github.com/JJApplication/Apollo/router/router_tasks"
	"github.com/JJApplication/Apollo/router/router_web"
	"github.com/JJApplication/Apollo/uds"
	"github.com/JJApplication/Apollo/utils"
)

const (
	APPName = "[Apollo]"
)

// 初始化运行时数据
func initGlobalConfig() {
	t := time.Now()
	err := config.InitGlobalConfig()
	if err != nil {
		fmt.Printf("%s init config failed: %s\n", APPName, err)
		return
	}
	fmt.Printf("%s load config in %dms\n", APPName, utils.TimeCalcUnix(t))
}

// 初始化数据库
func initMongo() {
	err := database.InitDBMongo()
	if err != nil {
		fmt.Printf("%s init mongo failed: %s\n", APPName, err.Error())
		database.MongoPing = false
		return
	}
	if err = database.Ping(); err != nil {
		fmt.Printf("%s ping mongo failed: %s\n", APPName, err.Error())
		database.MongoPing = false
		return
	}
	fmt.Println(APPName + " init mongo success")
}

func initAPPManager() {
	app_manager.InitAPPManager()
	app_manager.SaveToDB()
	app_manager.FirstLoad()
	logger.Logger.Info(APPName + " init APPManager done")
}

func initBackgroundJobs() {
	cron.InitBackgroundJobs()
	logger.Logger.Info(APPName + " init BackgroundJobs done")
}

func initEngine() {
	dirEngine := engine.NewEngine(&engine.EngineConfig{
		Host: config.ApolloConf.Server.Host,
		Port: config.ApolloConf.Server.Port,
	})
	dirEngine.Init()

	// load router
	router_app.Init(dirEngine.GetEngine())
	router_web.Init(dirEngine.GetEngine())
	router_tasks.Init(dirEngine.GetEngine())
	router_container.Init(dirEngine.GetEngine())

	err := dirEngine.RunServer()
	if err != nil && err != http.ErrServerClosed {
		logger.Logger.Error(APPName + " server start failed")
		logger.Logger.Error(fmt.Sprintf("%s %s", APPName, err.Error()))
		cron.InsureTickerExit()
		fmt.Println(APPName + " server start failed ⚠️")
		return
	}

	if err == http.ErrServerClosed {
		fmt.Println(APPName + " thanks for using ❤️️")
	}
}

func initUDS() {
	uds.Register()
}

func initDockerClient() {
	docker_manager.InitDockerClient()
}
