package indicator_manager

import (
	"encoding/json"
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
	"go.etcd.io/bbolt"
)

func IndicatorMem() []model.SystemMem {
	var data []model.SystemMem
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("MEMORY")).ForEach(func(k, v []byte) error {
			var s model.SystemMem
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

func IndicatorMemRun() {
	mem := utils.CalcMem()
	if err := kv.AddQuick(kv.KV, "MEMORY", mem.JSON()); err != nil {
		logger.LoggerSugar.Errorf("%s get indicator of MEMORY error: %s", ManagerPrefix, err.Error())
		return
	}
}
