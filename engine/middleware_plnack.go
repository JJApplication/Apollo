/*
Project: dirichlet middleware_plnack.go
Created: 2021/11/30 by Landers
*/

package engine

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/landers1037/dirichlet/logger"
	plnack_proto "github.com/landers1037/plnack-proto"
)

// MiddleWare_Plnack plnack数据编码
func MiddleWare_Plnack() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取响应数据
		if c.Request.Header.Get("proto") == "plnack" {
			d, _ := c.Get("data")
			pl := plnack_proto.PlnackData{
				Key:       "",
				Type:      plnack_proto.PtypeServer,
				Version:   "",
				AppName:   "Dirichlet",
				Data:      d,
				KeyVerify: false,
				Time:      time.Now(),
			}
			err := plnack_proto.EncodeGin(c, pl)
			if err != nil {
				logger.Logger.Error("encode for plnack failed")
			}
			return
		}
	}
}
