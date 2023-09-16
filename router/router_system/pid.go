/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

func SystemPID(c *gin.Context) {
	router.Response(c, utils.GetRuntimePID(), true)
}
