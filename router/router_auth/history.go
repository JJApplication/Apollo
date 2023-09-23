/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package router_auth

import (
	"fmt"
	"github.com/JJApplication/Apollo/app/token_manager"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

type history struct {
	IP   string `json:"loginIp"`
	Time string `json:"loginTime"`
}

func History(c *gin.Context) {
	var res []history
	t, _ := c.Cookie("token")
	fmt.Println("cookie", t)
	data := token_manager.GetTokenMap()
	for _, d := range data {
		res = append(res, history{
			IP:   d.LoginIP,
			Time: utils.GetTimeByFormat(utils.TimeFormat, d.LoginTime),
		})
	}

	router.Response(c, res, true)
}
