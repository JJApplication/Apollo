/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_alarm
package router_alarm

import (
	"github.com/JJApplication/Apollo/app/alarm_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetAllAlarm(c *gin.Context) {
	router.Response(c, alarm_manager.GetAllAlarm(), true)
}

func GetTopNAlarm(c *gin.Context) {
	router.Response(c, alarm_manager.GetTopNAlarm(), true)
}

func GetAlarmInfo(c *gin.Context) {
	id := c.Query("id")
	router.Response(c, alarm_manager.GetAlarmInfo(id), true)
}

func DeleteAlarm(c *gin.Context) {
	id := c.Query("id")
	err := alarm_manager.DeleteAlarm(id)
	if err != nil {
		router.Response(c, nil, false)
	}
	router.Response(c, nil, true)
}
