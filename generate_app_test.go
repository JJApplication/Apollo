/*
Project: dirichlet generate_app_test.go
Created: 2021/11/28 by Landers
*/

package main

import (
	"testing"

	"github.com/landers1037/dirichlet/app/app_manager"
)

var apps = []string{"NoEngine", "MySite", "Mgek", "MgekDoc", "Blog", "JJMail", "JJRobot",
	"CookBook", "JJService", "Redis", "Plume", "MgekFile", "JJGo", "Frp", "Blog_Front", "JJService_Front",
	"Plume_Front", "JJGo_Front", "Works",
}

// 生成app配置文件
func TestGenerateApps(t *testing.T) {
	for _, app := range apps {
		err := app_manager.NewApp(app)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func TestGenerateAppScripts(t *testing.T) {
	for _, app := range apps {
		err := app_manager.NewAppScript(app)
		if err != nil {
			t.Error(err.Error())
		}
	}
}
