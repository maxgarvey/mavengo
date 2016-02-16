package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Deploy(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in deploy function.\n") // debug

	// run the deploy command for the specified
	// project directory
	deployCommand := exec.Command(
		"mvn",
		"deploy",
		"-f",
		projectDirectory,
		"-s",
		settingsFile,
		"-Dmaven.test.skip=true", // we don't want to run tests
	)

	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		deployCommand = exec.Command(
			"mvn",
			"deploy",
			"-f",
			projectDirectory,
			"-s",
			settingsFile,
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

	return outputBytes, nil
}
