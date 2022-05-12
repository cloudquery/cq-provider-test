package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
	"github.com/cloudquery/cq-provider-test/resources"
)

func main() {
	outputPath := "./docs"
	if err := docs.GenerateDocs(resources.Provider(), outputPath, true); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
