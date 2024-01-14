/*
   Create: 2023/11/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import "github.com/JJApplication/Apollo/utils"

// 用于启动NoEngine的命令模板
// 内部服务统一监听80端口 外部端口动态生成

type NoEngineTemplate struct {
	// 域名默认由nginx.conf文件配置不需要单独写在配置文件中
	ServerDomain string           `json:"serverDomain"` // 缺省域名
	ServerName   string           `json:"serverName"`   // 微服务名称
	Volumes      []NoEngineVolume `json:"volumes"`      // 映射卷
	Ports        []NoEnginePort   `json:"ports"`        // 映射端口 开启随机端口时hostPort为随机生成
}

type NoEngineVolume struct {
	HostPath  string `json:"hostPath"`
	InnerPath string `json:"innerPath"`
}

type NoEnginePort struct {
	HostPort  string `json:"hostPort"`
	InnerPort string `json:"innerPort"`
	Proto     string `json:"proto"` // 默认为tcp
}

// GenerateTemplate 生成空模板
func GenerateTemplate(domain, name string) string {
	return utils.PrettyJson(NoEngineTemplate{
		ServerDomain: domain,
		ServerName:   name,
		Volumes:      []NoEngineVolume{},
		Ports:        []NoEnginePort{},
	})
}
