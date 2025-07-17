/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"github.com/JJApplication/Apollo/model"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/process"
)

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

func CalcIO() model.SystemIO {
	var totalReadBytes, totalWriteBytes uint64
	var totalReadCount, totalWriteCount uint64
	var totalReadTime, totalWriteTime, ioTime uint64

	counters, err := disk.IOCounters()
	if err != nil {
		return model.SystemIO{}
	}
	for _, counter := range counters {
		totalReadCount += counter.ReadCount
		totalWriteCount += counter.WriteCount
		totalReadBytes += counter.ReadBytes
		totalWriteBytes += counter.WriteBytes
		totalReadTime += counter.ReadTime
		totalWriteTime += counter.WriteTime
		ioTime += counter.IoTime
	}

	return model.SystemIO{
		TotalReadBytes:  totalReadBytes,
		TotalWriteBytes: totalWriteBytes,
		TotalReadCount:  totalReadCount,
		TotalWriteCount: totalWriteCount,
		TotalReadTime:   totalReadTime,
		TotalWriteTime:  totalWriteTime,
		IOTime:          ioTime,
	}
}
