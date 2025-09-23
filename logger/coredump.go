//go:build linux

/*
Create: 2022/8/22
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package logger
package logger

import (
	"os"
	"path/filepath"
	"syscall"
)

// 记录os.Exit(1)类型的fatal错误
// 不能保证程序的变量都可用 所以默认存储到tmp目录下

const (
	Coredump = "Apollo.coredump"
)

var (
	dumpLog = filepath.Join(os.TempDir(), Coredump)
)

// CoreDump 在程序主入口记录
// 重定向os.stderr
// 仅出错时记录 每次启动时删除旧的coredump
func CoreDump() {
	logFile, err := os.OpenFile(dumpLog, os.O_CREATE|os.O_RDWR|os.O_SYNC|os.O_TRUNC, 0644)
	if err != nil {
		return
	}

	err = syscall.Dup3(int(logFile.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		return
	}
	return
}
