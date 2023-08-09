package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
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

func StreamContainerLogs(ctx context.Context, client *client.Client, container string) (io.ReadCloser, error) {
	return client.ContainerLogs(ctx, container, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
}