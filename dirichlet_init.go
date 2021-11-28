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
	err := app_manager.LoadManagerCf()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("init APPConfig failed: %s", err.Error()))
	}

	logger.Logger.Info("init APPConfig done")
}

func initEngine() {
	dirEngine := engine.NewEngine(&engine.EngineConfig{})
	dirEngine.Init()
	err := dirEngine.Run()

	if err != nil {
		logger.Logger.Error("[Dirichlet] server start failed")
		return
	}
}