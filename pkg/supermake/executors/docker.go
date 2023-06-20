package executors

import (
	"context"
	"log"

	"github.com/KillianMeersman/Supermake/pkg/supermake/executables"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/moby/moby/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerEnvironment struct {
	Image      string
	Entrypoint string
}

func (d *DockerEnvironment) Execute(ctx context.Context, workingDir string, command executables.Command) error {
	client, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	dockerWorkingDir := "/supermake/"
	config := &container.Config{
		Image:        d.Image,
		Entrypoint:   []string{d.Entrypoint},
		Cmd:          command.GetShellCommands(),
		WorkingDir:   dockerWorkingDir,
		AttachStdout: true,
		AttachStderr: true,
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

	images, err := client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	pull := true
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
		log.Printf("pulling image '%s'", d.Image)
		_, err = client.ImagePull(ctx, d.Image, types.ImagePullOptions{})
		if err != nil {
			return err
		}
	} else {
		log.Printf("image '%s' already available, not pulling", d.Image)
	}

	container, err := client.ContainerCreate(ctx, config, hostConfig, networkConfig, platform, "test")
	if err != nil {
		return err
	}
	defer client.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	err = client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}
