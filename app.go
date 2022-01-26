package main

import (
	"fmt"

	"github.com/jrsmile/docker-Dasboard/adapters/docker"
	"github.com/jrsmile/docker-Dasboard/resources"

	"github.com/gin-gonic/gin"
)

func main() {
	containerEngine := docker.NewDockerContainerEngine()

	fmt.Printf("Started...\n")

	gin.SetMode(gin.ReleaseMode)

	r := resources.SetupRouter(containerEngine)
	r.Run() // listen and serve on 0.0.0.0:8080
}
