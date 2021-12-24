/*
Project: dirichlet middleware_noroute.go
Created: 2021/12/25 by Landers
*/

package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareNoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error/error404.tmpl", nil)
	}
}
