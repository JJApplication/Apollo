/*
Project: Apollo background_job.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"github.com/JJApplication/Apollo/app/discover_manager"
	"github.com/JJApplication/Apollo/app/noengine_manager"
	"path/filepath"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
)

// 启动时执行的轮询任务 用于随时刷新持久化数据
// 持久化数据用于恢复
const (
	DurationDbSaver             = 60 * 60
	DurationDbPersist           = 60 * 60 * 24
	DurationAppSync             = 60 * 60
	DurationAppRuntimeSyncDB    = 60
	DurationAppCheck            = 60 * 60
	DurationLogRotate           = 1 * 60 * 60 * 24
	DurationNoEngineRuntimeSync = 1 * 60 * 60
	DurationAppDiscover         = 60 * 10
	DurationNoEngineDiscover    = 60 * 10
)

func InitBackgroundJobs() {
	AddJobDBSaver()
	AddJobDBPersist()
	AddJobAPPSync()
	AddJobOctopusMetaSync()
	AddJobAPPDumps()
	AddJobAPPCheck()
	AddJobLogRotate()
	AddNoEngineSync()
	AddNoEngineReload()
}

// AddJobDBSaver 数据库刷新
func AddJobDBSaver() {
	logger.Logger.Info("job: database sync start")
	des := "同步刷新微服务信息到数据库"
	AddTicker(config.ApolloConf.Task.BackgroundJob.DBSave, DurationDbSaver, "DBSaver", des, func() {
		app_manager.SaveToDB()
	})
}

// AddJobDBPersist 数据库内容持久化
func AddJobDBPersist() {
	logger.Logger.Info("job: database persist start")
	des := "数据库信息持久化存储"
	AddTicker(config.ApolloConf.Task.BackgroundJob.DBPersist, DurationDbPersist, "DBPersist", des, func() {
		app_manager.Persist()
	})
}

// AddJobAPPSync app模型文件同步
// 从本地同步模型文件到Apollo
// 粒度为App
func AddJobAPPSync() {
	logger.Logger.Info("job: app config sync start")
	des := "同步微服务模型文件"
	AddTicker(config.ApolloConf.Task.BackgroundJob.AppSync, DurationAppSync, "AppConfigSync", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.Sync()
			if err != nil {
				logger.LoggerSugar.Errorf("job app config sync failed: %s", err.Error())
			}
			return true
		})
	})
}

// AddJobOctopusMetaSync 重载整个octopus目录
// 粒度为octopus 同步会合并数据，新增微服务删除不存在的服务
func AddJobOctopusMetaSync() {
	logger.Logger.Info("job: octopus-meta sync start")
	des := "App自动发现重载octopus模型文件"
	AddTicker(config.ApolloConf.Task.AutoDiscover.App, DurationAppDiscover, "OctopusMetaSync", des, func() {
		if discover_manager.GetAppDiscover().NeedDiscover() {
			err := app_manager.ReloadManagerMap()
			if err != nil {
				logger.LoggerSugar.Errorf("job octopus-meta sync failed: %s", err.Error())
			}
		}
		logger.LoggerSugar.Info("AutoDiscover task run")
	})
}

// AddJobAPPDumps 从缓存同步配置到mongo
// 用于同步配置参数和端口变量
func AddJobAPPDumps() {
	logger.Logger.Info("job: app runtime sync start")
	des := "同步存储缓存数据到数据库"
	AddTicker(config.ApolloConf.Task.BackgroundJob.AppRuntimeSync, DurationAppRuntimeSyncDB, "AppRuntimeSync", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.SyncDB()
			if err != nil {
				logger.LoggerSugar.Errorf("job app runtime syncDB failed: %s", err.Error())
			}
			return true
		})
	})
}

// AddJobAPPCheck app服务定时检查
func AddJobAPPCheck() {
	logger.Logger.Info("job: app check start")
	des := "微服务状态定时检查"
	AddTicker(config.ApolloConf.Task.BackgroundJob.AppCheck, DurationAppCheck, "AppChecker", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.Check()
			if err != nil {
				logger.LoggerSugar.Errorf("job app checker: %s check failed: %s try to restart", key, err.Error())
				_, err = app.ReStart()
				logger.LoggerSugar.Warnf("job app checker: %s restart result: %s", key, err.Error())
			}
			return true
		})
	})
}

// AddJobLogRotate 日志裁剪任务
// 对log目录下的所有日志绕接
func AddJobLogRotate() {
	logger.Logger.Info("job: log rotate start")
	des := "日志定时绕接"
	AddTicker(config.ApolloConf.Task.BackgroundJob.LogRotate, DurationLogRotate, "LogRotate", des, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			go func() {
				app := value.(app_manager.App)
				err := utils.Rotate(filepath.Join(config.ApolloConf.APPLogDir, app.Meta.Name, app.Meta.Name+".log"))
				if err != nil {
					if err != nil {
						logger.LoggerSugar.Errorf("job log rotate: failed: [%s] %s", app.Meta.Name, err.Error())
					}
				}
			}()
			return true
		})
	})
}

func AddNoEngineSync() {
	logger.Logger.Info("job: NoEngine sync start")
	des := "NoEngine服务信息同步"
	AddTicker(config.ApolloConf.Task.BackgroundJob.NoEngineRuntimeSync, DurationNoEngineRuntimeSync, "NoEngineSync", des, func() {
		_, err := noengine_manager.RefreshNoEngineMap()
		if err != nil {
			logger.LoggerSugar.Errorf("job NoEngine sync: failed: %s", err.Error())
		} else {
			logger.LoggerSugar.Info("job NoEngine sync: success")
		}
	})
}

func AddNoEngineReload() {
	logger.Logger.Info("job: NoEngine reload start")
	des := "NoEngine服务自动发现"
	AddTicker(config.ApolloConf.Task.AutoDiscover.NoEngine, DurationNoEngineDiscover, "NoEngineAutoDiscover", des, func() {
		if discover_manager.GetNoEngineDiscover().NeedDiscover() {
			noengine_manager.LoadAllNoEngineAPPs()
			logger.LoggerSugar.Info("job NoEngine reload: success")
		}
		logger.LoggerSugar.Info("AutoDiscover task run")
	})
}
