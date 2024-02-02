/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_alarm
package router_alarm

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerAlarm := r.Group("/api/alarm")
	{
		routerAlarm.GET("/all", GetAllAlarm)
		routerAlarm.GET("/top", GetTopNAlarm)
		routerAlarm.GET("/info", GetAlarmInfo)
	}
	routerAlarmWithAuth := r.Group("/api/alarm", middleware.MiddleWareAuth())
	{
		routerAlarmWithAuth.POST("/del", DeleteAlarm)
	}
}
