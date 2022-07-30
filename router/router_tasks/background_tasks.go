/*
Create: 2022/7/30
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_tasks
package router_tasks

import (
	"github.com/JJApplication/Apollo/app/task_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// apollo自己的背景任务

func GetAllBackgroundTasks(c *gin.Context) {
	router.Response(c, task_manager.GetAllBackgroundTasks(), true)
}
