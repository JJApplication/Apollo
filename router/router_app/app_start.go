/*
Project: Apollo start_app.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// StartApp 启动app
// 通过action判断是否异步执行 所有的异步任务由task manager管理
// @Summary 启动APP
// @Tags APP Manager
// @Description 启动APP接口
// @Accept application/json
// @Produce application/json
// @param app query string false "app名称"
// @Success 200
// @Router /api/app/start [post]
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

// StartAll
// @Summary 启动所有APP
// @Tags APP Manager
// @Description 启动所有APP接口
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /api/app/startall [post]
func StartAll(c *gin.Context) {
	_, err := app_manager.StartAll()
	if err != nil {
		router.Response(c, err, false)
		return
	}
	router.Response(c, err, true)
}
