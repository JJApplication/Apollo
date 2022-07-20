/*
Project: Apollo random_ports.go
Created: 2021/11/29 by Landers
*/

package utils

import (
	"github.com/gookit/goutil/mathutil"
)

// 当前仅支持随机的单个端口

const (
	MINPort = 10000
	MaxPort = 20000
)

func RandomPort() int {
	return mathutil.RandomInt(MINPort, MaxPort)
}
