package supermake

import (
	"context"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/moby/moby/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type CommandExecutor interface {
	Execute(ctx context.Context, workingDir string, command Command) error
}

type DockerEnvironment struct {
	Image      string
	Entrypoint string
}

func (d *DockerEnvironment) Execute(ctx context.Context, workingDir string, command Command) error {
	client, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	dockerWorkingDir := "/supermake/"
	config := &container.Config{
		Image:      d.Image,
		Entrypoint: []string{d.Entrypoint},
		Cmd:        command.GetShellCommands(),
		WorkingDir: dockerWorkingDir,
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: workingDir,
				Target: dockerWorkingDir,
			},
		},
		NetworkMode: "host",
	}
	networkConfig := &network.NetworkingConfig{}
	platform := &v1.Platform{}

	container, err := client.ContainerCreate(ctx, config, hostConfig, networkConfig, platform, "test")
	if err != nil {
		return err
	}
	defer client.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	return nil
}

type LocalEnvironment struct {
	Entrypoint string
}

func (l *LocalEnvironment) Execute(ctx context.Context, workingDir string, command Command) error {

	for _, command := range command.GetShellCommands() {
		parts := strings.Split(command, " ")
		command := parts[0]
		args := parts[1:]
		if l.Entrypoint != "" {
			command = l.Entrypoint
			args = parts
		}
		cmd := exec.Command(command, args...)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
