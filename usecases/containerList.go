package usecases

import (
	"log"

	"github.com/jrsmile/docker-Dasboard/entities"
)

func RetrieveContainer(containerEngine ContainerEngine) []entities.Container {
	if cont, err := containerEngine.GetContainer(); err == nil {
		return cont
	} else {
		log.Fatal(err)
		return nil
	}
}
