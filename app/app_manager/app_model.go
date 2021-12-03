/*
Project: dirichlet app_model.go
Created: 2021/11/20 by Landers
*/

package app_manager

import (
	"github.com/gookit/validate"
	"github.com/landers1037/dirichlet/logger"
)

// App model for app
type App struct {
	Name          string `json:"name" validate:"required" bson:"name"`
	ID            string `json:"id" validate:"required" bson:"id"`
	Type          string `json:"type" bson:"type"`                     // service | middleware
	ReleaseStatus string `json:"release_status" bson:"release_status"` // published | pending | testing
	EngDes        string `json:"eng_des" bson:"eng_des"`
	CHSDes        string `json:"chs_des" bson:"chs_des"`

	// 管理项
	ManageCMD CMD `json:"manage_cmd" bson:"manage_cmd"`
	// 元数据
	Meta Meta `json:"meta" bson:"meta"`
	// 动态依赖配置
	RunData RunData `json:"run_data" bson:"run_data"`
}

// CMD 服务的管理脚本
type CMD struct {
	Start     string `json:"start" bson:"start"`
	Stop      string `json:"stop" bson:"stop"`
	Restart   string `json:"restart" bson:"restart"`
	ForceKill string `json:"force_kill" bson:"force_kill"`
	Check     string `json:"check" bson:"check"`
}

// Meta 服务的元数据
type Meta struct {
	Author      string   `json:"author" bson:"author"`
	Language    []string `json:"language" bson:"language"`
	CreateDate  string   `json:"create_date" bson:"create_date"`
	Version     string   `json:"version" bson:"version"`
	DynamicConf bool     `json:"dynamic_conf" bson:"dynamic_conf"` // 是否需要生成配置文件
	ConfType    string   `json:"conf_type" bson:"conf_type"`       // nginx | gunicorn
	ConfPath    string   `json:"conf_path" bson:"conf_path"`       // 支持绝对和相对路径
}

// RunData 运行时依赖
type RunData struct {
	Envs       []string `json:"envs" bson:"envs"` // just like `Name=Diri`
	Ports      []int    `json:"ports" bson:"ports"`
	RandomPort bool     `json:"random_port" bson:"random_port"` // if using random port
	Host       string   `json:"host" bson:"host"`               // always must be localhost
}

// Validate 适用于model的检查器
func (app *App) Validate() bool {
	if !validate.Struct(app).Validate() {
		logger.Logger.Error(appCodeMsg("app validate failed") + validate.Struct(app).Errors.String())
		return false
	}

	return true
}
