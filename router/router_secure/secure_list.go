package router_secure

import (
	"github.com/JJApplication/Apollo/app/secure_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	data := secure_manager.GetSecureList()
	router.Response(c, data, true)
}
