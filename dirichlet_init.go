/*
Project: dirichlet dirichlet_init.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"fmt"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/config"
	"github.com/landers1037/dirichlet/engine"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/router/router_app"
)

// 初始化运行时数据
func initGlobalConfig() {
	err := config.InitGlobalConfig()
	if err != nil {
		fmt.Printf("[Dirichlet] init config failed %s\n", err)
		return
	}
}

func initAPPManager() {
	app_manager.InitAPPManager()
	logger.Logger.Info("init APPManager done")
}

func initEngine() {
	dirEngine := engine.NewEngine(&engine.EngineConfig{
		Host: config.DirichletConf.Server.Host,
		Port: config.DirichletConf.Server.Port,
	})
	dirEngine.Init()

	// load router
	router_app.Init(dirEngine.GetEngine())

	err := dirEngine.Run()
	if err != nil {
		logger.Logger.Error("[Dirichlet] server start failed")
		return
	}
}
