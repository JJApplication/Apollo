/*
Project: dirichlet app_model_types.go
Created: 2021/11/27 by Landers
*/

package app_manager

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
)

func appScriptPath(app, c string) string {
	return filepath.Join(APPScriptsRoot, app, c)
}

func wrapWithCode(envs []string) []string {
	return append([]string{
		fmt.Sprintf("%s=%d", "APP_STATUS_ERR", APPStatusError),
		fmt.Sprintf("%s=%d", "APP_START_ERR", APPStatusStart),
		fmt.Sprintf("%s=%d", "APP_STOP_ERR", APPStatusStop),
		fmt.Sprintf("%s=%d", "APP_EXIT_ERR", APPStatusExit),
		fmt.Sprintf("%s=%d", "APP_RESTART_ERR", APPStatusRestart),
		fmt.Sprintf("%s=%d", "APP_KILL_ERR", APPStatusKilled),
		fmt.Sprintf("%s=%d", "APP_RUN_ERR", APPStatusRunning),
	}, envs...)
}

func (app *App) Start() (bool, error) {
	var ret int
	for _, c := range app.ManageCMD.Start {
		_, err := utils.CMDRun(wrapWithCode(app.RunData.Envs), appScriptPath(app.Name, c))
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("%s execute cmd (%s) faield: %s", APPManager, appScriptPath(app.Name, c), err.Error()))
			ret = toCode(err.Error())
			return false, errors.New(appCodeMap[ret])
		}
	}

	return true, nil
}

func (app *App) Stop() (bool, error) {
	return true, nil
}

func (app *App) ReStart() (bool, error) {
	return true, nil
}

// ForceKill 查找进程树 全部强制kill
func (app *App) ForceKill() (bool, error) {
	return true, nil
}

// PostTodo 启动前的操作，生成环境变量或动态配置到app中
func (app *App) PostTodo() *App {
	return app
}

// Check 状态检查
func (app *App) Check() (bool, error) {
	return true, nil
}

// BackUp 备份服务
// 当前的备份比较简单 打包整个服务到tar中，去除所有的日志文件和缓存文件
func (app *App) BackUp() (bool, error) {
	return true, nil
}

// Reload 重载配置文件
func (app *App) Reload() (bool, error) {
	return true, nil
}

// Sync 同步配置文件
func (app *App) Sync() (bool, error) {
	return true, nil
}

// Info 获取app的基础信息
func (app *App) Info() (bool, error) {
	return true, nil
}

// Dump 安全的保存运行态数据
func (app *App) Dump() (bool, error) {
	return true, nil
}

// ToJSON 导出为json字符串
func (app *App) ToJSON() (string, error) {
	return "", nil
}
