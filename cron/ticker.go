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

type OneTicker struct {
	ch         chan bool
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Des        string `json:"des"`
	Stopped    bool   `json:"stopped"`
	CreateTime int64  `json:"create_time"`
}

func (tc *OneTicker) Stop() {
	tc.Stopped = true
	tc.ch <- true
}

var TickerMap = map[string]OneTicker{}

func init() {
	TickerMap = make(map[string]OneTicker, 1)
}

// AddTicker 以s为维度的执行轮询任务
func AddTicker(t int, taskName, des string, f func()) {
	ticker := time.NewTicker(time.Second * time.Duration(t))
	ch := make(chan bool)
	uuidStr := uuid.NewString()
	TickerMap[uuidStr] = OneTicker{
		ch:         ch,
		UUID:       uuidStr,
		Name:       taskName,
		Des:        des,
		Stopped:    false,
		CreateTime: time.Now().Unix(),
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				f()
				logger.Logger.Info(fmt.Sprintf("ticker {%s} [%s] task run at :%s", taskName, uuidStr, time.Now().String()))
			case sig := <-ch:
				if sig {
					ticker.Stop()
					logger.Logger.Info(fmt.Sprintf("ticker {%s} [%s] stop signal received", taskName, uuidStr))
				}
			}
		}
	}()
}

// InsureTickerExit 确保程序退出时关闭协程
func InsureTickerExit() {
	for _, t := range TickerMap {
		t.Stop()
	}
}
