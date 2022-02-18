/*
Project: dirichlet middleware_cf.go
Created: 2022/2/17 by Landers
*/

package engine

import (
	"fmt"
	"os"
	"time"

	"github.com/landers1037/configen"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
)

// 加载中间件的配置文件
// 分离路由和主程序的配置使逻辑更加清晰

type MiddleWareConfig struct {
	Name   string   `json:"name"`
	Mode   string   `json:"mode"`
	Active bool     `json:"active"`
	Urls   []string `json:"urls"`
}

const (
	GlobalMode       = "global"
	RouteMode        = "route"
	MiddleConfigRoot = "conf"
	MiddleConfigFile = "middleware.pig"
)

// LoadMiddleWareConfig 同步的加载配置
func LoadMiddleWareConfig() []MiddleWareConfig {
	if err := initEmptyConfig(); err != nil {
		logger.Logger.Warn(MiddleWare + " failed to init empty middleware config.")
	}
	var c []MiddleWareConfig
	t := time.Now()
	if err := configen.ParseConfig(&c, configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			MiddleConfigRoot,
			MiddleConfigFile)); err != nil {
		logger.Logger.Error(MiddleWare + " failed to parse middleware config.")
		logger.Logger.Warn(MiddleWare + " using default middleware config.")
	} else {
		logger.Logger.Info(fmt.Sprintf("%s load middleware config in %dms.", MiddleWare, utils.TimeCalcUnix(t)))
	}

	return c
}

// SaveMiddleWareConfig 同步的保存配置
func SaveMiddleWareConfig() error {
	return configen.SaveConfig(&PreInjectMiddle, configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			MiddleConfigRoot,
			MiddleConfigFile))
}

// 不存在配置文件时会自动生成基于当前接口的配置文件
func initEmptyConfig() error {
	if _, err := os.Stat(utils.CalDir(
		utils.GetAppDir(),
		MiddleConfigRoot,
		MiddleConfigFile)); os.IsNotExist(err) {
		logger.Logger.Warn(MiddleWare + " no config exists, using default middleware config.")
		return configen.SaveConfig(&DefaultMiddleWare,
			configen.Pig,
			utils.CalDir(
				utils.GetAppDir(),
				MiddleConfigRoot,
				MiddleConfigFile))
	}

	return nil
}

// 重载中间件配置，会调用Linux的专有特性fork子进程
// todo
