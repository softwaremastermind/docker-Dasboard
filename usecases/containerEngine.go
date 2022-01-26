package usecases

import "github.com/jrsmile/docker-Dasboard/entities"

type ContainerEngine interface {
	GetContainer() ([]entities.Container, error)
}
