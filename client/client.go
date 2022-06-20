package client

import (
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
)

type Configuration struct {
	Accounts []Account `hcl:"account,block" yaml:"account"`

	requestedFormat cqproto.ConfigFormat
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

func NewConfiguration(f cqproto.ConfigFormat) *Configuration {
	return &Configuration{
		requestedFormat: f,
	}
}

func (c Configuration) Example() string {
	switch c.requestedFormat {
	case cqproto.ConfigHCL:
		return `
  configuration {
    account "1" {
      id = "testid"
      regions = ["asdas"]
      resources = ["ab", "c"]
    }
  }`
	default:
		return `
#account:
#  name: "1"
#  id: testid
#  regions:
#    - asdas
`
	}
}

func (c Configuration) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
