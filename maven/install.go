package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func Install(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in install function.\n") // debug

	// run the clean install command for the specified
	// project directory
	installCommand := exec.Command(
		"mvn",
		"clean",
		"install",
		"-s",
		settingsFile,
		"-f",
		projectDirectory,
		"-Dmaven.test.skip=true", // we don't want to run tests
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
			"-s",
			settingsFile,
			"-f",
			projectDirectory,
			mavenOpts,
			"-Dmaven.test.skip=true", // we don't want to run tests
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

	return outputBytes, nil
}
