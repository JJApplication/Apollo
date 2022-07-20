/*
Project: Apollo middleware_nomethod.go
Created: 2021/12/25 by Landers
*/

package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareNoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusMethodNotAllowed, "error/error405.tmpl", nil)
	}
}
