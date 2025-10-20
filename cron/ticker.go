/*
Project: Apollo ticker.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"runtime"
	"time"

	"github.com/JJApplication/Apollo/app/task_manager"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/google/uuid"
)

func recoverTask(f func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			logger.LoggerSugar.Errorf("job: runtime error: %v", err)
		default:
			if err != nil {
				logger.LoggerSugar.Errorf("job: task error: %v", err)
			}
		}
	}()
	f()
}

// AddTicker 以s为维度的执行轮询任务
func AddTicker(cf, def int, taskName, des string, f func()) {
	var t int
	t = cf
	if cf <= 0 {
		t = def
	}
	ticker := time.NewTicker(time.Second * time.Duration(t))
	ch := make(chan bool)
	uuidStr := uuid.NewString()
	task_manager.TickerMap[uuidStr] = &task_manager.OneTicker{
		Ch:         ch,
		UUID:       uuidStr,
		Name:       taskName,
		Des:        des,
		Stopped:    false,
		CreateTime: time.Now().Unix(),
		Duration:   t,
		LastRun:    0,
	}

	task_manager.AddBackgroundTask(ch, uuidStr, taskName, des, t, 0)

	go func() {
		for {
			select {
			case <-ticker.C:
				recoverTask(f)
				updateLastRun(uuidStr)
				logger.LoggerSugar.Infof("ticker {%s} [%s] task run at :%s", taskName, uuidStr, utils.TimeNowFormat(utils.TimeForLogger))
			case sig := <-ch:
				if sig {
					ticker.Stop()
					logger.LoggerSugar.Infof("ticker {%s} [%s] stop signal received", taskName, uuidStr)
				}
			}
		}
	}()
}

// InsureTickerExit 确保程序退出时关闭协程
func InsureTickerExit() {
	for _, t := range task_manager.TickerMap {
		t.Stop()
	}
}

func updateLastRun(uuid string) {
	cronLock.Lock()
	defer cronLock.Unlock()
	task_manager.TickerMap[uuid].LastRun = time.Now().Unix()
}
