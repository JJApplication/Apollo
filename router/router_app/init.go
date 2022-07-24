/*
Project: Apollo init.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"github.com/gin-gonic/gin"
)

var routerApp *gin.RouterGroup

func Init(r *gin.Engine) {
	routerApp = r.Group("/api/app")
	{
		routerApp.GET("/all", StatusApp)
		routerApp.GET("/app", StatusApp)
		routerApp.GET("/tree", FileTree)

		routerApp.POST("/start", StartApp)
		routerApp.POST("/startall", StartAppAll)
	}
}
