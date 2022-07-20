/*
Project: Apollo app_test.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"testing"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/octopus_meta"
)

func init() {
	logger.InitLogger()
}

// app管理测试
func TestAppControl(t *testing.T) {
	var app app_manager.App
	meta := octopus_meta.App{Name: "Blog", ID: "test", ManageCMD: octopus_meta.CMD{
		Start:     "start.sh",
		Stop:      "stop.sh",
		Restart:   "restart.sh",
		ForceKill: "kill.sh",
		Check:     "check.sh",
	},
		RunData: octopus_meta.RunData{RandomPort: true, Ports: []int{1}}}
	app.Meta = meta
	app_manager.APPManager.APPManagerMap.Store(app.Meta.Name, app)
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
