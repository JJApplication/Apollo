/*
Create: 2022/8/22
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package task_manager
package task_manager

import "time"

// 定时器
// 以goroutine的方式执行定时轮询
// 所有轮询任务都拥有唯一的uuid存储在TickerMap中
// 可以通过通道关闭指定uuid的ticker
// todo如何恢复

var TickerMap = map[string]*OneTicker{}

// AddBackgroundTask 添加背景任务 已经存在则跳过
func AddBackgroundTask(ch chan bool, uuid string, name, des string, duration int, lastRun int64) {
	TaskManager.lock.Lock()
	defer TaskManager.lock.Unlock()

	// 从持久化数据中加载更新时间
	var ut int64
	ut = lastRun
	pt := GetPersist().GetBackgroundJob(name)
	if pt > 0 {
		ut = pt
	}
	TickerMap[uuid] = &OneTicker{
		Ch:         ch,
		UUID:       uuid,
		Name:       name,
		Des:        des,
		Stopped:    false,
		CreateTime: time.Now().Unix(),
		Duration:   duration,
		LastRun:    ut,
	}
}
