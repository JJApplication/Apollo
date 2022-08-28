/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package alarm_manager
package alarm_manager

import (
	"context"

	"github.com/JJApplication/Apollo/utils"
	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo操作
// mongo默认为全局初始化 直接通过mgm调用

type Alarm struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Level            string `json:"level" bson:"level"`
	Message          string `json:"message" bson:"message"`
}

func (app *Alarm) CollectionName() string {
	return "alarm"
}

func getAllAlarm() ([]Alarm, error) {
	var res []Alarm
	err := mgm.Coll(&Alarm{}).SimpleFind(&res, bson.M{})
	if err != nil {
		return nil, err
	}
	return toLocalTime(res), err
}

func getTopNAlarm() ([]Alarm, error) {
	var res []Alarm
	op := new(options.FindOptions)
	op.SetLimit(TopN)
	op.SetSort(bson.M{"created_at": -1})
	err := mgm.Coll(&Alarm{}).SimpleFind(&res, bson.M{}, op)
	if err != nil {
		return nil, err
	}
	return toLocalTime(res), err
}

func getAlarm(id string) (Alarm, error) {
	var res Alarm
	err := mgm.Coll(&Alarm{}).SimpleFind(&res, bson.M{"_id": id})
	res.CreatedAt = utils.TimeToLocal(res.CreatedAt)
	res.UpdatedAt = utils.TimeToLocal(res.UpdatedAt)
	return res, err
}

func deleteAlarm(id string) error {
	_, err := mgm.Coll(&Alarm{}).DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func getCountAlarm() int64 {
	count, err := mgm.Coll(&Alarm{}).CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0
	}
	return count
}

func deleteLastN() (int64, error) {
	// 按照日期的创建逆序排列 最新的N条
	count := getCountAlarm()
	if count >= AlarmSizeLimit {
		// 取最近的第N+1条数据
		var res []Alarm
		op := new(options.FindOptions)
		op.SetLimit(AlarmSizeTrim)
		op.SetSort(bson.M{"created_at": -1})
		err := mgm.Coll(&Alarm{}).SimpleFind(&res, bson.M{}, op)
		if err != nil {
			return 0, err
		}
		// 取条件日期
		date := res[len(res)-1].CreatedAt
		delRes, err := mgm.Coll(&Alarm{}).DeleteMany(context.Background(), bson.M{"created_at": bson.M{operator.Lt: date}})
		return delRes.DeletedCount, err
	}
	return 0, nil
}

func toLocalTime(res []Alarm) []Alarm {
	var result []Alarm
	for _, r := range res {
		r.CreatedAt = utils.TimeToLocal(r.CreatedAt)
		r.UpdatedAt = utils.TimeToLocal(r.UpdatedAt)

		result = append(result, r)
	}
	return result
}
