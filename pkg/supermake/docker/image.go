package docker

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/KillianMeersman/Supermake/pkg/supermake/log"
	"github.com/docker/docker/api/types"
	"github.com/moby/moby/client"
)

const defaultImageRegistry = "docker.io/library/"

var imagePulls = sync.Map{}

type ImageURL struct {
	Registry string
	Name     string
	Tag      string
}

func (u *ImageURL) String() string {
	return fmt.Sprintf("%s/%s:%s", u.Registry, u.Name, u.Tag)
}

func ParseImageURL(url string) (*ImageURL, error) {
	parts := strings.Split(url, "/")

	registry := "docker.io"
	path := parts
	if len(parts) > 1 {
		if strings.Contains(parts[0], ".") {
			registry = parts[0]
			path = parts[1:]
		}
	} else {
		path = []string{"library"}
		path = append(path, parts...)
	}

	tag := "latest"
	parts = strings.SplitN(path[len(path)-1], ":", 2)
	if len(parts) > 1 {
		tag = parts[1]
		path[len(path)-1] = parts[0]
	}

	return &ImageURL{
		Registry: registry,
		Name:     strings.Join(path, "/"),
		Tag:      tag,
	}, nil
}

func PullImage(ctx context.Context, client *client.Client, url *ImageURL) error {
	log.DefaultLogger().Info("pulling image", "image", url.String())
	output, err := client.ImagePull(ctx, url.String(), types.ImagePullOptions{
		Platform: "linux/amd64",
		All:      false,
	})
	if err != nil {
		return err
	}
	_, err = io.ReadAll(output)
	if err != nil {
		return err
	}
	defer output.Close()

	log.DefaultLogger().Info("pulled image", "image", url.String())
	return nil
}

func SearchLocalImages(ctx context.Context, client *client.Client, query string) (string, error) {
	images, err := client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return "", err
	}

	query = strings.TrimPrefix(query, defaultImageRegistry)

	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == query {
				return tag, nil
			}
		}
	}

	return "", nil
}

func GetOrPullImage(ctx context.Context, client *client.Client, url *ImageURL) error {
	lock, _ := imagePulls.LoadOrStore(url.String(), &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
	defer lock.(*sync.Mutex).Unlock()

	image, err := SearchLocalImages(ctx, client, url.String())
	if err != nil {
		return err
	}

	if image != "" {
		return nil
	}

	return PullImage(ctx, client, url)
}
