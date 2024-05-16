/*
   Create: 2024/5/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_system

import (
	"github.com/JJApplication/Apollo/app/cert_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/gin-gonic/gin"
)

func SystemCert(c *gin.Context) {
	router.Response(c, cert_manager.CertInfo(), true)
}
