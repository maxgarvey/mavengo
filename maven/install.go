package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Install(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in install function.\n") // debug

	originalDir, err := MoveToProjectDirectory(projectDirectory)
	if err != nil {
		return nil, err
	}

	// run the clean install command from the specified
	// project directory
	installCommand := exec.Command(
		"mvn",
		"clean",
		"install",
		"-Dmaven.test.skip=true", // we don't want to run tests
		"-s",
		settingsFile,
	)

	// add localcache option flag
	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		// fmt.Printf(mavenOpts) // debug
		installCommand = exec.Command(
			"mvn",
			"clean",
			"install",
			"-Dmaven.test.skip=true", // we don't want to run tests
			"-s",
			settingsFile,
			mavenOpts,
		)
	}
	output, err := installCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = installCommand.Start()
	if err != nil {
		return nil, err
	}
	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Printf("err:\n%v\n", err)
		return nil, err
	}
	installCommand.Wait()

	err = ChangeDirectory(originalDir)
	if err != nil {
		return nil, err
	}

	return outputBytes, nil
}
