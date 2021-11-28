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
	app := app_manager.App{Name: "Test", ID: "test", ManageCMD: app_manager.CMD{
		Start:     []string{"start.sh"},
		Stop:      nil,
		Restart:   nil,
		ForceKill: nil,
		Check:     "",
	}}

	ok, err := app.Start()
	if err != nil {
		t.Error("test: ", err.Error())
	}

	t.Logf("app start %v\n", ok)
}
