/*
Project: Apollo consts.go
Created: 2021/11/16 by Landers
*/

package app_manager

import "sync"

type appManager struct {
	// APPManagerMap 当前微服务数据的字典
	APPManagerMap sync.Map
	// APPUsingPorts 维护一个当前已经监听的端口组
	APPUsingPorts map[int]struct{}
}

// AppManagerMap 应用的全局字典 每次配置更新后重载到全局字典中
// 键值对appName: App

var APPManager = appManager{
	APPManagerMap: sync.Map{},
	APPUsingPorts: map[int]struct{}{},
}

const (
	APPConfigsRoot = "conf/app"
	APPScriptsRoot = "conf/manager" // 当前独立与zeus
	ConfigSuffix   = ".pig"
	APPStart       = "start"
	APPStop        = "stop"
	APPExit        = "exit"
	APPRestart     = "restart"
	APPKilled      = "killed"
	APPRunning     = "running"
)

const APPStatusOK = 0
const AppUnknown = 555
const (
	_ = 100 + iota
	APPStatusError
	APPStatusStart
	APPStatusStop
	APPStatusExit
	APPStatusRestart
	APPStatusKilled
	APPStatusRunning
)

// StatusMap 返回状态使用的状态表
var StatusMap = map[int]string{
	APPStatusOK:    "running",
	APPStatusError: "stopped",
	APPStatusExit:  "stopped",
	AppUnknown:     "unknown",
}

const (
	APPNotExist = "app not exist"
)

// app的默认类型
const (
	TypeService    = "Service"    // 服务
	TypeWebFront   = "FrontEnd"   // 前端
	TypeMiddleWare = "MiddleWare" // 中间件
	TypeDataStore  = "DataStore"  // 数据层
)

// app的发布状态
const (
	Published = "published" // 已发布
	Testing   = "testing"   // 测试中
	Pending   = "pending"   // 待发布
)

// 配置文件类型
const (
	ConfNginx    = "nginx"
	ConfGunicorn = "gunicorn"
)

// 校验随机端口的合法性
func (am *appManager) checkPorts(port int) bool {
	if _, ok := am.APPUsingPorts[port]; ok {
		return false
	}
	return true
}

// 添加监听端口到端口占用组
func (am *appManager) addPorts(port int) {
	if _, ok := am.APPUsingPorts[port]; !ok {
		am.APPUsingPorts[port] = struct{}{}
	}
}

// 删除不再使用的端口
func (am *appManager) delPorts(port int) {
	if _, ok := am.APPUsingPorts[port]; ok {
		delete(am.APPUsingPorts, port)
	}
}

// 是否存在此服务
func (am *appManager) hasApp(app string) (bool, App) {
	if a, ok := am.APPManagerMap.Load(app); ok {
		return true, a.(App)
	}
	return false, App{}
}
