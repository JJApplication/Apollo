/*
   Create: 2023/9/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"github.com/JJApplication/Apollo/model"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// CalcProcessNet 获取进程的网络堆栈
func CalcProcessNet(p *process.Process) int {
	conns, _ := p.Connections()
	return len(conns)
}

func CalcNetwork() model.SystemNetwork {
	currentCounter, err := net.IOCounters(false)
	if err != nil {
		return model.SystemNetwork{}
	}
	if len(currentCounter) == 0 {
		return model.SystemNetwork{}
	}
	n := currentCounter[0]
	return model.SystemNetwork{
		ByteRecv:    n.BytesRecv,
		ByteSent:    n.BytesSent,
		PacketsRecv: n.PacketsRecv,
		PacketsSent: n.PacketsSent,
	}
}
