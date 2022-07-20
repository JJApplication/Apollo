/*
Project: Apollo exec.go
Created: 2021/12/9 by Landers
*/

package uds

import (
	"fmt"
	"strings"

	"github.com/JJApplication/Apollo/logger"
)

// 执行cmds
// c 是一个空格分割的命令
func execCmd(c string) UDSRes {
	c = strings.TrimSpace(c)
	args := strings.Fields(c)
	if len(args) <= 0 {
		return UDSRes{
			Error: "empty cmd",
			Data:  "",
		}
	}
	logger.Logger.Info(fmt.Sprintf("recieved cmd from UDS: %s", c))
	fn, ok := cmdsMap[args[0]]
	if !ok {
		return UDSRes{
			Error: "incorrect cmd",
			Data:  "",
		}
	}
	funcRes, err := fn(args[1:]...)
	return UDSRes{
		Error: err,
		Data:  funcRes,
	}
}
