package utils

import "github.com/shirou/gopsutil/load"

// ProcLoad 计算系统负载1, 5, 15min
func ProcLoad() (float64, float64, float64) {
	info, err := load.Avg()
	if err != nil {
		return 0, 0, 0
	}
	return info.Load1, info.Load5, info.Load15
}
