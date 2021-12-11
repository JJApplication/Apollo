/*
Project: dirichlet check_async.go
Created: 2021/12/7 by Landers
*/

package utils

import (
	"github.com/gin-gonic/gin"
)

// CheckAsync 通过检查是否异步执行
func CheckAsync(query *gin.Context) bool {
	q := query.Query("action")
	return q == "async"
}
