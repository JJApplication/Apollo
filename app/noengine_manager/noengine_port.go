/*
   Create: 2024/1/11
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import (
	"fmt"
	"github.com/gookit/goutil/mathutil"
	"net"
)

// 用于生成随机端口 范围5000-8000

const (
	MINPort = 5000
	MaxPort = 8000
	Retry   = 5
)

func randomPort() int {
	port := mathutil.RandomInt(MINPort, MaxPort)
	for i := 0; i < Retry; i++ {
		if tryPort(port) {
			return port
		}
		continue
	}
	// 无法成功找到端口 返回0
	return 0
}

func tryPort(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return false
	}

	_ = ln.Close()
	return true
}
