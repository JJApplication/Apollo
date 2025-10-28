/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_app

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/app/process_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetAppProc(c *gin.Context) {
	app := c.Query("name")
	if !app_manager.Check(app) {
		router.Response(c, "app not exist", false)
		return
	}

	proc := process_manager.Get().GetProc(app)
	router.Response(c, proc, true)
}
