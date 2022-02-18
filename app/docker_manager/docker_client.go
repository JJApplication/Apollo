/*
Project: dirichlet docker_client.go
Created: 2022/2/18 by Landers
*/

package docker_manager

// docker的客户端

import (
	"fmt"
	"sync"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/landers1037/dirichlet/logger"
)

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
	client, err := docker.NewClient("")
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("%s failed to create docker client: %s", DockerManager, err.Error()))
		return nil
	}
	logger.Logger.Info(DockerManager + " init docker client success")
	return client
}
