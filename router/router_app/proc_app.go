/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_app

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

func GetAppProc(c *gin.Context) {
	app := c.Query("name")
	if !app_manager.Check(app) {
		router.Response(c, "app not exist", false)
		return
	}

	proc := utils.FilterProcess(app)
	if proc == nil {
		router.Response(c, utils.SysProc{}, false)
		return
	}
	router.Response(c, utils.SysProc{
		PID:            utils.GetProcessPID(proc),
		CPUPercent:     utils.CalcProcessCpu(proc),
		ProcessMemInfo: utils.CalcProcessMem(proc),
		ProcessIO:      utils.CalcProcessIO(proc),
		NetConnections: utils.CalcProcessNet(proc),
		Threads:        utils.GetProcessThreads(proc),
	}, true)
}
