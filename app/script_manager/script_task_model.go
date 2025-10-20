package script_manager

import "github.com/kamva/mgm/v3"

// 脚本任务模型

const (
	StatusRunning = "running"
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "stop"
)

type ScriptTask struct {
	mgm.DefaultModel `bson:",inline"`
	UUID             string `json:"uuid" bson:"uuid"`
	StartTime        int64  `json:"start_time" bson:"start_time"`
	EndTime          int64  `json:"end_time" bson:"end_time"`
	Script           string `json:"script" bson:"script"`
	Error            string `json:"error" bson:"error"`
	Status           string `json:"status" bson:"status"`
}
