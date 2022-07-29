/*
Project: Apollo manager.go
Created: 2021/11/18 by Landers
*/

package app_manager

import (
	"errors"
	"fmt"

	"github.com/JJApplication/Apollo/logger"
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

// GetApp 获取指定App
func GetApp(app string) (App, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		return a.(App), nil
	}

	return App{}, errors.New(APPNotExist)
}

// GetAllApp 获取所有App
func GetAllApp() ([]App, error) {
	var apps []App
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		apps = append(apps, value.(App))
		return true
	})
	return apps, nil
}

// Check 检查是否存在此App
func Check(app string) bool {
	if _, ok := APPManager.APPManagerMap.Load(app); ok {
		return true
	}

	return false
}

// Status 获取指定app状态
func Status(app string) (string, error) {
	if Check(app) {
		a, _ := APPManager.APPManagerMap.Load(app)
		b := a.(App)
		ok, err := b.Check()
		if err != nil {
			return StatusMap[toCode(err.Error())], err
		}
		if ok {
			return StatusMap[APPStatusOK], nil
		}
		return StatusMap[APPStatusStop], nil

	}

	return StatusMap[AppUnknown], errors.New(APPNotExist)
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
func StartAll() ([]string, error) {
	var e error
	var startList []string
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		app := value.(App)
		if ok, err := app.Start(); !ok {
			e = err
			startList = append(startList, fmt.Sprintf("[%s]: BAD", app.Meta.Name))
			logger.Logger.Error(fmt.Sprintf("%s %s start failed: %s", APPManagerPrefix, key, err.Error()))
		} else {
			startList = append(startList, fmt.Sprintf("[%s]: OK", app.Meta.Name))
			logger.Logger.Info(fmt.Sprintf("%s %s start success", APPManagerPrefix, key))
		}
		return true
	})

	if e != nil {
		return startList, errors.New("apps start has failure")
	}
	return startList, nil
}

// StopAll 停止所有服务
func StopAll() ([]string, error) {
	var e error
	var stopList []string
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		app := value.(App)
		if ok, err := app.Stop(); !ok {
			e = err
			stopList = append(stopList, fmt.Sprintf("[%s]: BAD", app.Meta.Name))
			logger.Logger.Error(fmt.Sprintf("%s %s stop failed: %s", APPManagerPrefix, key, err.Error()))
		} else {
			stopList = append(stopList, fmt.Sprintf("[%s]: OK", app.Meta.Name))
			logger.Logger.Info(fmt.Sprintf("%s %s stop success", APPManagerPrefix, key))
		}
		return true
	})

	if e != nil {
		return stopList, errors.New("apps stop has failure")
	}
	return stopList, nil
}

func StatusAll() ([]string, error) {
	var e error
	var statusList []string
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		app := value.(App)
		if ok, err := app.Check(); !ok {
			e = err
			statusList = append(statusList, fmt.Sprintf("[%s]: BAD", key))
			logger.Logger.Error(fmt.Sprintf("%s %s check failed: %s", APPManagerPrefix, key, err.Error()))
		} else {
			statusList = append(statusList, fmt.Sprintf("[%s]: OK", key))
			logger.Logger.Info(fmt.Sprintf("%s %s check success", APPManagerPrefix, key))
		}
		return true
	})

	if e != nil {
		return statusList, errors.New("apps check has failure")
	}
	return statusList, nil
}
