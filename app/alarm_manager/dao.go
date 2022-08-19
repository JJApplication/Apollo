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

	"github.com/kamva/mgm/v3"
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
	return res, err
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
	return res, err
}

func getAlarm(id string) (Alarm, error) {
	var res Alarm
	err := mgm.Coll(&Alarm{}).SimpleFind(&res, bson.M{"_id": id})
	return res, err
}

func deleteAlarm(id string) error {
	_, err := mgm.Coll(&Alarm{}).DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
