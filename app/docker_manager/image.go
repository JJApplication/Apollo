/*
Project: dirichlet image.go
Created: 2022/2/21 by Landers
*/

package docker_manager

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
)

// images操作

// ImageList 列举全部镜像
func ImageList() ([]types.ImageSummary, error) {
	list, err := DockerCli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func ImageInfo(id string) (types.ImageInspect, error) {
	info, _, err := DockerCli.ImageInspectWithRaw(context.Background(), id)
	return info, err
}

func ImagePull(ref string) (io.ReadCloser, error) {
	return DockerCli.ImagePull(context.Background(), ref, types.ImagePullOptions{})
}

func ImageRemove(id string) ([]types.ImageDeleteResponseItem, error) {
	return DockerCli.ImageRemove(context.Background(), id, types.ImageRemoveOptions{Force: true})
}