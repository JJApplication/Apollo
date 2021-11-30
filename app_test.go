/*
Project: dirichlet app_test.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"testing"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/logger"
)

func init() {
	logger.InitLogger()
}

// app管理测试
func TestAppControl(t *testing.T) {
	app := app_manager.App{Name: "Blog", ID: "test", ManageCMD: app_manager.CMD{
		Start:     "start.sh",
		Stop:      "stop.sh",
		Restart:   "restart.sh",
		ForceKill: "kill.sh",
		Check:     "check.sh",
	},
		RunData: app_manager.RunData{RandomPort: true, Ports: []int{1}}}
	app_manager.APPManager.APPManagerMap.Store(app.Name, app)
	// test app sync
	app.Sync()
	// test app toJson
	t.Log(app.ToJSON())
	// test app start
	ok, err := app.Start()
	if err != nil {
		t.Error("test: ", err.Error())
	}

	t.Logf("app start %v\n", ok)
	appKey, _ := app_manager.APPManager.APPManagerMap.Load("Blog")
	t.Logf("%+v", appKey)
	t.Logf("%+v", app_manager.APPManager)
}
