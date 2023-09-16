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

// UpdateConfig 更新运行时配置
// 传入的config是一个配置的子集
func UpdateConfig(cf RuntimeConf) {
	ApolloConf.lock.Lock()
	ApolloConf.ServiceRoot = cf.ServiceRoot
	ApolloConf.APPRoot = cf.APPRoot
	ApolloConf.APPManager = cf.APPManager
	ApolloConf.APPCacheDir = cf.APPCacheDir
	ApolloConf.APPLogDir = cf.APPLogDir
	ApolloConf.APPTmpDir = cf.APPTmpDir
	ApolloConf.APPBackUp = cf.APPBackUp

	ApolloConf.Server.AuthExpire = cf.AuthExpire
	ApolloConf.Server.UICache = cf.UICache
	ApolloConf.Server.UICacheTime = cf.UICacheTime
	ApolloConf.Server.UIRouter = cf.UIRouter

	ApolloConf.Log.LogFile = cf.LogFile
	ApolloConf.Log.EnableStack = cf.EnableStack
	ApolloConf.Log.EnableFunction = cf.EnableFunction
	ApolloConf.Log.EnableCaller = cf.EnableCaller
	defer ApolloConf.lock.Unlock()
}
