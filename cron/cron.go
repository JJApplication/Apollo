/*
Project: Apollo cron.go
Created: 2021/11/24 by Landers
*/

package cron

import (
	"errors"
	"sync"
	"time"

	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	CronID       cron.EntryID `json:"cron_id"`
	CronName     string       `json:"cron_name"`
	CreateTime   time.Time    `json:"create_time"`
	Start        time.Time    `json:"start"`
	ExecuteTimes int          `json:"execute_times"`
	Status       string       `json:"status"`
	Finish       bool         `json:"finish"`
	IsSaved      bool         `json:"is_saved"` // 持久化数据库
}

// CronManager cron内部实现基于cronV3
type CronManager struct {
	cronInstance *cron.Cron
	CreateTime   string
}

// CronJobMap 保存所有的定时任务 没必要使用syncmap 采用内部加锁的方式保证数据安全性
var CronJobMap map[cron.EntryID]*Cron

// InitCronManager 创建一个cron实例
func InitCronManager() CronManager {
	return CronManager{
		cron.New(cron.WithSeconds()),
		utils.TimeNowString(),
	}
}

// Check 检查cron的实例是否存在
func (c *CronManager) Check() bool {
	if c.cronInstance != nil {
		return true
	}
	return false
}

// 创建cron任务成功后，加入到全局任务内存
func (c *CronManager) innerNewCronTask(spec, name string, f func()) (cronID cron.EntryID, err error) {
	cronID, err = c.cronInstance.AddFunc(spec, f)
	if err == nil {
		l := sync.Mutex{}
		l.Lock()
		c := &Cron{
			CronID:       cronID,
			CronName:     name,
			CreateTime:   time.Now(),
			Start:        time.Now(),
			ExecuteTimes: 0,
			Status:       "",
			Finish:       false,
			IsSaved:      false,
		}

		CronJobMap[cronID] = c
		defer l.Unlock()
	}
	return
}

// AddTask 添加cron任务 添加任务仅是存储并不会立刻启动
func (c *CronManager) AddTask(spec, name string, f func()) (cronID cron.EntryID, err error) {
	if !c.Check() {
		return 0, errors.New(ErrCronNull)
	}
	id, err := c.innerNewCronTask(spec, name, f)
	if err != nil {
		logger.LoggerSugar.Errorf("cronjob {%s} [%s] created failed", name, spec)
	}
	logger.LoggerSugar.Infof("cronjob {%d} {%s} [%s] created successfully", id, name, spec)

	return id, err
}

// DelTask 直接删除定时任务 由cron保证从内存中去除
func (c *CronManager) DelTask(id cron.EntryID) error {
	if !c.Check() {
		return errors.New(ErrCronNull)
	}
	c.cronInstance.Remove(id)
	l := sync.Mutex{}
	delete(CronJobMap, id)
	defer l.Unlock()

	return nil
}

// StartTask 启动cron
func (c *CronManager) StartTask(id cron.EntryID) error {
	if !c.Check() {
		return errors.New(ErrCronNull)
	}
	c.cronInstance.Start()
	l := sync.Mutex{}
	CronJobMap[id].Status = "START"
	defer l.Unlock()
	return nil
}

// StopTask 停止cron
func (c *CronManager) StopTask(id cron.EntryID) error {
	if !c.Check() {
		return errors.New(ErrCronNull)
	}
	c.cronInstance.Stop()
	l := sync.Mutex{}
	CronJobMap[id].Status = "STOP"
	defer l.Unlock()
	return nil
}

// AllTask 返回cron内部的全部任务
func (c *CronManager) AllTask() ([]cron.Entry, error) {
	if !c.Check() {
		return nil, errors.New(ErrCronNull)
	}
	return c.cronInstance.Entries(), nil
}
