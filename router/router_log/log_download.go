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

func GetAPPLogDownload(c *gin.Context) {
	app := c.Query("app")
	logFile := c.Query("log")
	if app == "" || logFile == "" {
		router.Response(c, "", false)
		return
	}
	router.ResponseFile(c, log_manager.DownloadAPPLogFile(app, logFile))
}
