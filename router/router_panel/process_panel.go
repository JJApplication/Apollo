package router_panel

import (
	"github.com/JJApplication/Apollo/app/process_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func ProcessList(c *gin.Context) {
	data := process_manager.Get().GetProcList()
	router.Response(c, data, true)
}

func Process(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		router.Response(c, nil, false)
		return
	}
	data := process_manager.Get().GetProcHistory(name)
	router.Response(c, data, true)
}
