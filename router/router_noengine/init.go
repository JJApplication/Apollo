/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_noengine
package router_noengine

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerNoEngine := r.Group("/api/noengine")
	{
		routerNoEngine.GET("", GetNoEngineApp)
		routerNoEngine.GET("/all", GetAllNoEngineApp)
		routerNoEngine.GET("/status", GetNoEngineStatus)
	}
	routerNoEngineWithAuth := r.Group("/api/noengine", engine.MiddleWareAuth())
	{
		routerNoEngineWithAuth.POST("/start", StartNoEngineApp)
		routerNoEngineWithAuth.POST("/stop", StopNoEngineApp)
		routerNoEngineWithAuth.POST("/resume", ResumeNoEngineApp)
		routerNoEngineWithAuth.POST("/pause", PauseNoEngineApp)
		routerNoEngineWithAuth.POST("/remove", RemoveNoEngineApp)
		routerNoEngineWithAuth.POST("/status", GetNoEngineStatus)
		routerNoEngineWithAuth.POST("/refresh", RefreshNoEngineApp)
	}
	routerNoEngineInner := r.Group("/api/noengine/x", engine.MiddleWareXLocal())
	{
		routerNoEngineInner.POST("/start", StartNoEngineApp)
		routerNoEngineInner.POST("/stop", StopNoEngineApp)
		routerNoEngineInner.POST("/resume", ResumeNoEngineApp)
		routerNoEngineInner.POST("/pause", PauseNoEngineApp)
		routerNoEngineInner.POST("/remove", RemoveNoEngineApp)
		routerNoEngineInner.POST("/status", GetNoEngineStatus)
		routerNoEngineInner.POST("/refresh", RefreshNoEngineApp)
	}
}
