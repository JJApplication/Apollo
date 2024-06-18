/*
Project: Apollo app_model_types.go
Created: 2021/11/27 by Landers
*/

package app_manager

import (
	"errors"
	"fmt"
	"github.com/JJApplication/Apollo/app/noengine_manager"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/octopus_meta"
)

func appScriptPath(app, c string) string {
	return filepath.Join(config.ApolloConf.APPManager, app, c)
}

func wrapWithCode(envs []string) []string {
	return append([]string{
		fmt.Sprintf("%s=%d", "APP_STATUS_OK", APPStatusOK),
		fmt.Sprintf("%s=%d", "APP_STATUS_ERR", APPStatusError),
		fmt.Sprintf("%s=%d", "APP_START_ERR", APPStatusStart),
		fmt.Sprintf("%s=%d", "APP_STOP_ERR", APPStatusStop),
		fmt.Sprintf("%s=%d", "APP_EXIT_ERR", APPStatusExit),
		fmt.Sprintf("%s=%d", "APP_RESTART_ERR", APPStatusRestart),
		fmt.Sprintf("%s=%d", "APP_KILL_ERR", APPStatusKilled),
		fmt.Sprintf("%s=%d", "APP_RUN_ERR", APPStatusRunning),

		// 运行时服务路径
		fmt.Sprintf("%s=%s", "SERVICE_ROOT", config.ApolloConf.ServiceRoot),
		fmt.Sprintf("%s=%s", "APP_ROOT", config.ApolloConf.APPRoot),
		fmt.Sprintf("%s=%s", "APP_LOG", config.ApolloConf.APPLogDir),
		fmt.Sprintf("%s=%s", "APP_PID", config.ApolloConf.APPPidDir),
		fmt.Sprintf("%s=%s", "APP_TMP", config.ApolloConf.APPTmpDir),
	}, envs...)
}

// 生成运行时所需的环境变量
func attachEnvs(app *App) []string {
	var envs []string
	if app.Meta.Name != "" {
		envs = append(app.Meta.RunData.Envs, fmt.Sprintf("APP=%s", app.Meta.Name))
	}

	path := os.Getenv("PATH")
	if path != "" {
		envs = append(envs, fmt.Sprintf("PATH=%s", path+":/usr/local/bin"))
	}

	return wrapWithCode(envs)
}

// 加载固定端口时使用
func attachEnvsSp(app *App) []string {
	var envs = attachEnvs(app)
	if len(app.Meta.RunData.Ports) > 0 {
		var ports []string
		for _, p := range app.Meta.RunData.Ports {
			ports = append(ports, strconv.Itoa(p))
		}
		envs = append(envs, fmt.Sprintf("PORTS=%s", strings.Join(ports, " ")))
	}

	return envs
}

// 生成运行时所需的端口
// 通过appManager中的ports来去重
func attachEnvsWithPorts(app *App) []string {
	// 仅重试10次
	var s string
	var envs = attachEnvs(app)
	for i := 0; i < 10; i++ {
		p := utils.RandomPort()
		if APPManager.checkPorts(p) {
			s = fmt.Sprintf("PORTS=%d", p)
			// 记录port到app运行时 在启动失败后从manager中删除
			app.Meta.RunData.Ports = []int{p}
			APPManager.addPorts(p)
			app.ClonePorts()
			app.Dump()
			break
		}
	}

	envs = append(envs, s)
	return envs
}

// Start 启动服务
// 每次启动前应该强制校验 先停止服务
// rundata中的ports是随机分配的，只能在START中生效
func (app *App) Start() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}

	if app.NoEngineCheck() {
		return app.NoEngineStart()
	}

	var ret int
	if ok, _ := app.Check(); ok {
		return true, nil
	}
	if !app.DepCheck() {
		logger.LoggerSugar.Errorf("%s [%s] start failed, app's dependencies is not ready", APPManagerPrefix, app.Meta.Name)
		return false, errors.New("app's dependencies is not ready")
	}
	// 判断是否需要随机端口运行
	if app.Meta.RunData.RandomPort {
		_, err := utils.CMDRun(attachEnvsWithPorts(app), appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start))
		if err != nil {
			if len(app.Meta.RunData.Ports) > 0 {
				APPManager.delPorts(app.Meta.RunData.Ports[0])
				app.ClearPorts()
			}

			logger.LoggerSugar.Error("%s execute cmd (%s) failed: %s", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start), err.Error())
			ret = toCode(err.Error())
			return false, errors.New(errCode(ret))
		}
		app.startByApollo()
		logger.LoggerSugar.Infof("%s execute cmd (%s) success", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start))
		return true, err
	}

	_, err := utils.CMDRun(attachEnvsSp(app), appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start))
	if err != nil {
		logger.LoggerSugar.Errorf("%s execute cmd (%s) failed: %s", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start), err.Error())
		ret = toCode(err.Error())
		return false, errors.New(errCode(ret))
	}
	app.startByApollo()
	logger.LoggerSugar.Infof("%s execute cmd (%s) success", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Start))
	return true, nil
}

func (app *App) Stop() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}
	if app.NoEngineCheck() {
		return app.NoEngineStop()
	}
	var ret int
	_, err := utils.CMDRun(attachEnvs(app), appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Stop))
	if err != nil {
		// 停止失败时 保留原有的数据
		logger.LoggerSugar.Errorf("%s execute cmd (%s) failed: %s", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Stop), err.Error())
		ret = toCode(err.Error())
		return false, errors.New(errCode(ret))
	}
	// 停止成功时 清空保留的ports
	// 仅清除动态端口
	if app.Meta.RunData.RandomPort && len(app.Meta.RunData.Ports) > 0 {
		APPManager.delPorts(app.Meta.RunData.Ports[0])
		app.ClearPorts()
	}

	app.stopByApollo()
	logger.LoggerSugar.Infof("%s execute cmd (%s) success", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Stop))
	return true, nil
}

func (app *App) stopByApollo() {
	app.Meta.Runtime.StopOperation = true
	app.Dump()
}

func (app *App) startByApollo() {
	app.Meta.Runtime.StopOperation = false
	app.Dump()
}

func (app *App) ReStart() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}
	if app.NoEngineCheck() {
		return app.NoEngineReStart()
	}
	if app.Meta.Runtime.StopOperation {
		return true, nil
	}
	ok, err := app.Stop()
	if !ok || err != nil {
		return false, err
	}

	ok, err = app.Start()
	if !ok || err != nil {
		return false, err
	}

	return true, nil
}

// ForceKill 查找进程树 全部强制kill
func (app *App) ForceKill() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}
	return true, nil
}

// PostTodo 启动前的操作，生成环境变量或动态配置到app中
func (app *App) PostTodo() *App {
	return app
}

// Check 状态检查
func (app *App) Check() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}
	if app.NoEngineCheck() {
		return app.NoEngineStatus()
	}

	var ret int
	_, err := utils.CMDRun(attachEnvs(app), appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Check))
	if err != nil {
		logger.LoggerSugar.Errorf("%s execute cmd (%s) failed: %s", APPManagerPrefix, appScriptPath(app.Meta.Name, app.Meta.ManageCMD.Check), err.Error())
		ret = toCode(err.Error())
		return false, errors.New(errCode(ret))
	}

	return true, nil
}

// ClearPorts 删除运行时环境
func (app *App) ClearPorts() {
	app.Meta.RunData.Ports = []int{}
	app.Meta.Runtime.Ports = []int{}
}

// ClonePorts 同步缓存中的随机端口组到mongo
func (app *App) ClonePorts() {
	SavePort(app.Meta.Name, app.Meta.RunData.Ports)
}

// BackUp 备份服务
// 当前的备份比较简单 打包整个服务到tar中，去除所有的日志文件和缓存文件
func (app *App) BackUp() (bool, error) {
	if !app.CheckAppReleaseStatus() {
		return true, nil
	}
	return true, nil
}

// Reload 重载配置文件
func (app *App) Reload() (bool, error) {
	// 默认线程安全
	err := loadFromApp(app.Meta.Name)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Sync 从本地文件同步模型数据
// 能够同步的数据仅包括非运行时数据
func (app *App) Sync() (bool, error) {
	lock := sync.Mutex{}
	lock.Lock()
	// 运行时数据跳过
	var appClone App
	appClone = *app
	if len(app.Meta.RunData.Ports) > 0 {
		appClone.Meta.RunData.Ports = app.Meta.RunData.Ports
	} else {
		appClone.Meta.RunData.Ports = []int{}
	}

	var cf octopus_meta.App
	cf, err := octopus_meta.LoadApp(app.Meta.Name)
	if err != nil {
		return false, errors.New("can't load from config file, " + err.Error())
	}

	if !cf.Validate() {
		return false, errors.New("can't load from config file, validate failed")
	}
	appClone.Meta.ID = cf.ID
	appClone.Meta.Type = cf.Type
	appClone.Meta.ReleaseStatus = cf.ReleaseStatus
	appClone.Meta.EngDes = cf.EngDes
	appClone.Meta.CHSDes = cf.CHSDes
	appClone.Meta.Link = cf.Link
	appClone.Meta.ManageCMD = cf.ManageCMD
	appClone.Meta.Meta = cf.Meta
	appClone.Meta.ResourceLimit = cf.ResourceLimit
	APPManager.APPManagerMap.Store(app.Meta.Name, appClone)
	lock.Unlock()
	return true, nil
}

// Info 获取app的基础信息
func (app *App) Info() interface{} {
	return app.Meta
}

// Dump 安全的保存运行态数据到Map中
func (app *App) Dump() (bool, error) {
	APPManager.APPManagerMap.Store(app.Meta.Name, *app)
	return true, nil
}

// SyncDB 安全的保存运行态数据到DB中
func (app *App) SyncDB() (bool, error) {
	SaveRuntimeData(*app)
	return true, nil
}

// ToJSON 导出为json字符串
func (app *App) ToJSON() string {
	return utils.PrettyJson(app)
}

// CheckAppReleaseStatus 检查app发布状态 未发布时不进行操作
func (app *App) CheckAppReleaseStatus() bool {
	if app.Meta.ReleaseStatus == octopus_meta.Published {
		return true
	}
	logger.LoggerSugar.Warnf("%s [%s] releaseStatus is %s, skip operation", APPManagerPrefix, app.Meta.Name, app.Meta.ReleaseStatus)
	return false
}

// DepCheck 运行时依赖检查
// 依赖异常时 无法成功启动 停止不受影响
// 不存在依赖时跳过
func (app *App) DepCheck() bool {
	deps := app.Meta.RunData.RunDep
	if deps == nil || len(deps) <= 0 {
		return true
	}
	var apps []App
	for _, dep := range deps {
		if ok, a := APPManager.hasApp(dep); ok {
			apps = append(apps, a)
		}
	}

	// check
	for _, depApp := range apps {
		ok, err := depApp.Check()
		if !ok || err != nil {
			return false
		}
	}
	return true
}

// NoEngineCheck 检查是否为NoEngine服务
// NoEngine服务走单独的操作流程
func (app *App) NoEngineCheck() bool {
	return app.Meta.Type == noengine_manager.NoEngine
}

func (app *App) NoEngineStatus() (bool, error) {
	// 未创建容器时返回错误
	if noengine_manager.NoEngineAPPID(app.Meta.Name) == "" {
		logger.LoggerSugar.Errorf("%s [%s] status failed, error: container not exist", APPManagerPrefix, app.Meta.Name)
		return false, errors.New("container not exist")
	}
	// 判断容器状态
	status := noengine_manager.StatusNoEngineApp(app.Meta.Name)
	if status != "running" {
		return false, nil
	}
	return true, nil
}

func (app *App) NoEngineStart() (bool, error) {
	// 未创建容器时先创建
	if noengine_manager.NoEngineAPPID(app.Meta.Name) == "" {
		if err := noengine_manager.StartNoEngineApp(app.Meta.Name); err != nil {
			logger.LoggerSugar.Errorf("%s [%s] start failed, error: %s", APPManagerPrefix, app.Meta.Name, err.Error())
			return false, err
		}
		return true, nil
	}
	if err := noengine_manager.ResumeNoEngineApp(app.Meta.Name); err != nil {
		logger.LoggerSugar.Errorf("%s [%s] start failed, error: %s", APPManagerPrefix, app.Meta.Name, err.Error())
		return false, err
	}
	logger.LoggerSugar.Infof("%s [%s] start success", APPManagerPrefix, app.Meta.Name)
	return true, nil
}

func (app *App) NoEngineStop() (bool, error) {
	// 未创建 警告但不创建
	if noengine_manager.NoEngineAPPID(app.Meta.Name) == "" {
		logger.LoggerSugar.Warnf("%s [%s] stop failed, container not exist", APPManagerPrefix, app.Meta.Name)
		return false, errors.New("container not exist")
	}
	if err := noengine_manager.StopNoEngineApp(app.Meta.Name); err != nil {
		logger.LoggerSugar.Errorf("%s [%s] stop failed, error: %s", APPManagerPrefix, app.Meta.Name, err.Error())
		return false, err
	}
	logger.LoggerSugar.Infof("%s [%s] stop success", APPManagerPrefix, app.Meta.Name)
	return true, nil
}

func (app *App) NoEngineReStart() (bool, error) {
	// 未创建 警告但不创建
	if noengine_manager.NoEngineAPPID(app.Meta.Name) == "" {
		logger.LoggerSugar.Warnf("%s [%s] restart failed, container not exist", APPManagerPrefix, app.Meta.Name)
		return false, errors.New("container not exist")
	}
	if err := noengine_manager.StopNoEngineApp(app.Meta.Name); err != nil {
		logger.LoggerSugar.Errorf("%s [%s] stop failed, error: %s", APPManagerPrefix, app.Meta.Name, err.Error())
		return false, err
	}
	if err := noengine_manager.ResumeNoEngineApp(app.Meta.Name); err != nil {
		logger.LoggerSugar.Errorf("%s [%s] start failed, error: %s", APPManagerPrefix, app.Meta.Name, err.Error())
		return false, err
	}
	logger.LoggerSugar.Infof("%s [%s] restart success", APPManagerPrefix, app.Meta.Name)
	return true, nil
}
