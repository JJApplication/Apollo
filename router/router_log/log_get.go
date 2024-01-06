/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

package router_log

import (
	"github.com/JJApplication/Apollo/app/log_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetAPPLogDir(c *gin.Context) {
	app := c.Query("app")
	if app == "" {
		router.Response(c, "", false)
		return
	}
	router.Response(c, log_manager.GetAPPLogList(app), true)
}

func GetAPPLog(c *gin.Context) {
	app := c.Query("app")
	if app == "" {
		router.Response(c, "", false)
		return
	}
	router.Response(c, log_manager.GetAPPLog(app), true)
}
