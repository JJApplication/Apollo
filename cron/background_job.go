/*
Project: dirichlet background_job.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"fmt"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/logger"
)

// 启动时执行的轮询任务 用于随时刷新持久化数据
// 持久化数据用于恢复
const (
	Duration_DBSaver   = 60 * 60
	Duration_DBPersist = 60 * 60 * 24
	Duration_APPSync   = 60 * 60
	Duration_APPSyncDB = 60
	Duration_APPCheck  = 60 * 60
)

func InitBackgroundJobs() {
	AddJobDBSaver()
	AddJobDBPersist()
	AddJobAPPSync()
	AddJobAPPDumps()
	AddJobAPPCheck()
}

// AddJobDBSaver 数据库刷新
func AddJobDBSaver() {
	logger.Logger.Info("job: database sync start")
	des := ""
	AddTicker(Duration_DBSaver, "DBSaver", des, func() {
		app_manager.SaveToDB()
	})
}

// AddJobDBPersist 数据库内容持久化
func AddJobDBPersist() {
	logger.Logger.Info("job: database persist start")
	des := ""
	AddTicker(Duration_DBPersist, "DBPersist", des, func() {
		app_manager.Persist()
	})
}

// AddJobAPPSync app配置文件同步
// 同步配置文件到配置文件目录
func AddJobAPPSync() {
	logger.Logger.Info("job: app config sync start")
	des := ""
	AddTicker(Duration_APPSync, "AppConfigSync", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.Sync()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("job app config sync failed: %s", err.Error()))
			}
			return true
		})
	})
}

// AddJobAPPDumps 从缓存同步配置到mongo
// 用于同步配置参数和端口变量
func AddJobAPPDumps() {
	logger.Logger.Info("job: app runtime sync start")
	des := ""
	AddTicker(Duration_APPSyncDB, "AppRuntimeSync", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.SyncDB()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("job app runtime syncDB failed: %s", err.Error()))
			}
			return true
		})
	})
}

// AddJobAPPCheck app服务定时检查
func AddJobAPPCheck() {
	logger.Logger.Info("job: app check start")
	des := ""
	AddTicker(Duration_APPCheck, "AppChecker", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.Check()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("job app checker: %s check failed: %s try to restart", key, err.Error()))
				_, err = app.ReStart()
				logger.Logger.Warn(fmt.Sprintf("job app checker: %s restart result: %s", key, err.Error()))
			}
			return true
		})
	})
}
