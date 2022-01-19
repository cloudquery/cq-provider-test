package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/cloudquery/cq-provider-test/resources"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "test",
		Provider: resources.Provider(),
	})
}
