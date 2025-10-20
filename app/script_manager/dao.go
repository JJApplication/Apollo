package script_manager

import (
	"context"
	"errors"
	"github.com/JJApplication/Apollo/database"
	"github.com/JJApplication/Apollo/logger"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// 数据库操作

func (app *ScriptTask) CollectionName() string {
	return "scriptTask"
}

func getTaskList() []ScriptTask {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return nil
	}
	var data []ScriptTask
	if err := mgm.Coll(&ScriptTask{}).SimpleFind(&data, bson.M{}); err != nil {
		logger.LoggerSugar.Errorf("%s getTaskList error:%v", ScriptManagerPrefix, err)
		return nil
	}

	return data
}

func getTaskListByName(scriptName string) []ScriptTask {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return nil
	}
	var data []ScriptTask
	if err := mgm.Coll(&ScriptTask{}).SimpleFind(&data, bson.M{"script": scriptName}); err != nil {
		logger.LoggerSugar.Errorf("%s getTaskList of [%s] error:%v", ScriptManagerPrefix, scriptName, err)
		return nil
	}

	return data
}

// 获取某条任务的状态详情
func getTaskStatus(uuid string) (ScriptTask, error) {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return ScriptTask{}, errors.New("failed to connect to mongo")
	}
	var data ScriptTask
	if err := mgm.Coll(&ScriptTask{}).FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&data); err != nil {
		logger.LoggerSugar.Errorf("%s getTaskStatus of [%s] error:%v", ScriptManagerPrefix, uuid, err)
		return ScriptTask{}, err
	}

	return data, nil
}

// 检查是否有正在运行的任务
func hasTaskRunning(scriptName string) bool {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return true
	}
	var data ScriptTask
	if err := mgm.Coll(&ScriptTask{}).FindOne(context.Background(), bson.M{"script": scriptName, "status": StatusRunning}).Decode(&data); err != nil {
		logger.LoggerSugar.Errorf("%s get Running task of [%s] error:%v", ScriptManagerPrefix, scriptName, err)
		return true
	}

	return false
}

// 开启任务
func startTask(scriptName string) (string, error) {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return "", errors.New("failed to connect to mongo")
	}
	var data = ScriptTask{
		UUID:      uuid.NewString(),
		StartTime: time.Now().Unix(),
		Script:    scriptName,
		Status:    StatusRunning,
	}
	if err := mgm.Coll(&ScriptTask{}).Create(&data); err != nil {
		logger.LoggerSugar.Errorf("%s start task: %s error:%v", ScriptManagerPrefix, scriptName, err)
		return "", err
	}

	return data.UUID, nil
}

// 更新任务的状态
func updateTaskStatus(status, uuid, errMsg string) error {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return errors.New("failed to connect to mongo")
	}
	// 先查找
	var data ScriptTask
	if err := mgm.Coll(&ScriptTask{}).FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&data); err != nil {
		logger.LoggerSugar.Errorf("%s find task of [%s] error:%v", ScriptManagerPrefix, uuid, err)
		return errors.New("failed to find task: " + uuid)
	}
	data.Status = status
	data.Error = errMsg
	if status != StatusRunning {
		data.EndTime = time.Now().Unix()
	}

	if err := mgm.Coll(&ScriptTask{}).Update(&data); err != nil {
		logger.LoggerSugar.Errorf("%s update task status of [%s] error:%v", ScriptManagerPrefix, uuid, err)
		return errors.New("failed to update task: " + uuid)
	}
	return nil
}

func deleteTask(uuid string) error {
	// 如果任务正在运行 则停止并删除
	scriptLock.Lock()
	defer scriptLock.Unlock()

	for _, sc := range scriptManager {
		if sc.uuid == uuid && sc.status == StatusRunning {
			// 继续执行 不允许删除
			return nil
		}
	}

	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return errors.New("failed to connect to mongo")
	}

	if _, err := mgm.Coll(&ScriptTask{}).DeleteOne(context.Background(), bson.M{"uuid": uuid}); err != nil {
		return errors.New("failed to delete task: " + uuid)
	}
	
	return nil
}
