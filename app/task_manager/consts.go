/*
Project: Apollo consts.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"github.com/fatih/set"
)

const (
	TaskManagerPrefix = "[Task Manager]"
)

var TaskManager *taskManager
var TaskGroup []interface{}

func InitTaskManager() {
	TickerMap = make(map[string]*OneTicker, 1)
	TaskManager = new(taskManager)
	TaskManager.CronJobs = make(map[string]task, 1)
	TaskManager.BackGroundJobs = TickerMap
	LoadFromData()
	SyncData()
	logger.LoggerSugar.Infof("%s init all tasks to TaskManager", TaskManagerPrefix)
}

// GetSetOfTasks 任务去重处理，理论上不存在重复的uuid
// 仅针对Task
func GetSetOfTasks() {
	s := set.New(set.ThreadSafe)
	for k, _ := range TaskManager.CronJobs {
		s.Add(k)
	}
	t := s.List()
	TaskGroup = t
}
