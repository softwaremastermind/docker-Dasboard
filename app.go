package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type Event struct {
	ID      int
	Message string
}

func streamEvents(c *gin.Context) {
	listener := make(chan interface{})
	go sendEvents(listener)
	c.Stream(func(w io.Writer) bool {
		event := <-listener
		if event != nil {
			c.SSEvent("message", event)
			return true
		}
		return false
	})
}

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

	r.GET("/stream", streamEvents)

	r.GET("/status", func(c *gin.Context) {
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
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

func sendEvents(listener chan interface{}) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	msgs, errs := cli.Events(ctx, types.EventsOptions{})
	var i int
	for {
		i ++
		select {
			case err := <-errs:
				fmt.Println(err)
				listener <- Event{i, "update"}
			case msg := <-msgs:
				fmt.Println(msg)
				listener <- Event{i, "update"}
		}
	}
}