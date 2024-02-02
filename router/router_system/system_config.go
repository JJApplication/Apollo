/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"fmt"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
)

// 配置文件处理

func GetConfig(c *gin.Context) {
	router.Response(c, config.RuntimeConf{
		ServiceRoot:      config.ApolloConf.ServiceRoot,
		APPRoot:          config.ApolloConf.APPRoot,
		APPManager:       config.ApolloConf.APPManager,
		APPCacheDir:      config.ApolloConf.APPCacheDir,
		APPLogDir:        config.ApolloConf.APPLogDir,
		APPTmpDir:        config.ApolloConf.APPTmpDir,
		APPBackUp:        config.ApolloConf.APPBackUp,
		EnableStack:      config.ApolloConf.Log.EnableStack,
		EnableFunction:   config.ApolloConf.Log.EnableFunction,
		EnableCaller:     config.ApolloConf.Log.EnableCaller,
		LogFile:          config.ApolloConf.Log.LogFile,
		UICache:          config.ApolloConf.Server.UICache,
		UICacheTime:      config.ApolloConf.Server.UICacheTime,
		UIRouter:         config.ApolloConf.Server.UIRouter,
		AuthExpire:       config.ApolloConf.Server.AuthExpire,
		PID:              utils.GetRuntimePID(),
		Port:             config.ApolloConf.Server.Port,
		UDS:              config.ApolloConf.Server.Uds,
		Mongo:            fmt.Sprintf("%s@%s", config.ApolloConf.DB.Mongo.Name, config.ApolloConf.DB.Mongo.URL),
		DockerApi:        config.ApolloConf.CI.DockerHost,
		DockerApiVersion: config.ApolloConf.CI.DockerAPIVersion,
		Goroutines:       runtime.NumGoroutine(),
		MaxProcs:         runtime.NumCPU(),
	}, true)
}

// UpdateConfig 传入后台的数据处理后更新到全局配置中
// 因为字段不一致 所以单独定义模型
func UpdateConfig(c *gin.Context) {
	var rcf config.RuntimeConf
	if err := c.BindJSON(&rcf); err != nil {
		router.Response(c, false, false)
		return
	}

	if (reflect.DeepEqual(rcf, config.RuntimeConf{})) {
		router.Response(c, false, false)
		return
	}

	config.UpdateConfig(rcf)
	router.Response(c, true, true)
}

func SaveConfig(c *gin.Context) {
	if err := config.SaveGlobalConfig(); err != nil {
		router.Response(c, false, false)
		return
	}
	router.Response(c, true, true)
}

// ReloadConfig todo
func ReloadConfig(c *gin.Context) {
	router.Response(c, true, true)
}
