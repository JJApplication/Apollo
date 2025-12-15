package router_secure

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	routerSecure := r.Group("/api/secure")
	{
		routerSecure.GET("/list", List)
	}
}
