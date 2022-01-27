package resources

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jrsmile/docker-Dasboard/usecases"
)

// content holds our static web server content.
//go:embed static
// var staticFS embed.FS

//go:embed static/index.html
var index_html string

func SetupRouter(containerEngine usecases.ContainerEngine) *gin.Engine {
	r := gin.Default()
	r.GET("/status", createStatusHandler(containerEngine))
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, index_html)
	})
	return r
}

// func mustFS() http.FileSystem {
// 	sub, err := fs.Sub(staticFS, "public")

// 	if err != nil {
// 		panic(err)
// 	}

// 	return http.FS(sub)
// }
func createStatusHandler(containerEngine usecases.ContainerEngine) gin.HandlerFunc {

	containers := usecases.RetrieveContainer(containerEngine)
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, containers)
	}
}
