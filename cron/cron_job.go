/*
Create: 2022/8/22
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package cron
package cron

import (
	"strconv"

	"github.com/JJApplication/Apollo/app/alarm_manager"
	"github.com/JJApplication/Apollo/app/task_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/fushin/cron"
	"github.com/JJApplication/fushin/utils/files"
)

const (
	SpecAPPBackup = "0 0 0 1/7 * ?"
)

var cronGroup *cron.CronGroup

func InitCronJobs() {
	addCronJobBackup()
	addCronJobClearAlarm()
}

func init() {
	cronGroup = cron.NewGroup(SpecAPPBackup)
}

func addCronJobBackup() {
	logger.Logger.Info("cron job: rsync-backup start")
	id, err := cronGroup.AddFunc(func() {
		err := files.RsyncAndTar(config.ApolloConf.APPRoot, config.ApolloConf.APPBackUp, true)
		if err != nil {
			logger.LoggerSugar.Errorf("cron job backup failed: %s", err.Error())
		}
	})

	id2Str := strconv.Itoa(id)
	task_manager.AddCronTask(id2Str, "RsyncBackup", SpecAPPBackup)
	if err != nil {
		logger.LoggerSugar.Errorf("init cron job backup failed: %s", err.Error())
	}
}

func addCronJobClearAlarm() {
	logger.Logger.Info("cron job: clear-alarms start")
	id, err := cronGroup.AddFunc(func() {
		err := alarm_manager.ClearAlarmLimit()
		if err != nil {
			logger.LoggerSugar.Errorf("cron job clear-alarms failed: %s", err.Error())
		}
	})

	id2Str := strconv.Itoa(id)
	task_manager.AddCronTask(id2Str, "ClearAlarms", SpecAPPBackup)
	if err != nil {
		logger.LoggerSugar.Errorf("init cron job clear-alarms failed: %s", err.Error())
	}
}
