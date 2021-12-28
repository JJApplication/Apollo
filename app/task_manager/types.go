/*
Project: dirichlet types.go
Created: 2021/12/27 by Landers
*/

package task_manager

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/landers1037/dirichlet/cron"
	"github.com/landers1037/dirichlet/logger"
)

type taskInterface interface {
	Start() (uuid string, err error)
	Stop() (uuid string, err error)
	Check() (uuid string, err error)
	Delete() (uuid string, err error)
	Save() (uuid string, err error)
	GetStatus() (uuid string, err error)
}

type task struct {
	TaskID      string `json:"task_id"`
	TaskName    string `json:"task_name"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
	GoroutineID int    `json:"goroutine_id"` // 只有在任务启动时才会分配
	Stopped     bool   `json:"stopped"`
	IsDeadLine  bool   `json:"is_deadline"` // 设置deadline时 超时则强制终止，回收goroutine资源
	MaxTimeOut  int    `json:"max_timeout"` // 不设置默认无限即goroutine永久存在
	Status      status `json:"status"`
	Fn          func() `json:"_,omitempty"`
}

type status int

const (
	Created status = iota
	// Started
	Stopped
	Running
	Done
	Exited // goroutine被回收强制停止
	Unknown
)

const EmptyName = ""

type taskManager struct {
	CronJobs       map[string]interface{}
	BackGroundJobs map[string]cron.OneTicker
	Tasks          map[string]task
}

func NewTask(name string, timeout int) *task {
	if name == "" {
		name = EmptyName
	}
	now := time.Now().Unix()
	return &task{
		TaskID:      uuid.NewString(),
		TaskName:    name,
		CreateTime:  now,
		UpdateTime:  now,
		GoroutineID: 0,
		Stopped:     false,
		IsDeadLine:  false,
		MaxTimeOut:  timeout,
		Status:      Created,
	}
}

func (t *task) AttachFn(fn func()) (uuid string, err error) {
	if t.Fn == nil {
		t.Fn = fn
		return t.TaskID, err
	}
	return t.TaskID, errors.New(ErrTaskAlreadyAttach)
}

func (t *task) Start() (uuid string, err error) {
	// 创建新协程 启动记录uuid
	if t.TaskID != "" && !t.Stopped && !t.IsDeadLine && t.Status == Created && t.Fn != nil {
		if t.MaxTimeOut > 0 {
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(t.MaxTimeOut))
				defer cancel()
				t.sureRunning()
				t.Fn()
				t.sureDone()
				go func() {
					for {
						select {
						case <-ctx.Done():
							t.sureTimeout()
							logger.Logger.Error(fmt.Sprintf("[task %s]: %s", t.TaskID, ErrTaskTimeout))
							runtime.Goexit()
						}
					}
				}()
			}()
		} else {
			go func() {
				t.sureRunning()
				t.Fn()
				t.sureDone()
			}()
		}
		return t.TaskID, err
	}
	return "", errors.New(ErrTaskStart)
}

func (t *task) Stop() (uuid string, err error) {
	return "", err
}

func (t *task) Check() (uuid string, err error) {
	return "", err
}

func (t *task) Delete() (uuid string, err error) {
	return "", err
}

func (t *task) Save() (uuid string, err error) {
	return "", err
}

func (t *task) GetStatus() (uuid string, err error) {
	return "", err
}

func (t *task) Info() (task task, err error) {
	return *t, err
}

func (t *task) TimeFmt() (s string) {
	return ""
}

// 状态修改
func (t *task) sureRunning() {
	t.Stopped = false
	t.IsDeadLine = false
	t.Status = Running
}

func (t *task) sureDone() {
	t.Stopped = true
	t.Status = Done
}

func (t *task) sureTimeout() {
	t.Stopped = true
	t.Status = Exited
	t.IsDeadLine = true
}

// 手动停止
func (t *task) sureStop() {
	t.Stopped = true
	t.Status = Stopped
}

// 在轮询时产生的特殊状态 unknown及任务为running态 但是轮询任务器距离createtime超出最大限制1天
func (t *task) sureUnknown() {
	t.Status = Unknown
}
