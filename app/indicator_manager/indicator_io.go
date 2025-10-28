package indicator_manager

import (
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
	"go.etcd.io/bbolt"
)

func IndicatorIO() []model.SystemIO {
	var data []model.SystemIO
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("IO")).ForEach(func(k, v []byte) error {
			var s model.SystemIO
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

func IndicatorIORun() {
	io := utils.CalcIO()
	if err := kv.AddQuick(kv.KV, "IO", io.JSON()); err != nil {
		logger.LoggerSugar.Errorf("%s get indicator of IO error: %s", ManagerPrefix, err.Error())
		return
	}
}
