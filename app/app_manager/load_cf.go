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
	"github.com/gookit/goutil/reflects"
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
			logger.LoggerSugar.Infof("store app [%s] config", k)
			APPManager.APPManagerMap.Store(k, App{Meta: v})
		}
	}

	// 未刷新保持缓存的map

	return ok
}

// ReloadManagerMap 运行时刷新数据
// 只做增量更新 数据的更新由AppSync任务完成
// 以octopus模型中注册的微服务为准，即新增octopus则注册服务，不存在的服务即被卸载的服务->删除
func ReloadManagerMap() error {
	apps, err := octopus_meta.AutoLoad()
	if err != nil {
		return err
	}
	logger.LoggerSugar.Infof("%s reload %d apps", APPManagerPrefix, len(apps))
	logger.LoggerSugar.Infof("%s reload map %+v", APPManagerPrefix, apps)

	var oldKeys []string
	// 先执行合并任务
	APPManager.APPManagerMap.Range(func(key, value any) bool {
		oldKeys = append(oldKeys, key.(string))
		return true
	})
	for _, key := range oldKeys {
		if _, ok := apps[key]; !ok {
			APPManager.APPManagerMap.Delete(key)
		}
	}

	for k, v := range apps {
		// 存在app时跳过
		app, ok := APPManager.APPManagerMap.Load(k)
		if ok && app.(App).Meta.Name == k {
			continue
		}
		// 不存在且为空对象时跳过
		if !ok && reflects.IsEqual(v, octopus_meta.App{}) {
			continue
		}

		APPManager.APPManagerMap.Store(k, App{Meta: v})
	}
	return nil
}
