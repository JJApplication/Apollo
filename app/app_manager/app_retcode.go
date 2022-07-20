/*
Project: Apollo app_retcode.go
Created: 2021/11/27 by Landers
*/

package app_manager

import (
	"fmt"
	"regexp"
	"strconv"
)

// 返回码对照表

const APPManagerPrefix = "[APP Manager]"

var appCodeMap = map[int]string{
	APPStatusError:   appCodeMsg("get app result failed"),
	APPStatusStart:   appCodeMsg("app started failed"),
	APPStatusStop:    appCodeMsg("app stopped failed"),
	APPStatusRestart: appCodeMsg("app restarted failed"),
	APPStatusExit:    appCodeMsg("app exited failed"),
	APPStatusKilled:  appCodeMsg("app killed"),
	APPStatusRunning: appCodeMsg("app running"),
}

func appCodeMsg(s string) string {
	return fmt.Sprintf("%s %s", APPManagerPrefix, s)
}

func errCode(code int) string {
	if v, ok := appCodeMap[code]; ok {
		return v
	}
	return appCodeMsg("invalid code")
}

func toCode(s string) int {
	// 正则匹配出返回码
	reg, err := regexp.Compile("[0-9]*$")
	if err != nil {
		return APPStatusError
	}

	s = reg.FindString(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		return APPStatusError
	}

	return i
}
