/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
	"math"
	"time"
)

// CalcCpuLoad 当前系统cpu使用率
func CalcCpuLoad() float64 {
	data, err := cpu.Percent(time.Second, true)
	if err != nil {
		return 0
	}
	// 计算总值
	var all float64
	for _, v := range data {
		all += v
	}
	return math.Floor(all / float64(len(data)))
}

// CalcProcessCpu 获取当前进程的CPU占用
func CalcProcessCpu(p *process.Process) float64 {
	res, err := p.CPUPercent()
	if err != nil {
		return 0
	}

	return res
}
