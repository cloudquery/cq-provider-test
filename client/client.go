package client

import (
	"github.com/hashicorp/go-hclog"
)

type Configuration struct {
	Accounts []Account `hcl:"account,block"`
}

type Account struct {
	Name      string   `hcl:"name,label"`
	Id        string   `hcl:"id"`
	Regions   []string `hcl:"regions,optional"`
	Resources []string `hcl:"resources,optional"`
}

type TestClient struct {
	L hclog.Logger
}

func (t TestClient) Logger() hclog.Logger {
	return t.L
}

func (Configuration) Example() string {
	return `
  configuration {
    account "1" {
      id = "testid"
      regions = ["asdas"]
      resources = ["ab", "c"]
    }
  }`
}
