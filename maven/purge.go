package maven

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func PurgeLocal(localCache, projectDirectory, settingsFile string) ([]byte, error) {
	// fmt.Printf("in purge function.\n") // debug

	// run the clean purge command for the specified local cache
	purgeCommand := exec.Command(
		"mvn",
		"dependency:purge-local-repository",
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
		purgeCommand = exec.Command(
			"mvn",
			"dependency:purge-local-repository",
			"-s",
			settingsFile,
			"-f",
			projectDirectory,
			mavenOpts,
			"-Dmaven.test.skip=true", // we don't want to run tests
		)
	}
	output, err := purgeCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = purgeCommand.Start()
	if err != nil {
		return nil, err
	}
	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Printf("err:\n%v\n", err)
		return nil, err
	}
	purgeCommand.Wait()

	return outputBytes, nil
}
