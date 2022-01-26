package resources_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jrsmile/docker-Dasboard/entities"
	"github.com/jrsmile/docker-Dasboard/resources"
)

type ContainerEngineMock struct{}

func (c *ContainerEngineMock) GetContainer() ([]entities.Container, error) {
	return []entities.Container{
		{
			Id:    "test-1",
			Name:  "test Container 1",
			Image: "example.com/testimage:latest",
			State: "running",
			Port:  8080,
		},
	}, nil
}
func TestStatus(t *testing.T) {
	containerEngine := ContainerEngineMock{}
	router := resources.SetupRouter(&containerEngine)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/status", nil)
	if err != nil {
		t.Fatal("Error: cannot initialize GET /status http request")
	}
	router.ServeHTTP(recorder, request)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Error("/status should return status code 200")
	}
}
