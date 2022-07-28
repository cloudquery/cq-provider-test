package client

import (
	"github.com/hashicorp/go-hclog"
)

type Configuration struct {
	Accounts []Account `yaml:"account"`
}

type Account struct {
	Name      string   `yaml:"name"`
	Id        string   `yaml:"id"`
	Regions   []string `yaml:"regions,omitempty"`
	Resources []string `yaml:"resources,omitempty"`
}

type TestClient struct {
	L hclog.Logger
}

func (t TestClient) Logger() hclog.Logger {
	return t.L
}

func (Configuration) Example() string {
	return `
#account:
#  name: "1"
#  id: testid
#  regions:
#    - asdas
`
}
