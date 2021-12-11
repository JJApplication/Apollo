/*
Project: dirichlet dao.go
Created: 2021/12/4 by Landers
*/

package app_manager

import (
	"context"
	"fmt"

	"github.com/kamva/mgm/v3"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// DaoAPP 数据库操作类
type DaoAPP struct {
	mgm.DefaultModel `bson:",inline"`
	App
}

// SaveToDB 批量插入
func SaveToDB() {
	var data DaoAPP
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		if !checkExist(bson.M{"app.name": key}) {
			data = DaoAPP{
				App: value.(App),
			}
			res, e := mgm.Coll(&DaoAPP{}).InsertOne(context.Background(), data)
			if e != nil {
				logger.Logger.Error(fmt.Sprintf("%s insert app %s to db failed: %s", APPManagerPrefix, key, e.Error()))
			} else {
				logger.Logger.Info(fmt.Sprintf("%s insert app %s to db, resID: %v", APPManagerPrefix, key, res.InsertedID))
			}
		}
		return true
	})
}

// Persist 持久化为bson数据
// todo 序列化bson文件
func Persist() {
	var data []DaoAPP
	err := mgm.Coll(&DaoAPP{}).SimpleFind(&data, bson.M{})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s persist %s from db failed: %s", APPManagerPrefix, err.Error()))
		return
	}
	err = utils.SaveBson(&data, "app.bson")
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s persist %s from db failed: %s", APPManagerPrefix, err.Error()))
	}
}

func checkExist(filter interface{}) bool {
	if e := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), filter).Err(); e != nil {
		return false
	}
	return true
}
