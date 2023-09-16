/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"strings"
	"time"
)

type SysProc struct {
	PID            int32   `json:"pid"`
	CPUPercent     float64 `json:"cpuPercent"`
	ProcessMemInfo `json:",inline"`
	ProcessIO      `json:",inline"`
	NetConnections int   `json:"netConnections"`
	Threads        int32 `json:"threads"`
}

// CalcBoot 主机的启动时间
func CalcBoot() string {
	start, err := host.BootTime()
	if err != nil {
		return ""
	}
	t := time.Unix(int64(start), 0)
	return t.Local().Format("2006-01-02 15:04:05")
}

// CalcKernel 主机的内核信息
func CalcKernel() string {
	k, err := host.KernelVersion()
	if err != nil {
		return ""
	}
	return k
}

// CalcPlatform 主机平台信息
// 平台 家族 版本
func CalcPlatform() (string, string, string) {
	p, f, v, err := host.PlatformInformation()
	if err != nil {
		return "", "", ""
	}
	return p, f, v
}

// GetProcessCount 获取进程总数
func GetProcessCount() int {
	ps, err := process.Processes()
	if err != nil {
		return 0
	}
	return len(ps)
}

// GetProcess 获取全部进程
func GetProcess() []*process.Process {
	ps, err := process.Processes()
	if err != nil {
		return ps
	}
	return ps
}

// FilterProcess 过滤得到指定名称的进程
func FilterProcess(name string) *process.Process {
	ps, err := process.Processes()
	if err != nil {
		return nil
	}
	for _, p := range ps {
		processName, _ := p.Cmdline()
		if strings.Contains(processName, name) {
			return p
		}
	}
	return nil
}

// GetProcessInfo 获取进程PID
func GetProcessInfo(p *process.Process) string {
	return p.String()
}

func GetProcessPID(p *process.Process) int32 {
	if p != nil {
		return p.Pid
	}
	return 0
}

// GetProcessThreads 获取进程的线程数
func GetProcessThreads(p *process.Process) int32 {
	t, _ := p.NumThreads()
	return t
}
