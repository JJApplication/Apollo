/*
Project: dirichlet manager.go
Created: 2021/11/18 by Landers
*/

package app_manager

import (
	"errors"
	"fmt"

	"github.com/landers1037/dirichlet/logger"
)

// 服务管理
// 启动停止脚本来自于配置数组 所有服务都位于统一的路径下

// InitAPPManager 初始化app manager
func InitAPPManager() {
	err := LoadManagerCf()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("init APPManager failed: %s", err.Error()))
	}
}

func GetApp(app string) (App, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		return a.(App), nil
	}

	return App{}, errors.New(APPNotExist)
}

func Check(app string) bool {
	if _, ok := APPManager.APPManagerMap.Load(app); ok {
		return true
	}

	return false
}

// Start app快速启动
func Start(app string) (bool, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		b := a.(App)
		return b.Start()
	}

	return false, errors.New(APPNotExist)
}

// Stop app快速停止
func Stop(app string) (bool, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		b := a.(App)
		return b.Stop()
	}

	return false, errors.New(APPNotExist)
}

// ReStart app快速重启
func ReStart(app string) (bool, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		b := a.(App)
		return b.ReStart()
	}

	return false, errors.New(APPNotExist)
}

// StartAll 启动所有服务
// 服务的启动可以异步 并且不受其他服务报错的影响
func StartAll() error {
	var e error
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		app := value.(App)
		if ok, err := app.Start(); !ok {
			e = err
			logger.Logger.Error(fmt.Sprintf("%s %s start failed: %s", APPManagerPrefix, key, err.Error()))
		} else {
			logger.Logger.Info(fmt.Sprintf("%s %s start success", APPManagerPrefix, key))
		}
		return true
	})

	if e != nil {
		return errors.New("apps start has failure")
	}
	return nil
}

// StopAll 停止所有服务
func StopAll() error {
	var e error
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		app := value.(App)
		if ok, err := app.Stop(); !ok {
			e = err
			logger.Logger.Error(fmt.Sprintf("%s %s stop failed: %s", APPManagerPrefix, key, err.Error()))
		} else {
			logger.Logger.Info(fmt.Sprintf("%s %s stop success", APPManagerPrefix, key))
		}
		return true
	})

	if e != nil {
		return errors.New("apps stop has failure")
	}
	return nil
}
