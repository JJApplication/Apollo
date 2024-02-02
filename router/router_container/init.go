/*
Create: 2022/7/31
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package router_container
package router_container

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerContainer := r.Group("/api/container")
	{
		routerContainer.GET("/containers", GetAllContainer)
		routerContainer.GET("/images", GetAllImages)
	}
	routerContainerAuth := r.Group("/api/container", middleware.MiddleWareAuth())
	{
		routerContainerAuth.POST("/start", StartContainer)
		routerContainerAuth.POST("/stop", StopContainer)
		routerContainerAuth.POST("/remove", RemoveContainer)
	}
}
