/*
   Create: 2024/1/11
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package docker_manager

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
)

const (
	LocalHost = "127.0.0.1"
	TCP       = "tcp"
	UDP       = "udp"
)

type ContainerConfig struct {
	ImageName     string
	ContainerName string
	PortBind      []ContainerPortConfig // Host: Port
	MountBind     []ContainerMount
	Network       ContainerNetwork
}

type ContainerPortConfig struct {
	Listen   string // 80/tcp
	HostIP   string // 默认127.0.0.1
	HostPort string
}

type ContainerMount struct {
	Type     string
	ReadOnly bool
	Source   string
	Target   string
}

type ContainerNetwork struct {
	Host bool
	Name string
}

func convertPortBind(c ContainerConfig) nat.PortMap {
	pm := nat.PortMap{}
	for _, config := range c.PortBind {
		pm[nat.Port(config.Listen)] = []nat.PortBinding{
			{HostIP: config.HostIP, HostPort: config.HostPort},
		}
	}

	return pm
}

// 默认bind的端口全部放开
func convertPortExpose(c ContainerConfig) map[nat.Port]struct{} {
	pm := map[nat.Port]struct{}{}
	for _, config := range c.PortBind {
		pm[nat.Port(config.Listen)] = struct{}{}
	}

	return pm
}

func convertMountBind(c ContainerConfig) []mount.Mount {
	var mm []mount.Mount
	for _, containerMount := range c.MountBind {
		mm = append(mm, mount.Mount{
			Type:     mount.Type(containerMount.Type),
			Source:   containerMount.Source,
			Target:   containerMount.Target,
			ReadOnly: containerMount.ReadOnly,
		})
	}

	return mm
}
