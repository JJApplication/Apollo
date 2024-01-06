/*
   Create: 2023/6/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package log_manager
// 全局的日志管理模块
// 负责读取日志 日志最大限制只能读取5mb
// 超过部分允许下载日志
// 日志的清理
package log_manager

import (
	"errors"
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/logger"
	"sync"
)

const (
	LogManager = "[LogManager]"
)

var (
	LogManagerPool sync.Map
)

// InitLogManager 初始化日志采集器
//
// 对Apollo运行中各个微服务的日志采集器都是一个单独的协程，互不影响
// 采集器在启动时不会创建，只有首次采集任务执行时会创建，并且微服务在存在采集任务时进入锁状态
func InitLogManager() {
	LogManagerPool = sync.Map{}
	createSelfLogCollector()
	logger.LoggerSugar.Infof("%s init log manager success", LogManager)
}

func createSelfLogCollector() {
	c := Collector{AppName: "Apollo"}
	c.init()
	LogManagerPool.Store("Apollo", c)
}

func getCollector(app string) (Collector, error) {
	val, ok := LogManagerPool.Load(app)
	if !ok {
		return Collector{}, errors.New(ErrorLogCollectorNotExist)
	}

	return val.(Collector), nil
}

func initCollector(app string) error {
	// 查询app是否存在
	if !app_manager.Check(app) {
		return errors.New(ErrorLogAPPNotExist)
	}
	c := Collector{AppName: app}
	c.init()
	LogManagerPool.Store(app, c)
	return nil
}
