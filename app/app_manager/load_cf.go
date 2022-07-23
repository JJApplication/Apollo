/*
Project: Apollo load_cf.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"fmt"
	"os"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/octopus_meta"
)

// LoadManagerCf 加载所有配置文件到全局的字典中
func LoadManagerCf() error {
	// 保证读取到配置后再刷新字典
	AppRoot := os.Getenv("APP_ROOT")
	if AppRoot == "" {
		os.Setenv("APP_ROOT", config.ApolloConf.APPRoot)
	}

	tm, ok := octopus_meta.AutoLoad()
	// 每次刷新
	if ok == nil {
		APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			APPManager.APPManagerMap.Delete(key)
			return true
		})

		for k, v := range tm {
			logger.Logger.Info(fmt.Sprintf("store app [%s] config: %+v", k, v))
			APPManager.APPManagerMap.Store(k, App{Meta: v})
		}
	}

	// 未刷新保持缓存的map

	return ok
}
