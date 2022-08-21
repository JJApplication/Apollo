/*
Create: 2022/7/30
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_tasks
package router_tasks

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerTask := r.Group("/api/task")
	{
		routerTask.GET("/bg", GetAllBackgroundTasks)
		routerTask.GET("/cron", GetAllCronTasks)
	}
}
