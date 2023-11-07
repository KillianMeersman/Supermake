package supermake

import (
	"context"
	"fmt"
	"os/user"

	"github.com/KillianMeersman/Supermake/pkg/supermake/docker"
	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
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

func (d *DockerEnvironment) Execute(ctx context.Context, execCtx ExecutorContext, command Command) error {
	mobyClient, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	imageURL, err := docker.ParseImageURL(d.Image)
	if err != nil {
		return err
	}

	dockerWorkingDir := "/supermake/"

	err = docker.GetOrPullImage(ctx, mobyClient, imageURL)
	if err != nil {
		return err
	}

	entrypoint, cmd := ParseInterpreterCommand(execCtx.EnvVars, d.Entrypoint, command.GetShellCommands())

	user, err := user.Current()
	if err != nil {
		execCtx.Logger.Panic(err.Error())
	}

	config := &container.Config{
		Image:      imageURL.String(),
		Entrypoint: []string{entrypoint},
		Cmd:        cmd,
		WorkingDir: dockerWorkingDir,
		Tty:        true,
		Env:        execCtx.EnvVars.EnvStrings(),
		User:       fmt.Sprintf("%s:%s", user.Uid, user.Gid),
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
	defer docker.RemoveContainer(ctx, mobyClient, dockerContainer.ID)

	execCtx.Logger.Debug("starting container", "image", imageURL.String())
	err = mobyClient.ContainerStart(ctx, dockerContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	logStream, err := docker.StreamContainerLogs(ctx, mobyClient, dockerContainer.ID)
	if err != nil {
		return err
	}
	defer logStream.Close()
	log.StreamReaderNewLines(execCtx.Logger.Info, logStream)

	waitChan, errChan := mobyClient.ContainerWait(ctx, dockerContainer.ID, container.WaitConditionNotRunning)
	select {
	case wait := <-waitChan:
		err := wait.Error
		if err != nil {
			return fmt.Errorf("error running container: %s", err)
		}

		statusCode := wait.StatusCode
		if statusCode != 0 {
			logs, err := docker.GetContainerLogs(ctx, mobyClient, dockerContainer.ID)
			if err != nil {
				execCtx.Logger.Panic(err.Error())
			}
			execCtx.Logger.Error(string(logs))
			return fmt.Errorf("container returned non-zero exit code: %d", statusCode)
		}
	case err := <-errChan:
		return err
	}

	execCtx.Logger.Debug("container finished", "image", imageURL.String())
	if err != nil {
		return err
	}

	return nil
}
