/*
Project: dirichlet cron.go
Created: 2021/11/24 by Landers
*/

package cron

import (
	"sync"
	"time"
)

type Cron struct {
	rwLock       *sync.RWMutex
	CronID       string    `json:"cron_id"`
	CronName     string    `json:"cron_name"`
	CreateTime   time.Time `json:"create_time"`
	Start        time.Time `json:"start"`
	ExecuteTimes int       `json:"execute_times"`
	TimeOut      int       `json:"time_out"`
	Status       string    `json:"status"`
	Finish       bool      `json:"finish"`
	IsSaved      bool      `json:"is_saved"` // 持久化数据库
}

type CronManager struct {
}
