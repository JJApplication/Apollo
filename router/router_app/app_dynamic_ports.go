/*
   Create: 2023/9/17
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_app

import (
	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func GetDynamicPortApp(c *gin.Context) {
	router.Response(c, app_manager.GetRuntimePortApp(), true)
}
