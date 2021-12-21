/*
Project: dirichlet dirichlet_init.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"fmt"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/config"
	"github.com/landers1037/dirichlet/cron"
	"github.com/landers1037/dirichlet/database"
	"github.com/landers1037/dirichlet/engine"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/router/router_app"
	"github.com/landers1037/dirichlet/router/router_web"
	"github.com/landers1037/dirichlet/uds"
)

// 初始化运行时数据
func initGlobalConfig() {
	err := config.InitGlobalConfig()
	if err != nil {
		fmt.Printf("[Dirichlet] init config failed %s\n", err)
		return
	}
}

// 初始化数据库
func initMongo() {
	err := database.InitDBMongo()
	if err != nil {
		fmt.Printf("[Dirichlet] init mongo failed %s\n", err)
		return
	}
}

func initAPPManager() {
	app_manager.InitAPPManager()
	app_manager.SaveToDB()
	logger.Logger.Info("init APPManager done")
}

func initBackgroundJobs() {
	cron.InitBackgroundJobs()
	logger.Logger.Info("init BackgroundJobs done")
}

func initEngine() {
	dirEngine := engine.NewEngine(&engine.EngineConfig{
		Host: config.DirichletConf.Server.Host,
		Port: config.DirichletConf.Server.Port,
	})
	dirEngine.Init()

	// load router
	router_app.Init(dirEngine.GetEngine())
	router_web.Init(dirEngine.GetEngine())

	err := dirEngine.Run()
	if err != nil {
		logger.Logger.Error("[Dirichlet] server start failed")
		logger.Logger.Error(err.Error())
		cron.InsureTickerExit()
		return
	}
}

func initUDS() {
	uds.Register()
}
