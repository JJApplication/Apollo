/*
Project: dirichlet dao.go
Created: 2021/12/4 by Landers
*/

package app_manager

import (
	"context"
	"fmt"

	"github.com/kamva/mgm/v3"
	"github.com/landers1037/dirichlet/database"
	"github.com/landers1037/dirichlet/logger"
	"github.com/landers1037/dirichlet/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// DaoAPP 数据库操作类
type DaoAPP struct {
	mgm.DefaultModel `bson:",inline"`
	App
}

func (app *DaoAPP) CollectionName() string {
	return "app"
}

// SaveToDB 批量插入
func SaveToDB() {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
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
		} else {
			// 更新操作
			var data DaoAPP
			v := value.(App)
			mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.name": key}).Decode(&data)
			if len(v.RunData.Ports) > 0 {
				data.RunData.Ports = v.RunData.Ports
			}
			if len(v.RunData.Envs) > 0 {
				data.RunData.Envs = v.RunData.Envs
			}
			data.ID = v.ID
			data.Type = v.Type
			data.CHSDes = v.CHSDes
			data.EngDes = v.EngDes
			data.ReleaseStatus = v.ReleaseStatus
			data.ManageCMD = v.ManageCMD
			data.Meta = v.Meta

			err := mgm.Coll(&DaoAPP{}).Update(&data)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("%s update [%s] data to db failed: %s", APPManagerPrefix, key, err.Error()))
			} else {
				logger.Logger.Info(fmt.Sprintf("%s update [%s] data to db", APPManagerPrefix, key))
			}
		}
		return true
	})
}

// FirstLoad 在异常情况下重启，先从mongo拿数据保存后续再刷新
func FirstLoad() {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		if checkExist(bson.M{"app.name": key}) {
			var data DaoAPP
			err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.name": key}).Decode(&data)
			if err == nil {
				app := value.(App)
				if app.RunData.Ports == nil || len(app.RunData.Ports) == 0 {
					app.RunData.Ports = data.RunData.Ports
				}
				if app.RunData.Envs == nil || len(app.RunData.Envs) == 0 {
					app.RunData.Envs = data.RunData.Envs
				}
				APPManager.APPManagerMap.Store(key, app)
				logger.Logger.Info(fmt.Sprintf("%s first load %s runtime data to appCache", APPManagerPrefix, key))
			}
		}
		return true
	})
}

// Persist 持久化为bson数据
// todo 序列化bson文件
func Persist() {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	var data []DaoAPP
	err := mgm.Coll(&DaoAPP{}).SimpleFind(&data, bson.M{})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s persist from db failed: %s", APPManagerPrefix, err.Error()))
		return
	}
	err = utils.SaveBson(&data, "app.bson")
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s persist from db failed: %s", APPManagerPrefix, err.Error()))
	}
}

func SavePort(app string, port []int) {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	var data DaoAPP
	err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.name": app}).Decode(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime port to db failed: %s", APPManagerPrefix, app, err.Error()))
		return
	}
	data.RunData.Ports = port
	err = mgm.Coll(&DaoAPP{}).Update(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime port to db failed: %s", APPManagerPrefix, app, err.Error()))
	}
	logger.Logger.Info(fmt.Sprintf("%s save [%s] runtime port to db", APPManagerPrefix, app))
}

func SaveRuntimeData(app App) {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	var data DaoAPP
	err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.name": app.Name}).Decode(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Name, err.Error()))
		return
	}
	data.RunData = app.RunData
	err = mgm.Coll(&DaoAPP{}).Update(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Name, err.Error()))
	}
	logger.Logger.Info(fmt.Sprintf("%s save [%s] runtime data to db", APPManagerPrefix, app.Name))
}

func checkExist(filter interface{}) bool {
	if e := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), filter).Err(); e != nil {
		return false
	}
	return true
}
