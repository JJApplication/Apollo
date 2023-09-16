/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import "github.com/shirou/gopsutil/process"

// CalcProcessNet 获取进程的网络堆栈
func CalcProcessNet(p *process.Process) int {
	conns, _ := p.Connections()
	return len(conns)
}
