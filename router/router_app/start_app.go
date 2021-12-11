/*
Project: dirichlet start_app.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/router"
)

// StartApp 启动app
// 通过action判断是否异步执行 所有的异步任务由task manager管理
func StartApp(c *gin.Context) {
	app := c.Query("app")
	if app == "" {
		router.Response(c, "", false)
	}

	isOk, err := app_manager.Start(app)
	if err != nil || !isOk {
		router.Response(c, err, false)
		return
	}

	router.Response(c, err, true)
}

func StartAppAll(c *gin.Context) {
	err := app_manager.StartAll()
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, err, true)
}
