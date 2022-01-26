package docker

import (
	"context"

	"github.com/jrsmile/docker-Dasboard/entities"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerContainerEngine struct {
	client client.Client
}

func (d *DockerContainerEngine) GetContainer() ([]entities.Container, error) {
	if containers, err := d.client.ContainerList(context.TODO(), types.ContainerListOptions{}); err == nil {
		var result []entities.Container = make([]entities.Container, 0)
		for _, c := range containers {
			result = append(result, entities.Container{Id: c.ID, Image: c.Image, Name: c.Names[0]})
		}
		return result, nil
	} else {
		return nil, err
	}
}

func NewDockerContainerEngine() *DockerContainerEngine {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return &DockerContainerEngine{client: *cli}
}
