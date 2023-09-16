/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import "github.com/shirou/gopsutil/process"

type ProcessIO struct {
	ReadCount  uint64 `json:"readCount"`
	WriteCount uint64 `json:"writeCount"`
	ReadBytes  uint64 `json:"readBytes"`
	WriteBytes uint64 `json:"writeBytes"`
}

// CalcProcessIO 获取进程IO读写
func CalcProcessIO(p *process.Process) ProcessIO {
	io, _ := p.IOCounters()
	return ProcessIO{
		ReadCount:  io.ReadCount,
		WriteCount: io.WriteCount,
		ReadBytes:  io.ReadBytes,
		WriteBytes: io.WriteBytes,
	}
}
