// Package router_env
package router_env

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerEnv := r.Group("/api/env")
	{
		routerEnv.GET("/list", ListServices)
	}
	routerEnvWithAuth := r.Group("/api/env", middleware.MiddleWareAuth())
	{
		routerEnvWithAuth.POST("/show", GetEnvs)
		routerEnvWithAuth.POST("/get", GetEnv)
		routerEnvWithAuth.POST("/decrypt", GetEnvWithAES)
		routerEnvWithAuth.POST("/set", SetEnv)
		routerEnvWithAuth.POST("/delete", DeleteEnv)
	}
}
