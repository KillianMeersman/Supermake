package supermake

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/moby/moby/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerEnvironment struct {
	Image      string
	Entrypoint []string
}

type imageURL struct {
	Registry string
	Name     string
	Tag      string
}

func parseImageURL(url string) (*imageURL, error) {
	parts := strings.Split(url, "/")

	registry := "docker.io"
	name := strings.Join(parts, "/")
	if len(parts) > 1 {
		if strings.Contains(parts[0], ".") {
			registry = parts[0]
			name = strings.Join(parts[1:], "/")
		}
	}

	tag := "latest"
	parts = strings.SplitN(parts[len(parts)-1], ":", 2)
	if len(parts) > 1 {
		tag = parts[1]
		name = strings.TrimSuffix(name, fmt.Sprintf(":%s", tag))
	}

	return &imageURL{
		Registry: registry,
		Name:     name,
		Tag:      tag,
	}, nil
}

func (u *imageURL) String() string {
	return fmt.Sprintf("%s/%s:%s", u.Registry, u.Name, u.Tag)
}

func (d *DockerEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command Command) error {
	mobyClient, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	imageURL, err := parseImageURL(d.Image)
	if err != nil {
		return err
	}

	dockerWorkingDir := "/supermake/"

	pull := true
	images, err := mobyClient.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}
searchimages:
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == d.Image {
				pull = false
				break searchimages
			}
		}
	}

	if pull {
		execCtx.Logger.Debug("pulling image", "image", imageURL.String())

		output, err := mobyClient.ImagePull(ctx, imageURL.String(), types.ImagePullOptions{
			Platform: "linux/amd64",
			All:      false,
		})
		if err != nil {
			return err
		}
		// data, err := ioutil.ReadAll(output)
		// if err != nil {
		// 	return err
		// }

		// fmt.Print(string(data))

		defer output.Close()
	} else {
		execCtx.Logger.Debug("image already available", "image", imageURL.String())
	}

	entrypoint, cmd := ParseInterpreterCommand(d.Entrypoint, command.GetShellCommands())

	config := &container.Config{
		Image:      imageURL.String(),
		Entrypoint: []string{entrypoint},
		Cmd:        cmd,
		WorkingDir: dockerWorkingDir,
		Tty:        true,
		Env:        execCtx.EnvVars.EnvStrings(),
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: execCtx.WorkingDir,
				Target: dockerWorkingDir,
			},
		},
		NetworkMode: "host",
	}
	networkConfig := &network.NetworkingConfig{}
	platform := &v1.Platform{}

	dockerContainer, err := mobyClient.ContainerCreate(ctx, config, hostConfig, networkConfig, platform, "")
	if err != nil {
		return err
	}
	defer mobyClient.ContainerRemove(ctx, dockerContainer.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	execCtx.Logger.Info("starting container", "image", imageURL.String())
	err = mobyClient.ContainerStart(ctx, dockerContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	logReader, err := mobyClient.ContainerLogs(ctx, dockerContainer.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return err
	}
	defer logReader.Close()

	reader := bufio.NewReader(logReader)
	for {
		line, err := reader.ReadBytes('\n')

		lineStr := string(bytes.TrimRight(line, "\n\r"))
		execCtx.Logger.Info(lineStr)
		if err != nil {
			if err != io.EOF {
				execCtx.Logger.Error(fmt.Sprintf("error streaming logs from container: %s", err))
				return err
			}
			break
		}
	}

	waitChan, errChan := mobyClient.ContainerWait(ctx, dockerContainer.ID, container.WaitConditionNotRunning)
	select {
	case wait := <-waitChan:
		err := wait.Error
		if err != nil {
			return fmt.Errorf("error running container: %s", err)
		}

		statusCode := wait.StatusCode
		if statusCode != 0 {
			return fmt.Errorf("container returned non-zero exit code: %d", statusCode)
		}
	case err := <-errChan:
		return err
	}

	execCtx.Logger.Info("container finished", "image", imageURL.String())
	if err != nil {
		return err
	}

	return nil
}
