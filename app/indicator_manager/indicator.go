package indicator_manager

import (
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
)

// 操作系统的指标聚合

const (
	ManagerPrefix = "[Indicator Manager]"
)

func GetSystemInfo() model.SystemInfo {
	p, f, v := utils.CalcPlatform()
	return model.SystemInfo{
		Kernel:   utils.CalcKernel(),
		Platform: p,
		Family:   f,
		Version:  v,
	}
}
