package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"encoding/json"

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

	fmt.Printf("Started...\n")

    gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/status", func(c *gin.Context) {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		u, err := json.Marshal(containers)
        if err != nil {
            panic(err)
        }

		c.DataFromReader(http.StatusOK,
        int64(len(string(u))), gin.MIMEJSON, strings.NewReader(string(u)), nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}