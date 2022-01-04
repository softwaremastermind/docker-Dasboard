package main

import (
	//"net/http"
	"context"
	"bytes"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"

)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Started...:\n")

    gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}
		
		var buffer bytes.Buffer
		buffer.WriteString("{")

		for _, container := range containers {
			buffer.WriteString("{id: \"")
			buffer.WriteString(container.ID[:10])
			buffer.WriteString("\",name: \"")
			buffer.WriteString(container.Names[0])
			buffer.WriteString("\",image: \"")
			buffer.WriteString(container.Image)
			buffer.WriteString("\"},")
			//containerjson, err := cli.ContainerInspect(context.Background(), container.ID[:10])
			//if err != nil {
			//	panic(err)
			//}
		}
		buffer.WriteString("}")
		c.JSON(200, gin.H{"message": buffer.String(),})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}