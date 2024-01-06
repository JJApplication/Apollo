/*
   Create: 2024/1/6
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_log

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerLog := r.Group("/api/log")
	{
		routerLog.GET("/list", GetAPPLogDir)
		routerLog.GET("", GetAPPLog)
	}
	routerLogWithAuth := r.Group("/api/log", engine.MiddleWareAuth())
	{
		routerLogWithAuth.GET("/download", GetAPPLogDownload)
	}
}
