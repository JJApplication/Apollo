package indicator_manager

import (
	"github.com/JJApplication/Apollo/kv"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/model"
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
	"go.etcd.io/bbolt"
)

func IndicatorNetwork() []model.SystemNetwork {
	var data []model.SystemNetwork
	err := kv.KV.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("NETWORK")).ForEach(func(k, v []byte) error {
			var s model.SystemNetwork
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

func IndicatorNetworkRun() {
	net := utils.CalcNetwork()
	if err := kv.AddQuick(kv.KV, "NETWORK", net.JSON()); err != nil {
		logger.LoggerSugar.Errorf("%s get indicator of NETWORK error: %s", ManagerPrefix, err.Error())
		return
	}
}
