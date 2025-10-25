package router_panel

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerPanel := r.Group("/api/panel")
	{
		routerPanel.GET("/network/stat", NetworkStat)
		routerPanel.GET("/network/domain", NetworkDomain)
		routerPanel.GET("/network/geo", NetworkGeoIP)
	}
}
