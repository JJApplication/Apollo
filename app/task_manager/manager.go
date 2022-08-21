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

func GetAllBackgroundTasks() []OneTickerRes {
	var res []OneTickerRes
	for _, v := range TickerMap {
		res = append(res, OneTickerRes{
			UUID:       v.UUID,
			Name:       v.Name,
			Des:        v.Des,
			Stopped:    v.Stopped,
			CreateTime: v.CreateTime,
			Duration:   v.Duration,
			LastRun:    v.LastRun,
		})
	}

	// 排序
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})
	return res
}

// GetAllCronTasks 定时任务
func GetAllCronTasks() []taskRes {
	var res []taskRes
	for _, v := range TaskManager.CronJobs {
		res = append(res, taskRes{
			TaskID:      v.TaskID,
			TaskName:    v.TaskName,
			Spec:        v.Spec,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			GoroutineID: v.GoroutineID,
			Stopped:     v.Stopped,
			IsDeadLine:  v.IsDeadLine,
			MaxTimeOut:  v.MaxTimeOut,
			Status:      v.Status,
		})
	}
	sort.SliceStable(res, func(i, j int) bool {
		return res[i].TaskName < res[j].TaskName
	})
	return res
}
