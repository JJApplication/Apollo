/*
Project: Apollo manager.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"errors"
	"fmt"
	"sort"

	"github.com/JJApplication/Apollo/cron"
)

// 任务管理
// 所有的任务类 都有Start Stop Check方法

// ListAllTasks 列出全部任务
func ListAllTasks() ([]string, error) {
	if len(TaskManager.Tasks) <= 0 {
		return nil, errors.New("no tasks")
	}
	var list []string
	for _, v := range TaskManager.Tasks {
		list = append(list, fmt.Sprintf("Name: %s ID: %s Create: %v, Status: %v",
			v.TaskName, v.TaskID, v.CreateTime, v.Status))
	}
	return list, nil
}

func GetAllBackgroundTasks() []*cron.OneTicker {
	var res []*cron.OneTicker
	for _, v := range cron.TickerMap {
		res = append(res, v)
	}

	// 排序
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})
	return res
}
