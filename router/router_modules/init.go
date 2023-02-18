/*
Create: 2023/2/18
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_modules 动态路由
package router_modules

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerAlarm := r.Group("/api/modules")
	{
		routerAlarm.GET("", func(c *gin.Context) {
			router.Response(c, engine.GetModules(), true)
		})
		routerAlarm.GET("/all", func(c *gin.Context) {
			router.Response(c, engine.GetModules(), true)
		})
		routerAlarm.GET("/:path", func(c *gin.Context) {

		})
	}
}
