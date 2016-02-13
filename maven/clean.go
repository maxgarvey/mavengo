package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Clean(localCache, projectDirectory string) ([]byte, error) {
	// fmt.Printf("in clean function.\n") // debug

	originalDir, err := MoveToProjectDirectory(projectDirectory)
	if err != nil {
		return nil, err
	}

	// run the clean command from the specified
	// project directory
	cleanCommand := exec.Command(
		"mvn",
		"clean",
	)

	// add localcache option flag
	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		// fmt.Printf(mavenOpts) // debug
		cleanCommand = exec.Command(
			"mvn",
			"clean",
			mavenOpts,
		)
	}
	output, err := cleanCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cleanCommand.Start()
	if err != nil {
		return nil, err
	}
	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Printf("err:\n%v\n", err)
		return nil, err
	}
	cleanCommand.Wait()

	err = ChangeDirectory(originalDir)
	if err != nil {
		return nil, err
	}

	return outputBytes, nil
}
