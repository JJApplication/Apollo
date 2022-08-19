/*
Project: Apollo dao.go
Created: 2021/12/4 by Landers
*/

package app_manager

import (
	"context"
	"fmt"

	"github.com/JJApplication/Apollo/database"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// DaoAPP 数据库操作类
type DaoAPP struct {
	mgm.DefaultModel `bson:",inline"`
	App              `json:"app" bson:"app"`
}

const CollectName = "microservice"
const PersistFile = "apolloServices.bson"

// CollectionName 定义存储名称 不能与app冲突
func (app *DaoAPP) CollectionName() string {
	return CollectName
}

// SaveToDB 批量插入
func SaveToDB() {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	var data DaoAPP
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		if !checkExist(bson.M{"app.meta.name": key}) {
			data = DaoAPP{
				App: value.(App),
			}
			e := mgm.Coll(&DaoAPP{}).Create(&data)
			if e != nil {
				logger.Logger.Error(fmt.Sprintf("%s insert app %s to db failed: %s", APPManagerPrefix, key, e.Error()))
			} else {
				logger.Logger.Info(fmt.Sprintf("%s insert app %s to db", APPManagerPrefix, key))
			}
		} else {
			// 更新操作
			var data DaoAPP
			v := value.(App)
			mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": key}).Decode(&data)
			if len(v.Meta.RunData.Ports) > 0 {
				data.Meta.RunData.Ports = v.Meta.RunData.Ports
			}
			if len(v.Meta.RunData.Envs) > 0 {
				data.Meta.RunData.Envs = v.Meta.RunData.Envs
			}
			data.Meta.ID = v.Meta.ID
			data.Meta.Type = v.Meta.Type
			data.Meta.CHSDes = v.Meta.CHSDes
			data.Meta.EngDes = v.Meta.EngDes
			data.Meta.ReleaseStatus = v.Meta.ReleaseStatus
			data.Meta.ManageCMD = v.Meta.ManageCMD
			data.Meta.Meta = v.Meta.Meta

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
		if checkExist(bson.M{"app.meta.name": key}) {
			var data DaoAPP
			err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": key}).Decode(&data)
			if err == nil {
				app := value.(App)
				if app.Meta.RunData.Ports == nil || len(app.Meta.RunData.Ports) == 0 {
					app.Meta.RunData.Ports = data.Meta.RunData.Ports
				}
				if app.Meta.RunData.Envs == nil || len(app.Meta.RunData.Envs) == 0 {
					app.Meta.RunData.Envs = data.Meta.RunData.Envs
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
	err = utils.SaveBson(&data, PersistFile)
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
	err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": app}).Decode(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime port to db failed: %s", APPManagerPrefix, app, err.Error()))
		return
	}
	data.Meta.RunData.Ports = port
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
	err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": app.Meta.Name}).Decode(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Meta.Name, err.Error()))
		return
	}
	data.Meta.RunData = app.Meta.RunData
	err = mgm.Coll(&DaoAPP{}).Update(&data)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Meta.Name, err.Error()))
	}
	logger.Logger.Info(fmt.Sprintf("%s save [%s] runtime data to db", APPManagerPrefix, app.Meta.Name))
}

func checkExist(filter interface{}) bool {
	if e := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), filter).Err(); e != nil {
		return false
	}
	return true
}
