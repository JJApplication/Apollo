/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_script

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerScript := r.Group("/api/script")
	{
		routerScript.GET("/list", List)
	}
}
