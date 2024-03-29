/*
Create: 2022/8/20
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package alarm_manager
package alarm_manager

import (
	"github.com/JJApplication/Apollo/logger"
)

const (
	TopN           = 10
	AlarmManager   = "[Alarm Manager]"
	AlarmSizeLimit = 50 // 超过50开始删除
	AlarmSizeTrim  = 20 // 保留最近的20条
)

// GetAllAlarm 获取全部告警信息
// 按照时间排序 从新到旧
func GetAllAlarm() []Alarm {
	res, err := getAllAlarm()
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to get all Alarms, error: %s", AlarmManager, err.Error())
		return nil
	}
	return res
}

// GetTopNAlarm 获取最新的10条告警
func GetTopNAlarm() []Alarm {
	res, err := getTopNAlarm()
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to get topN Alarms, error: %s", AlarmManager, err.Error())
		return nil
	}
	return res
}

func GetAlarmInfo(id string) Alarm {
	res, err := getAlarm(id)
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to get Alarm [%s], error: %s", AlarmManager, id, err.Error())
		return Alarm{}
	}
	return res
}

// DeleteAlarm 根据告警id删除告警信息
func DeleteAlarm(id string) error {
	err := deleteAlarm(id)
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to delete Alarm [%s], error: %s", AlarmManager, id, err.Error())
	}
	return nil
}

// ClearAlarmLimit 告警清除任务
// 满足条件 数量大于指定条数 则删除最旧的M条
func ClearAlarmLimit() error {
	count, err := deleteLastN()
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to delete Last Alarm, error: %s", AlarmManager, err.Error())
		return err
	}
	logger.LoggerSugar.Infof("%s delete Last Alarm success, count %d", AlarmManager, count)
	return nil
}
