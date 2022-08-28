/*
Project: Apollo time_util.go
Created: 2022/2/17 by Landers
*/

package utils

import (
	"time"
)

// 时间计算

const (
	TimeFormat    = "2006/1/2 15:04:05"
	TimeFormatDay = "2006/1/2"
	TimeForLogger = "2006-1-2 15:04:05"
	TimeBetterSep = "2006-1-2-15-04-05"
)

// TimeNowUnix 计算当前unix时间
func TimeNowUnix() int64 {
	return time.Now().Unix()
}

// TimeNowString 计算当前日志的年月日时分秒
func TimeNowString() string {
	return time.Now().Format(TimeFormat)
}

// TimeNowYearDay 计算当前日期的年月日
func TimeNowYearDay() string {
	return time.Now().Format(TimeFormatDay)
}

// TimeNowFormat 返回符合格式的时间戳
func TimeNowFormat(layout string) string {
	return time.Now().Format(layout)
}

// TimeNowBetterSep 形如"2006-1-2-15-04-05"
func TimeNowBetterSep() string {
	return time.Now().Format(TimeBetterSep)
}

// TimeCalcUnix 计算ms差
func TimeCalcUnix(last time.Time) int64 {
	return time.Since(last).Milliseconds()
}

// TimeCalcString 计算时间差的string
func TimeCalcString(last time.Time) string {
	return time.Since(last).String()
}

// TimeToLocal 时区转换
// 转换为当地时区 ->东8区
func TimeToLocal(t time.Time) time.Time {
	return t.Local()
}
