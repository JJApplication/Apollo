/*
Project: dirichlet middleware_cors.go
Created: 2022/2/18 by Landers
*/

package engine

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MiddleWareCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowWebSockets = true
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{}

	return cors.New(config)
}
