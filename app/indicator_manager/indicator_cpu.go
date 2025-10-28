package indicator_manager

import (
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
	"go.etcd.io/bbolt"
)

// CPU指标

func IndicatorCPU() []model.SystemCPU {
	var data []model.SystemCPU
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("CPU")).ForEach(func(k, v []byte) error {
			var s model.SystemCPU
			err := json.Unmarshal(v, &s)
			if err != nil {
				return err
			}
			data = append(data, s)
			return nil
		})
	})
	if err != nil {
		return nil
	}

	return data
}

func IndicatorCPURun() {
	if err := kv.AddQuick(kv.KV, "CPU", (&model.SystemCPU{
		Percent: utils.CalcCpuLoad(),
	}).JSON()); err != nil {
		logger.LoggerSugar.Errorf("%s get indicator of CPU error: %s", ManagerPrefix, err.Error())
		return
	}
}
