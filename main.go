package main

import (
	"fmt"
	"os"

	"github.com/maxgarvey/mavengo/maven"
)

func main() {
	fmt.Printf("in main function.\n")

	output, err := maven.Install("/Users/mgarve/maven_stuff/artifactory_stuff/trunk")
	if err != nil {
		fmt.Printf(
			"error running maven install:\n%s\n",
			err.Error(),
		)
		os.Exit(1)
	}

	fmt.Printf("output:\n%s\n", string(output))
}
