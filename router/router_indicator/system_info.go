package router_indicator

import (
	"github.com/JJApplication/Apollo/app/indicator_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

// SystemInfo 系统指标信息 非实时查询数据而是来自历史指标探测
// 由定时器定时获取性能指标 存储在内存数据库中
func SystemInfo(c *gin.Context) {
	router.Response(c, indicator_manager.GetSystemInfo(), true)
}
