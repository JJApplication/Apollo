/*
Create: 2022/8/22
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package task_manager
package task_manager

// 定时器
// 以goroutine的方式执行定时轮询
// 所有轮询任务都拥有唯一的uuid存储在TickerMap中
// 可以通过通道关闭指定uuid的ticker
// todo如何恢复

var TickerMap = map[string]*OneTicker{}

type OneTicker struct {
	Ch         chan bool `json:",omitempty"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Des        string    `json:"des"`
	Stopped    bool      `json:"stopped"`
	CreateTime int64     `json:"create_time"`
	Duration   int       `json:"duration"`
	LastRun    int64     `json:"lastRun"`
}

type OneTickerRes struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Des        string `json:"des"`
	Stopped    bool   `json:"stopped"`
	CreateTime int64  `json:"create_time"`
	Duration   int    `json:"duration"`
	LastRun    int64  `json:"lastRun"`
}

func (tc *OneTicker) Start() (uuid string, err error) {
	tc.Stopped = false
	return tc.UUID, nil
}

func (tc *OneTicker) Stop() (uuid string, err error) {
	tc.Stopped = true
	tc.Ch <- true
	return tc.UUID, nil
}
