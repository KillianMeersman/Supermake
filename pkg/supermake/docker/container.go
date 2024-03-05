package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/moby/moby/client"
)

func GetContainerLogs(ctx context.Context, client *client.Client, container string) ([]byte, error) {
	logReader, err := client.ContainerLogs(ctx, container, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return nil, err
	}
	defer logReader.Close()
	return io.ReadAll(logReader)
}

func StreamContainerLogs(ctx context.Context, client *client.Client, containerID string) (io.ReadCloser, error) {
	return client.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
}

func RemoveContainer(ctx context.Context, client *client.Client, id string) error {
	return client.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		Force: true,
	})
}
