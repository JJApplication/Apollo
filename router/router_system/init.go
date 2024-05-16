/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerSystem := r.Group("/api/system")
	{
		routerSystem.GET("/overview", SystemOverview)
		routerSystem.GET("/pid", SystemPID)
		routerSystem.GET("/config", GetConfig)
		// 保留接口 系统信息包括cpu 内存 其他信息等
		routerSystem.GET("/info", SystemInfo)
		// 更新系统配置 -> 运行时生效
		routerSystem.POST("/config", middleware.MiddleWareAuth(), UpdateConfig)
		// 保存系统配置
		routerSystem.POST("/save", middleware.MiddleWareAuth(), SaveConfig)
		// 软重启
		routerSystem.POST("/reload", middleware.MiddleWareAuth(), ReloadConfig)
		// 清除web缓存
		routerSystem.POST("/clear", middleware.MiddleWareAuth(), ClearCache)
		// SSL证书信息
		routerSystem.GET("/cert", SystemCert)
	}
}
