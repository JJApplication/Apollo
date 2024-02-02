/*
Project: Apollo middleware_cf.go
Created: 2022/2/17 by Landers
*/

package middleware

import (
	"time"

	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/landers1037/configen"
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
		logger.LoggerSugar.Warnf("%s failed to init empty middleware config.", MiddleWare)
	}
	var c []MiddleWareConfig
	t := time.Now()
	if err := configen.ParseConfig(&c, configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			MiddleConfigRoot,
			MiddleConfigFile)); err != nil {
		logger.LoggerSugar.Errorf("%s failed to parse middleware config: %s", MiddleWare, err.Error())
		logger.LoggerSugar.Warnf("%s using default middleware config.", MiddleWare)
	} else {
		logger.LoggerSugar.Infof("%s load middleware config in %dms.", MiddleWare, utils.TimeCalcUnix(t))
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
	if utils.FileNotExist(utils.CalDir(
		utils.GetAppDir(),
		MiddleConfigRoot,
		MiddleConfigFile)) {
		logger.LoggerSugar.Warnf("%s no config exists, using default middleware config.", MiddleWare)
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
