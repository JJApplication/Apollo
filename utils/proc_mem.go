/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

type ProcessMemInfo struct {
	MemPercent float32 `json:"memPercent"`
	MemRss     uint64  `json:"memRss"`  // 物理内存
	MemVms     uint64  `json:"memVms"`  // 虚拟内存
	MemSwap    uint64  `json:"memSwap"` // 交换内存
}

// CalcMemUsed 当前内存使用率
// used/total
func CalcMemUsed() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return v.UsedPercent
}

// CalcMemAvail 当前内存可用大小
func CalcMemAvail() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return float64(v.Available)
}

func CalcProcessMem(p *process.Process) ProcessMemInfo {
	m, _ := p.MemoryPercent()
	mi, _ := p.MemoryInfo()
	return ProcessMemInfo{
		MemPercent: m,
		MemRss:     mi.RSS,
		MemVms:     mi.VMS,
		MemSwap:    mi.Swap,
	}
}
