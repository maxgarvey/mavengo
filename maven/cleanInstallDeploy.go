package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func CleanInstallDeploy(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in install function.\n") // debug

	// run the clean install deploy command for the specified
	// project directory
	cleanInstallCommand := exec.Command(
		"mvn",
		"clean",
		"install",
		"deploy",
		"-f",
		projectDirectory,
		"-s",
		settingsFile,
		"-Dmaven.test.skip=true", // we don't want to run tests
	)

	// add localcache option flag
	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		// fmt.Printf(mavenOpts) // debug
		cleanInstallCommand = exec.Command(
			"mvn",
			"clean",
			"install",
			"deploy",
			"-f",
			projectDirectory,
			"-s",
			settingsFile,
			"-Dmaven.test.skip=true", // we don't want to run tests
			mavenOpts,
		)
	}
	output, err := cleanInstallCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cleanInstallCommand.Start()
	if err != nil {
		return nil, err
	}
	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Printf("err:\n%v\n", err)
		return nil, err
	}
	cleanInstallCommand.Wait()

	return outputBytes, nil
}
