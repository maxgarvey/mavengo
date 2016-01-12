package maven

import (
	"fmt"
	// "os"
	"os/exec"
)

func Deploy(projectDirectory string) ([]byte, error) {
	fmt.Printf("in install function.\n") // debug

	originalDir, err := MoveToProjectDirectory(projectDirectory)
	if err != nil {
		return nil, err
	}

	// run the clean install command from the specified
	// project directory
	installCommand := exec.Command(
		"mvn",
		"deploy",
	)
	output, err := installCommand.CombinedOutput()
	if err != nil {
		return nil, err
	}

	err = ChangeDirectory(originalDir)
	if err != nil {
		return nil, err
	}

	return output, nil
}
