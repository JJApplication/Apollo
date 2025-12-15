package secure_manager

import (
	"github.com/JJApplication/Apollo/database"
	"github.com/JJApplication/Apollo/logger"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *SecureAudit) CollectionName() string {
	return "SecureAudit"
}

func GetSecureList() []SecureAudit {
	if !database.MongoPing {
		logger.Logger.Warn("failed to connect to mongo")
		return nil
	}
	var data []SecureAudit
	if err := mgm.Coll(&SecureAudit{}).SimpleFind(&data, bson.M{}); err != nil {
		logger.LoggerSugar.Errorf("%s getSecureList error:%v", SecureManagerPrefix, err)
		return nil
	}

	return data
}

// 存储审计信息到数据库
//
// TODO 当IP已经存在且时间更新时，刷新时间戳
func saveSecureAudit(data *SecureAudit) {
	if !database.MongoPing {
		return
	}
	if err := mgm.Coll(&SecureAudit{}).Create(data); err != nil {
		logger.LoggerSugar.Errorf("%s saveSecureAudit error:%v", SecureManagerPrefix, err)
	}
}
