package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Deploy(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in deploy function.\n") // debug

	originalDir, err := MoveToProjectDirectory(projectDirectory)
	if err != nil {
		return nil, err
	}

	// run the deploy command from the specified
	// project directory
	deployCommand := exec.Command(
		"mvn",
		"deploy",
		"-Dmaven.test.skip=true", // we don't want to run tests
	)

	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		deployCommand = exec.Command(
			"mvn",
			"deploy",
			"-Dmaven.test.skip=true", // we don't want to run tests
			mavenOpts,
		)
	}
	output, err := deployCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = deployCommand.Start()
	if err != nil {
		return nil, err
	}

	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		return nil, err
	}

	deployCommand.Wait()

	err = ChangeDirectory(originalDir)
	if err != nil {
		return nil, err
	}

	return outputBytes, nil
}
