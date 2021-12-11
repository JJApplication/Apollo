/*
Project: dirichlet exec.go
Created: 2021/12/9 by Landers
*/

package uds

import (
	"fmt"
	"strings"

	"github.com/landers1037/dirichlet/logger"
)

// 执行cmds
// c 是一个空格分割的命令
func execCmd(c string) string {
	c = strings.TrimSpace(c)
	args := strings.Fields(c)
	if len(args) <= 0 {
		return "empty cmd"
	}
	logger.Logger.Info(fmt.Sprintf("recieved cmd from UDS: %s", c))
	fn, ok := cmdsMap[args[0]]
	if !ok {
		return "incorrect cmd"
	}
	return fn(args[1:]...)
}
