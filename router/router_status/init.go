/*
   Create: 2023/7/31
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_status

import (
	"github.com/JJApplication/Apollo/app/status_manager"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerApp := r.Group("/api/status")
	{
		routerApp.StaticFile(status_manager.StatusOptFile, status_manager.GetStatusOptFile())
	}
}
