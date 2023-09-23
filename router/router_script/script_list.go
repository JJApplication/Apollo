/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_script

import (
	"github.com/JJApplication/Apollo/app/script_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	router.Response(c, script_manager.GetScripts(), true)
}
