/*
Project: Apollo background_job.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"fmt"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
)

// 启动时执行的轮询任务 用于随时刷新持久化数据
// 持久化数据用于恢复
const (
	DurationDbSaver   = 60 * 60
	DurationDbPersist = 60 * 60 * 24
	DurationAppsync   = 60 * 60
	DurationAppSyncdb = 60
	DurationAppCheck  = 60 * 60
	DurationLogRotate = 7 * 60 * 60 * 24
)

func InitBackgroundJobs() {
	AddJobDBSaver()
	AddJobDBPersist()
	AddJobAPPSync()
	AddJobAPPDumps()
	AddJobAPPCheck()
	AddJobLogRotate()
}

// AddJobDBSaver 数据库刷新
func AddJobDBSaver() {
	logger.Logger.Info("job: database sync start")
	des := "同步刷新微服务信息到数据库"
	AddTicker(DurationDbSaver, "DBSaver", des, func() {
		app_manager.SaveToDB()
	})
}

// AddJobDBPersist 数据库内容持久化
func AddJobDBPersist() {
	logger.Logger.Info("job: database persist start")
	des := "数据库信息持久化存储"
	AddTicker(DurationDbPersist, "DBPersist", des, func() {
		app_manager.Persist()
	})
}

// AddJobAPPSync app配置文件同步
// 同步配置文件到配置文件目录
func AddJobAPPSync() {
	logger.Logger.Info("job: app config sync start")
	des := "同步存储微服务模型文件"
	AddTicker(DurationAppsync, "AppConfigSync", des, func() {
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
	des := "同步存储缓存数据到数据库"
	AddTicker(DurationAppSyncdb, "AppRuntimeSync", des, func() {
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
	des := "微服务状态定时检查"
	AddTicker(DurationAppCheck, "AppChecker", des, func() {
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

// AddJobLogRotate 日志裁剪任务
func AddJobLogRotate() {
	logger.Logger.Info("job: log rotate start")
	des := "日志定时绕接"
	AddTicker(DurationLogRotate, "LogRotate", des, func() {
		if config.ApolloConf.Log.EnableLog == "yes" && config.ApolloConf.Log.LogFile != "" {
			err := utils.ArchiveFile(
				config.ApolloConf.Log.LogFile,
				fmt.Sprintf("%s-%s", config.ApolloConf.Log.LogFile, utils.TimeNowBetterSep()),
				true)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("job log rotate: failed: %s", err.Error()))
			}
		}
	})
}
