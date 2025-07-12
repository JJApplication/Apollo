// Package router_env
package router_env

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerAlarm := r.Group("/api/env")
	{
		routerAlarm.GET("/list", ListServices)
	}
	routerAlarmWithAuth := r.Group("/api/env", middleware.MiddleWareAuth())
	{
		routerAlarmWithAuth.POST("/show", GetEnvs)
		routerAlarmWithAuth.POST("/get", GetEnv)
		routerAlarmWithAuth.POST("/decrypt", GetEnvWithAES)
		routerAlarmWithAuth.POST("/set", SetEnv)
		routerAlarmWithAuth.POST("/delete", DeleteEnv)
	}
}
