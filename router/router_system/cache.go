/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/engine"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func ClearCache(c *gin.Context) {
	store := engine.GetStore()
	if err := store.Flush(); err != nil {
		router.Response(c, false, false)
		return
	}

	router.Response(c, true, true)
}
