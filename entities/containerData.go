package entities

type Container struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	State string `json:"state"`
	Port  int    `json:"port"`
}
