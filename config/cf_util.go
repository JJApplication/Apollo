/*
Project: Apollo cf_util.go
Created: 2021/11/22 by Landers
*/

package config

import (
	"sync"

	"github.com/JJApplication/Apollo/utils"
	"github.com/landers1037/configen"
)

// InitGlobalConfig 初始化全局配置文件
func InitGlobalConfig() error {
	ApolloConf.lock = new(sync.Mutex)
	return configen.ParseConfig(
		&ApolloConf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(),
			GlobalConfigRoot,
			GlobalConfigFile))
}

// SaveGlobalConfig 持久化配置
func SaveGlobalConfig() error {
	return configen.SaveConfig(
		&ApolloConf,
		configen.Pig,
		utils.CalDir(
			utils.GetAppDir(), GlobalConfigRoot, GlobalConfigFile))
}
