/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"github.com/JJApplication/Apollo/app/token_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type authBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// Login 根据请求body体中的用户名和密码认证
// 默认返回生成的token 同时设置cookie
// 兼容http场景 默认关闭secure和httpOnly属性
func Login(c *gin.Context) {

	var ab authBody
	if err := c.BindJSON(&ab); err != nil {
		router.Response(c, false, false)
		return
	}
	// 与配置文件比较
	if ab.Account != config.ApolloConf.Server.Account || ab.Password != config.ApolloConf.Server.PassWd {
		router.Response(c, false, false)
		return
	}
	ip := c.Request.RemoteAddr
	if ip == "" {
		router.Response(c, false, false)
		return
	}
	ip = utils.GetRemoteIP(c.Request)
	loginTime := time.Now().Unix()
	token := token_manager.GenerateToken(ip, loginTime)
	token_manager.SetToken(token)
	c.SetCookie("token", token.String(), config.ApolloConf.Server.AuthExpire, "", "", false, false)
	router.Response(c, map[string]string{
		"token":     token.String(),
		"loginIp":   ip,
		"loginTime": utils.GetTimeByFormat(utils.TimeFormat, loginTime),
	}, true)
}
