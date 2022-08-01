package client

import (
	"github.com/hashicorp/go-hclog"
)

type Configuration struct {
	Accounts []Account `hcl:"account,block" yaml:"account"`
}

type Account struct {
	Name      string   `hcl:"name,label" yaml:"name"`
	Id        string   `hcl:"id" yaml:"id"`
	Regions   []string `hcl:"regions,optional" yaml:"regions,omitempty"`
	Resources []string `hcl:"resources,optional" yaml:"resources,omitempty"`
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
