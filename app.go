package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

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

		var buffer bytes.Buffer
		buffer.WriteString("[")

		for _, container := range containers {
			buffer.WriteString("{" + `"` + "id" + `"` + ":")
			buffer.WriteString(`"` + container.ID[:10] +`"`)
			buffer.WriteString("," + `"` + "name" + `"` + ":")
			buffer.WriteString(`"` + container.Names[0] + `"`)
			buffer.WriteString("," + `"` + "image" + `"` + ":")
			buffer.WriteString(`"` + container.Image + `"`)
			buffer.WriteString("," + `"` + "state" + `"` + ":")
			buffer.WriteString(`"` + container.State + `"`)
			buffer.WriteString("},")
		}
		buffer.Truncate(buffer.Len()-1)
		buffer.WriteString("]")
		c.DataFromReader(http.StatusOK,
        int64(len(buffer.String())), gin.MIMEJSON, strings.NewReader(buffer.String()), nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}