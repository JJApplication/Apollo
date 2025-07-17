package router_indicator

import (
	"github.com/JJApplication/Apollo/app/indicator_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func IndicatorLoad(c *gin.Context) {
	router.Response(c, indicator_manager.IndicatorLoad(), true)
}

func IndicatorCPU(c *gin.Context) {
	router.Response(c, indicator_manager.IndicatorCPU(), true)
}

func IndicatorMem(c *gin.Context) {
	router.Response(c, indicator_manager.IndicatorMem(), true)
}

func IndicatorIO(c *gin.Context) {
	router.Response(c, indicator_manager.IndicatorLoad(), true)
}

func IndicatorNetwork(c *gin.Context) {
	router.Response(c, indicator_manager.IndicatorNetwork(), true)
}
