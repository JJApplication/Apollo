/*
Project: Apollo load_cf.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
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
			logger.LoggerSugar.Infof("store app [%s] config: %+v", k, v)
			APPManager.APPManagerMap.Store(k, App{Meta: v})
		}
	}

	// 未刷新保持缓存的map

	return ok
}

// ReloadManagerMap 运行时刷新数据
// 只做增量更新 数据的更新由AppSync任务完成
func ReloadManagerMap() error {
	apps, err := octopus_meta.AutoLoad()
	if err != nil {
		return err
	}
	logger.LoggerSugar.Infof("%s reload %d apps", APPManagerPrefix, len(apps))
	logger.LoggerSugar.Infof("%s reload map %+v", APPManagerPrefix, apps)
	// 存在app时跳过
	for k, v := range apps {
		app, ok := APPManager.APPManagerMap.Load(k)
		if ok && app.(App).Meta.Name == k {
			continue
		}
		if !ok {
			continue
		}

		APPManager.APPManagerMap.Store(k, App{Meta: v})
	}
	return nil
}
