/*
Project: Apollo dao.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"os"
	"time"
)

// 持久化任务到任务数据库中

func enable() bool {
	return config.ApolloConf.Experiment.TaskData.Path != ""
}

func PersistToData() {
	if !enable() {
		return
	}
	var tm PersistManager
	for _, t := range TaskManager.CronJobs {
		tm.CronJobs = append(tm.CronJobs, struct {
			Name       string `json:"name"`
			UpdateTime int64  `json:"update_time"`
		}{
			t.TaskName,
			t.UpdateTime,
		})
	}

	for _, t := range TaskManager.BackGroundJobs {
		tm.BackGroundJobs = append(tm.BackGroundJobs, struct {
			Name       string `json:"name"`
			UpdateTime int64  `json:"update_time"`
		}{
			t.Name,
			t.LastRun,
		})
	}

	data, err := utils.MarshalMsgPack(tm)
	if err != nil {
		return
	}
	file := config.ApolloConf.Experiment.TaskData.Path
	if err = os.WriteFile(file, data, 0644); err != nil {
		logger.LoggerSugar.Errorf("persist task data >%s failed: %s", file, err.Error())
		return
	}
}

func LoadFromData() {
	if !enable() {
		return
	}
	file := config.ApolloConf.Experiment.TaskData.Path
	data, err := os.ReadFile(file)
	if err != nil {
		logger.LoggerSugar.Errorf("load tasks from persistance: %s failed: %s", file, err.Error())
		return
	}

	var tmp PersistManager
	err = utils.UnmarshalMsgPack(data, &tmp)
	if err != nil {
		logger.LoggerSugar.Errorf("load tasks from persistance: %s failed: %s", file, err.Error())
		return
	}
	persistData = tmp
}

func SyncData() {
	if !enable() {
		return
	}
	t := config.ApolloConf.Experiment.TaskData.Duration
	if t <= 0 {
		t = 60
	}
	go func() {
		ticker := time.NewTicker(time.Duration(t) * time.Second)
		for {
			select {
			case <-ticker.C:
				PersistToData()
			}
		}
	}()
}
