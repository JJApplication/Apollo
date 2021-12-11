/*
Project: dirichlet cf_util.go
Created: 2021/11/22 by Landers
*/

package config

import (
	"sync"

	"github.com/landers1037/configen"
	"github.com/landers1037/dirichlet/utils"
)

// InitGlobalConfig 初始化全局配置文件
func InitGlobalConfig() error {
	DirichletConf.lock = new(sync.Mutex)
	return configen.ParseConfig(
		&DirichletConf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			GlobalConfigRoot,
			GlobalConfigFile))
}

// SaveGlobalConfig 持久化配置
func SaveGlobalConfig() error {
	return configen.SaveConfig(
		&DirichletConf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(), GlobalConfigRoot, GlobalConfigFile))
}
