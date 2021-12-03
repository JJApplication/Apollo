/*
Project: dirichlet ticker.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/landers1037/dirichlet/logger"
)

// 定时器
// 以goroutine的方式执行定时轮询
// 所有轮询任务都拥有唯一的uuid存储在TickerMap中
// 可以通过通道关闭指定uuid的ticker
// todo如何恢复

type oneTicker struct {
	ch         chan bool
	uuid       string
	name       string
	stopped    bool
	createTime int64
}

func (tc *oneTicker) Stop() {
	tc.ch <- true
}

func (tc *oneTicker) SetStop() {
	tc.stopped = true
}

var TickerMap = map[string]oneTicker{}

func init() {
	TickerMap = make(map[string]oneTicker, 1)
}

// AddTicker 以s为维度的执行轮询任务
func AddTicker(t int, taskName string, f func()) {
	ticker := time.NewTicker(time.Second * time.Duration(t))
	ch := make(chan bool)
	uuidStr := uuid.NewString()
	TickerMap[uuidStr] = oneTicker{
		ch:         ch,
		uuid:       uuidStr,
		name:       taskName,
		stopped:    false,
		createTime: time.Now().Unix(),
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				f()
				logger.Logger.Info(fmt.Sprintf("ticker [%s] task run at :%s", uuidStr, time.Now().String()))
			case sig := <-ch:
				if sig {
					ticker.Stop()
					logger.Logger.Info(fmt.Sprintf("ticker [%s] stop signal received", uuidStr))
				}
			}
		}
	}()
}

// InsureTickerExit 确保程序退出时关闭协程
func InsureTickerExit() {
	for _, t := range TickerMap {
		t.Stop()
		t.SetStop()
	}
}
