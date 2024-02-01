/*
Project: Apollo middleware_cors.go
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
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTION"}
	config.AllowHeaders = []string{"token", " access-token", "Token", "auth", "Auth", "Mgek", "jjapp", "JJApp", "plnack", "content-type", "ContentType"}

	return cors.New(config)
}
