/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	router.Response(c, nil, true)
}
