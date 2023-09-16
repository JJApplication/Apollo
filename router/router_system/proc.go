/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

func SystemInfo(c *gin.Context) {
	p, f, v := utils.CalcPlatform()
	router.Response(c, struct {
		Kernel   string  `json:"kernel"`
		Platform string  `json:"platform"`
		Family   string  `json:"family"`
		Version  string  `json:"version"`
		CPU      float64 `json:"cpu"`
		MemAvail float64 `json:"memAvail"`
		MemUsed  float64 `json:"memUsed"`
	}{
		Kernel:   utils.CalcKernel(),
		Platform: p,
		Family:   f,
		Version:  v,
		CPU:      utils.CalcCpuLoad(),
		MemAvail: utils.CalcMemAvail(),
		MemUsed:  utils.CalcMemUsed(),
	}, true)
}
