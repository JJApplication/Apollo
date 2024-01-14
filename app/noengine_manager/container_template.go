/*
   Create: 2024/1/11
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package noengine_manager

import (
	"fmt"
	"github.com/JJApplication/Apollo/app/docker_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/containerd/containerd/pkg/cri/util"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	EngineRoot = "$engine" // NoEngine的根目录
	// Root hostPath配置为$root/dir映射为$NoEngineAPP/APP_NAME
	// 默认直接解析为相对路径
	// 绝对路径标识符? 例如?/root -> /root
	Root          = "@root" // NoEngineAPP/$APP的根目录环境变量
	APP           = "@app"  // 微服务$APP的根目录
	AbsRoot       = "?"
	NoEngineImage = "openresty/openresty"
	OpenrestyConf = "/usr/local/openresty/nginx/conf/nginx.conf"
)

// 生成容器模板
//
// 所有容器使用公用模板路径$NoEngineAPPConf
// 在HOST上会默认监听80端口，使用https时需要增加443端口的监听
// 转发到实际微服务监听的端口，当转发的端口不存在时异常由Nginx处理
// 何时使用随机端口? 解析的APP配置HostPort存在时则不使用随机端口
var (
	preInitVolume []NoEngineVolume
	preInitPorts  = []NoEnginePort{
		{HostPort: "", InnerPort: "80"},
	}
)

// 创建微服务的NoEngine容器 出错时返回错误，否则会创建名为name的容器用于查询
func createContainer(template NoEngineTemplate) (error, NoEngineTemplate) {
	var temp = template
	preInitVolume = []NoEngineVolume{
		{HostPath: NoEngineAPPConf, InnerPath: OpenrestyConf},
	}
	// 格式化volume
	for i, _ := range temp.Volumes {
		hp := temp.Volumes[i].HostPath
		// 解析标识符
		if hp == "" {
			// 目录为空时 直接映射的就是NoEngine/$APP的根目录
			temp.Volumes[i].HostPath = filepath.Join(NoEngineAPP, temp.ServerName, hp)
		} else if strings.HasPrefix(hp, Root) {
			realhp := strings.TrimPrefix(hp, Root)
			temp.Volumes[i].HostPath = filepath.Join(NoEngineAPP, temp.ServerName, realhp)
		} else if strings.HasPrefix(hp, APP) {
			realhp := strings.TrimPrefix(hp, APP)
			temp.Volumes[i].HostPath = filepath.Join(config.ApolloConf.APPRoot, temp.ServerName, realhp)
		} else if strings.HasPrefix(hp, EngineRoot) {
			realhp := strings.TrimPrefix(hp, EngineRoot)
			temp.Volumes[i].HostPath = filepath.Join(NoEngineAPP, realhp)
		} else if strings.HasPrefix(hp, AbsRoot) {
			p, e := filepath.Abs(strings.TrimPrefix(hp, AbsRoot))
			if e != nil {
				temp.Volumes[i].HostPath = filepath.Join(NoEngineAPP, hp)
			}
			temp.Volumes[i].HostPath = p
		} else {
			temp.Volumes[i].HostPath = filepath.Join(config.ApolloConf.APPRoot, temp.ServerName, hp)
		}
	}

	// 预设volume
	temp.Volumes = append(temp.Volumes, preInitVolume...)
	// 预设ports
	if len(temp.Ports) <= 0 {
		var tp []NoEnginePort
		_ = util.DeepCopy(&tp, preInitPorts)
		for i, _ := range tp {
			tp[i].HostPort = strconv.Itoa(randomPort())
		}
		temp.Ports = tp
	} else {
		// 遍历配置的端口 host为空的设置随机端口
		for i, t := range temp.Ports {
			if t.HostPort == "" {
				temp.Ports[i].HostPort = strconv.Itoa(randomPort())
			}
		}
	}

	// 使用配置完毕的模板启动容器
	cf := buildContainerConfig(temp)
	resp, err := docker_manager.ContainerCreate(cf)
	if err != nil {
		return err, temp
	}
	return docker_manager.ContainerStart(resp.ID), temp
}

func buildContainerConfig(t NoEngineTemplate) docker_manager.ContainerConfig {
	var cf docker_manager.ContainerConfig
	cf.ContainerName = t.ServerName
	cf.ImageName = NoEngineImage
	for _, port := range t.Ports {
		proto := port.Proto
		if proto == "" {
			proto = docker_manager.TCP
		}
		cf.PortBind = append(cf.PortBind, docker_manager.ContainerPortConfig{
			Listen:   fmt.Sprintf("%s/%s", port.InnerPort, proto),
			HostIP:   docker_manager.LocalHost,
			HostPort: port.HostPort,
		})
	}

	for _, volume := range t.Volumes {
		cf.MountBind = append(cf.MountBind, docker_manager.ContainerMount{
			Type:     "bind",
			ReadOnly: false,
			Source:   volume.HostPath,
			Target:   volume.InnerPath,
		})
	}

	cf.Network.Name = config.ApolloConf.APPBridge

	return cf
}
