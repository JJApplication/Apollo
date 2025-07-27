package router_indicator

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerSystem := r.Group("/api/indicator")
	{
		routerSystem.GET("/sys", SystemInfo)
		routerSystem.GET("/load", IndicatorLoad)
		routerSystem.GET("/cpu", IndicatorCPU)
		routerSystem.GET("/mem", IndicatorMem)
		routerSystem.GET("/io", IndicatorIO)
		routerSystem.GET("/network", IndicatorNetwork)
	}
}
