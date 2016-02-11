package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/maxgarvey/mavengo/maven"
)

func main() {
	fmt.Printf("in main function.\n")

	envLock := &sync.Mutex{}
	output, err := maven.Install(
		"",
		"./",
		envLock,
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
