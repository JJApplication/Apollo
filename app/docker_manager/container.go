/*
Project: Apollo container.go
Created: 2022/2/21 by Landers
*/

package docker_manager

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types/container"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
)

// 容器操作

// ContainerList 返回容器列表 默认返回全部状态
func ContainerList() ([]types.Container, error) {
	list, err := DockerCli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	return list, err
}

func ContainerNetworkCreate(network ContainerNetwork) (string, error) {
	if network.Name == "" {
		return "", errors.New("network name is empty")
	}
	driver := "bridge"
	if network.Host {
		driver = "host"
	}
	net, err := DockerCli.NetworkCreate(context.Background(), network.Name, types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         driver,
	})

	return net.ID, err
}

func ContainerCreate(c ContainerConfig) (container.ContainerCreateCreatedBody, error) {

	return DockerCli.ContainerCreate(context.Background(), &container.Config{
		Image:        c.ImageName,
		ExposedPorts: convertPortExpose(c),
	}, &container.HostConfig{
		LogConfig:    container.LogConfig{},
		PortBindings: convertPortBind(c),
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
		Resources:   container.Resources{},
		Mounts:      convertMountBind(c),
		NetworkMode: container.NetworkMode(c.Network.Name),
	},
		nil,
		nil,
		c.ContainerName,
	)
}

func ContainerCreateExec(c string) (types.IDResponse, error) {
	return DockerCli.ContainerExecCreate(context.Background(), c, types.ExecConfig{})
}

func ContainerRemove(id string) error {
	return DockerCli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{Force: true})
}

func ContainerInfo(id string) (types.ContainerJSON, error) {
	return DockerCli.ContainerInspect(context.Background(), id)
}

func ContainerStart(id string) error {
	return DockerCli.ContainerStart(context.Background(), id, types.ContainerStartOptions{
		CheckpointID:  "",
		CheckpointDir: "",
	})
}

func ContainerStop(id string) error {
	return DockerCli.ContainerStop(context.Background(), id, nil)
}

func ContainerRestart(id string) error {
	return DockerCli.ContainerRestart(context.Background(), id, nil)
}

func ContainerKill(id string) error {
	return DockerCli.ContainerKill(context.Background(), id, "9")
}

func ContainerPause(id string) error {
	return DockerCli.ContainerPause(context.Background(), id)
}

func ContainerResume(id string) error {
	return DockerCli.ContainerUnpause(context.Background(), id)
}

func ContainerRename(id string, name string) error {
	return DockerCli.ContainerRename(context.Background(), id, name)
}

func ContainerLogs(id string) (io.ReadCloser, error) {
	return DockerCli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStdout: false,
		ShowStderr: false,
		Since:      "",
		Until:      "",
		Timestamps: false,
		Follow:     false,
		Tail:       "",
		Details:    false,
	})
}

// GetContainerIDByName 根据容器名称获取容器ID
func GetContainerIDByName(name string) string {
	list, err := DockerCli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return ""
	}
	for _, c := range list {
		for _, cn := range c.Names {
			if strings.Trim(cn, "/") == name {
				return c.ID
			}
		}
	}

	return ""
}

// GetContainerNetworkIP 获取网卡的ip地址
// 默认使用第一个即可
func GetContainerNetworkIP(name string) []string {
	netInfo, err := DockerCli.NetworkInspect(context.Background(), name, types.NetworkInspectOptions{})
	if err != nil {
		return nil
	}

	config := netInfo.IPAM.Config
	// 解析IP地址
	var res []string
	for _, c := range config {
		res = append(res, c.Gateway)
	}

	return res
}
