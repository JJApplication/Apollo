/*
Project: Apollo dao.go
Created: 2021/12/4 by Landers
*/

package app_manager

import (
	"context"

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

// SaveToDB 批量插入数据到数据库中
// 每次插入都视为全量更新
func SaveToDB() {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	apps := utils.NewSet()
	// 数据 -> 数据库
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		apps.Add(value.(App).Meta.Name)

		if !checkExist(bson.M{"app.meta.name": key}) {
			var data DaoAPP
			data = DaoAPP{
				App: value.(App),
			}
			e := mgm.Coll(&DaoAPP{}).Create(&data)
			if e != nil {
				logger.LoggerSugar.Errorf("%s insert app %s to db failed: %s", APPManagerPrefix, key, e.Error())
			} else {
				logger.LoggerSugar.Infof("%s insert app %s to db", APPManagerPrefix, key)
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
				logger.LoggerSugar.Errorf("%s update [%s] data to db failed: %s", APPManagerPrefix, key, err.Error())
			} else {
				logger.LoggerSugar.Infof("%s update [%s] data to db", APPManagerPrefix, key)
			}
		}
		return true
	})
	// 数据库自查
	var data []DaoAPP
	if err := mgm.Coll(&DaoAPP{}).SimpleFind(&data, bson.M{}); err != nil {
		logger.LoggerSugar.Errorf("%s load from db failed: %s", APPManagerPrefix, err.Error())
		return
	}

	for _, d := range data {
		if !apps.Contains(d.Meta.Name) {
			logger.LoggerSugar.Infof("%s [%s] unload from db", APPManagerPrefix, d.Meta.Name)
			if _, err := mgm.Coll(&DaoAPP{}).DeleteOne(context.Background(), bson.M{"app.meta.name": d.Meta.Name}); err != nil {
				logger.LoggerSugar.Errorf("%s [%s] unload from db failed: %s", APPManagerPrefix, d.Meta.Name, err.Error())
				return
			}
		}
	}
}

// FirstLoad 在异常情况下重启，先从mongo拿数据保存后续再刷新
//
// 可以从缓存中继承的数据runtime运行时数据
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
				if app.Meta.Runtime.Pid == "" && data.Meta.Runtime.Pid != "" {
					app.Meta.Runtime.Pid = data.Meta.Runtime.Pid
				}
				if app.Meta.Runtime.Ports == nil && len(data.Meta.Runtime.Ports) > 0 {
					app.Meta.Runtime.Ports = data.Meta.Runtime.Ports
				}
				if app.Meta.Runtime.StopOperation != data.Meta.Runtime.StopOperation {
					app.Meta.Runtime.StopOperation = data.Meta.Runtime.StopOperation
				}
				APPManager.APPManagerMap.Store(key, app)
				logger.LoggerSugar.Infof("%s first load %s runtime data to appCache", APPManagerPrefix, key)
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
		logger.LoggerSugar.Errorf("%s persist from db failed: %s", APPManagerPrefix, err.Error())
		return
	}
	err = utils.SaveBson(&data, PersistFile)
	if err != nil {
		logger.LoggerSugar.Errorf("%s persist from db failed: %s", APPManagerPrefix, err.Error())
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
		logger.LoggerSugar.Errorf("%s save [%s] runtime port to db failed: %s", APPManagerPrefix, app, err.Error())
		return
	}
	data.Meta.RunData.Ports = port
	data.Meta.Runtime.Ports = port
	err = mgm.Coll(&DaoAPP{}).Update(&data)
	if err != nil {
		logger.LoggerSugar.Errorf("%s save [%s] runtime port to db failed: %s", APPManagerPrefix, app, err.Error())
	}
	logger.LoggerSugar.Infof("%s save [%s] runtime port to db", APPManagerPrefix, app)
}

func SaveRuntimeData(app App) {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return
	}
	var data DaoAPP
	err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": app.Meta.Name}).Decode(&data)
	if err != nil {
		logger.LoggerSugar.Errorf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Meta.Name, err.Error())
		return
	}
	data.Meta.Runtime = app.Meta.Runtime
	err = mgm.Coll(&DaoAPP{}).Update(&data)
	if err != nil {
		logger.LoggerSugar.Errorf("%s save [%s] runtime data to db failed: %s", APPManagerPrefix, app.Meta.Name, err.Error())
	}
	logger.LoggerSugar.Infof("%s save [%s] runtime data to db", APPManagerPrefix, app.Meta.Name)
}

// GetRuntimePortApp 从数据库中获取存在运行时动态端口的服务
func GetRuntimePortApp() []App {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return []App{}
	}

	var dynamicPortApp []App
	APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
		if checkExist(bson.M{"app.meta.name": key}) {
			var data DaoAPP
			err := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), bson.M{"app.meta.name": key}).Decode(&data)
			if err == nil {
				app := value.(App)
				if app.Meta.RunData.RandomPort {
					if len(data.Meta.RunData.Ports) > 0 {
						dynamicPortApp = append(dynamicPortApp, app)
					}
				}
			}
		}
		return true
	})

	return dynamicPortApp
}

func checkExist(filter interface{}) bool {
	if e := mgm.Coll(&DaoAPP{}).FindOne(context.Background(), filter).Err(); e != nil {
		return false
	}
	return true
}
