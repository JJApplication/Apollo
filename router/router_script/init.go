/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_script

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerScript := r.Group("/api/script")
	{
		routerScript.GET("/list", List)
		routerScript.GET("/task/list", ScriptTaskList)
		routerScript.GET("/task", ScriptTaskByName)
		routerScript.POST("/task/start", middleware.MiddleWareAuth(), ScriptTaskStart)
		routerScript.POST("/task/stop", middleware.MiddleWareAuth(), ScriptTaskStop)
		routerScript.POST("/task/delete", middleware.MiddleWareAuth(), ScriptTaskDelete)
	}
}
