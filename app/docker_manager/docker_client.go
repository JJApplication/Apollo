/*
Project: Apollo docker_client.go
Created: 2022/2/18 by Landers
*/

package docker_manager

import (
	"fmt"
	"sync"
	"time"

	docker "github.com/docker/docker/client"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
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
		DockerCli = getClient()
	})
}

func getClient() *docker.Client {
	client, err := docker.NewClientWithOpts(
		docker.WithHost(config.ApolloConf.CI.DockerHost),
		docker.WithTimeout(time.Duration(config.ApolloConf.CI.DockerTimeout)*time.Second),
		docker.WithVersion(config.ApolloConf.CI.DockerAPIVersion),
	)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s failed to create docker client: %s", DockerManager, err.Error()))
		return nil
	}
	logger.Logger.Info(DockerManager + " init docker client success")
	logger.Logger.Info(DockerManager + " connected to host: " + config.ApolloConf.CI.DockerHost)
	return client
}
