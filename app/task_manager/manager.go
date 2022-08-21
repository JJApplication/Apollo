/*
Project: Apollo manager.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"sort"
)

// 任务管理
// 所有的任务类 都有Start Stop Check方法

// ListAllTasks 列出全部任务
func ListAllTasks() ([]string, error) {
	var list []string
	return list, nil
}

func GetAllBackgroundTasks() []*OneTicker {
	var res []*OneTicker
	for _, v := range TickerMap {
		res = append(res, v)
	}

	// 排序
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})
	return res
}

// GetAllCronTasks 定时任务
func GetAllCronTasks() []task {
	var res []task
	for _, v := range TaskManager.CronJobs {
		res = append(res, v)
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].TaskName < res[j].TaskName
	})
	return res
}
