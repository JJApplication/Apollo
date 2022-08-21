/*
Project: Apollo docker_client.go
Created: 2022/2/18 by Landers
*/

package docker_manager

import (
	"sync"
	"time"

	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	docker "github.com/docker/docker/client"
)

// docker的客户端
//ubuntu18.04 配置docker service
// vi /etc/systemd/system/multi-user.target.wants/docker.service
// ExecStart=... -H tcp://0.0.0.0:2375
// systemctl daemon-reload
// systemctl restart docker

var DockerCli *docker.Client

const (
	DockerManager = "[DockerManager]"
)

// InitDockerClient 初始化Docker连接
func InitDockerClient() {
	once := sync.Once{}
	once.Do(func() {
		DockerCli = createClient()
	})
}

// 初始化docker客户端
func createClient() *docker.Client {
	client, err := docker.NewClientWithOpts(
		docker.WithHost(config.ApolloConf.CI.DockerHost),
		docker.WithTimeout(time.Duration(config.ApolloConf.CI.DockerTimeout)*time.Second),
		docker.WithVersion(config.ApolloConf.CI.DockerAPIVersion),
	)
	if err != nil {
		logger.LoggerSugar.Errorf("%s failed to create docker client: %s", DockerManager, err.Error())
		return nil
	}
	logger.LoggerSugar.Infof("%s init docker client success", DockerManager)
	logger.LoggerSugar.Infof("%s connected to host: %s", DockerManager, config.ApolloConf.CI.DockerHost)
	return client
}

func getClient() *docker.Client {
	return DockerCli
}
