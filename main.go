package main

import (
	"fmt"
	"os"

	"github.com/maxgarvey/mavengo/maven"
)

func main() {
	fmt.Printf("in main function.\n")

	// local stuff, for proof of concept
	mavenBinary := "/usr/local/bin/mvn"
	settingsFile := "/Users/mgarve/.m2/settings.xml"

	output, err := maven.Install(
		mavenBinary,
		"",
		"./",
		settingsFile,
	)
	if err != nil {
		fmt.Printf(
			"error running maven install:\n%s\n",
			err.Error(),
		)
		os.Exit(1)
	}

	fmt.Printf("output:\n%s\n", string(output))
}
