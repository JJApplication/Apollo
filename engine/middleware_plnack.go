/*
Project: Apollo middleware_plnack.go
Created: 2021/11/30 by Landers
*/

package engine

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/JJApplication/Apollo/logger"
	plnack_proto "github.com/landers1037/plnack-proto"
)

const (
	KeyPlnack = "plnack"
)

// MiddlewarePlnack plnack数据编码
func MiddlewarePlnack() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取响应数据
		// 校验headers的proto判断是否需要使用plnack编码
		if strings.ToLower(c.Request.Header.Get("proto")) == "plnack" {
			d, _ := c.Get("data")
			pl := plnack_proto.PlnackData{
				Key:       "",
				Type:      plnack_proto.PtypeServer,
				Version:   "",
				AppName:   "Apollo",
				Data:      d,
				KeyVerify: false,
				Time:      time.Now(),
			}
			plnack_proto.PLNACK_LOG = false
			plnack_proto.PlnackVerify = false
			err := plnack_proto.EncodeGin(c, pl)
			if err != nil {
				logger.Logger.Error("encode for plnack failed")
			}
			return
		}
	}
}
