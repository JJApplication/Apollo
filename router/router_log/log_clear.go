/*
   Create: 2025/7/12
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_log

import (
	"github.com/JJApplication/Apollo/app/log_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// ClearAPPLog 清空全部gz压缩的日志
func ClearAPPLog(c *gin.Context) {
	if err := log_manager.ClearAPPLog(); err != nil {
		router.Response(c, err.Error(), false)
		return
	}
	router.Response(c, nil, true)
}
