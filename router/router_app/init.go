/*
Project: Apollo init.go
Created: 2021/11/30 by Landers
*/

package router_app

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerApp := r.Group("/api/app")
	{
		routerApp.GET("/all", StatusApp)
		routerApp.GET("/info", InfoApp)
		routerApp.GET("/status", StatusApp)
		routerApp.GET("/tree", FileTree)
		routerApp.GET("/proc", GetAppProc)
		routerApp.GET("/ports", GetDynamicPortApp)
	}
	routerWithAuth := r.Group("/api/app", engine.MiddleWareAuth())
	{
		routerWithAuth.POST("/start", StartApp)
		routerWithAuth.POST("/startall", StartAll)
		routerWithAuth.POST("/stop", StopApp)
		routerWithAuth.POST("/stopall", StopAll)
		routerWithAuth.POST("/restart", ReStartApp)
		routerWithAuth.POST("/upload", Upload)
		routerWithAuth.POST("/remove", Remove)
	}
}
