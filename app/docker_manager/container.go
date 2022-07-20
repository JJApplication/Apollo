/*
Project: Apollo container.go
Created: 2022/2/21 by Landers
*/

package docker_manager

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
)

// 容器操作

func ContainerList() ([]types.Container, error) {
	list, err := DockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ContainerCreate(c string) (types.IDResponse, error) {
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
