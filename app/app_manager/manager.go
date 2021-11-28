/*
Project: dirichlet manager.go
Created: 2021/11/18 by Landers
*/

package app_manager

import (
	"errors"
)

// 服务管理
// 启动停止脚本来自于配置数组 所有服务都位于统一的路径下

func GetApp(app string) (*App, error) {
	if Check(app) {
		a, _ := AppManagerMap.Load(app)
		return a.(*App), nil
	}

	return nil, errors.New(APPNotExist)
}

func Check(app string) bool {
	if _, ok := AppManagerMap.Load(app); ok {
		return true
	}

	return false
}

// Start app快速启动
func Start(app string) (bool, error) {
	if Check(app) {
		a, _ := AppManagerMap.Load(app)
		return a.(*App).Start()
	}

	return false, errors.New(APPNotExist)
}

// Stop app快速停止
func Stop(app string) (bool, error) {
	if Check(app) {
		a, _ := AppManagerMap.Load(app)
		return a.(*App).Stop()
	}

	return false, errors.New(APPNotExist)
}

// ReStart app快速重启
func ReStart(app string) (bool, error) {
	if Check(app) {
		a, _ := AppManagerMap.Load(app)
		return a.(*App).ReStart()
	}

	return false, errors.New(APPNotExist)
}
