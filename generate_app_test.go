/*
Project: Apollo generate_app_test.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"testing"

	"github.com/JJApplication/Apollo/app/app_manager"
)

var apps = []string{"NoEngine", "MySite", "Mgek", "MgekDoc", "Blog", "JJMail", "JJRobot",
	"CookBook", "JJService", "Redis", "Plume", "MgekFile", "JJGo", "Frp", "BlogFront", "JJServiceFront",
	"PlumeFront", "JJGoFront", "Works",
}

// 生成app配置文件
func TestGenerateApps(t *testing.T) {
	return
	for _, app := range apps {
		err := app_manager.NewApp(app)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func TestGenerateAppScripts(t *testing.T) {
	// 测试态下会覆盖原有文件
	return
	for _, app := range apps {
		err := app_manager.NewAppScript(app)
		if err != nil {
			t.Error(err.Error())
		}
	}
}
