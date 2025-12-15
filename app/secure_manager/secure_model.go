package secure_manager

import "github.com/kamva/mgm/v3"

// 安全模块模型

const (
	TypeSSH  = "ssh"
	TypeHttp = "http"
	TypePing = "ping"
)

type SecureAudit struct {
	mgm.DefaultModel `bson:",inline"`
	SecureType       string `bson:"secureType" json:"secureType"`
	BlockIP          string `bson:"blockIP" json:"blockIP"`
	BlockStart       string `bson:"blockStart" json:"blockStart"` // 屏蔽日期
	BlockTime        string `bson:"blockTime" json:"blockTime"`   // 屏蔽时长
	Remark           string `bson:"remark" json:"remark"`
}
