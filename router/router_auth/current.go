/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"github.com/JJApplication/Apollo/app/token_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

// Current 返回当前登录生效用户的登入时间和IP
func Current(c *gin.Context) {
	current := token_manager.GetActiveToken()
	if current.LoginIP == "" {
		router.Response(c, false, false)
		return
	}

	router.Response(c, map[string]string{
		"account":   current.LoginUser,
		"loginIp":   current.LoginIP,
		"loginTime": utils.GetTimeByFormat(utils.TimeFormat, current.LoginTime),
	}, true)
}
