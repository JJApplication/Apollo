/*
   Create: 2024/5/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package logger

import (
	"fmt"
	"strings"
)

// app manager使用的快速日志记录
// 日志记录包括:
// log level
// [xxx manager]
// log message
// error detail

type GLogger struct{}

var (
	G = GLogger{}
)

func wrapAppName(app string) string {
	if strings.HasPrefix(app, "[") && strings.HasSuffix(app, "]") {
		return app
	}
	return fmt.Sprintf("[%s]", app)
}

func (g GLogger) INFO(app string, args ...any) {
	if len(args) <= 1 {
		LoggerSugar.Info(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprint(args...)))
		return
	}
	fmtStr := args[0]
	LoggerSugar.Infof(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprintf(fmtStr.(string), args[1:]...)))
}

func (g GLogger) Error(app string, args ...any) {
	if len(args) <= 1 {
		LoggerSugar.Error(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprint(args...)))
		return
	}
	fmtStr := args[0]
	LoggerSugar.Errorf(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprintf(fmtStr.(string), args[1:]...)))
}

func (g GLogger) WARN(app string, args ...any) {
	if len(args) <= 1 {
		LoggerSugar.Warn(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprint(args...)))
		return
	}
	fmtStr := args[0]
	LoggerSugar.Warnf(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprintf(fmtStr.(string), args[1:]...)))
}

func (g GLogger) DEBUG(app string, args ...any) {
	if len(args) <= 1 {
		LoggerSugar.Debug(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprint(args...)))
		return
	}
	fmtStr := args[0]
	LoggerSugar.Debugf(fmt.Sprintf("%s %s", wrapAppName(app), fmt.Sprintf(fmtStr.(string), args[1:]...)))
}
