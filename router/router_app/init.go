/*
Project: dirichlet init.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var routerApp *gin.RouterGroup

func Init(r *gin.Engine) {
	r.GET("/a", func(context *gin.Context) {
		fmt.Println(context.Request.Header)
		context.String(200, "%s", "sb")
	})
	routerApp = r.Group("/app")
	{
		routerApp.POST("/start", StartApp)
		routerApp.POST("/startall", StartAppAll)
	}
}
