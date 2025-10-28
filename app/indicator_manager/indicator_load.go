package indicator_manager

import (
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
	"go.etcd.io/bbolt"
)

func IndicatorLoad() []model.SystemLoad {
	var data []model.SystemLoad
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("LOAD")).ForEach(func(k, v []byte) error {
			var s model.SystemLoad
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

func IndicatorLoadRun() {
	m1, m5, m15 := utils.ProcLoad()
	if err := kv.AddQuick(kv.KV, "LOAD", (&model.SystemLoad{
		Minute1:  m1,
		Minute5:  m5,
		Minute15: m15,
	}).JSON()); err != nil {
		logger.LoggerSugar.Errorf("%s get indicator of LOAD error: %s", ManagerPrefix, err.Error())
		return
	}
}
