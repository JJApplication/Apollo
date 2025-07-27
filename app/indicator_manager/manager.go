package indicator_manager

import (
	"time"

	"github.com/JJApplication/Apollo/logger"
)

// 定时任务用于存储最新数据

func AllocIndicator() {
	logger.LoggerSugar.Infof("%s init alloc indicator", ManagerPrefix)
	ticker := time.NewTicker(time.Minute * 15)
	go func() {
		for {
			select {
			case <-ticker.C:
				IndicatorLoadRun()
				IndicatorCPURun()
				IndicatorMemRun()
				IndicatorIORun()
				IndicatorNetworkRun()
			}
		}
	}()
}
