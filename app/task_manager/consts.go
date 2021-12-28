/*
Project: dirichlet consts.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"github.com/fatih/set"
)

var TaskManager taskManager
var TaskGroup []interface{}

func init() {
	TaskManager = *new(taskManager)
}

// GetSetofTasks 任务去重处理，理论上不存在重复的uuid
// 仅针对Task
func GetSetofTasks() {
	s := set.New(set.ThreadSafe)
	for k, _ := range TaskManager.Tasks {
		s.Add(k)
	}
	t := s.List()
	TaskGroup = t
}
