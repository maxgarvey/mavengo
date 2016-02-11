package maven

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

func Deploy(localCache, projectDirectory string, envLock *sync.Mutex) ([]byte, error) {
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
	)
	output, err := deployCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// obtain environmental variable mutex
	envLock.Lock()
	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		os.Setenv("MAVEN_OPTS", mavenOpts)
	}
	err = deployCommand.Start()

	// return environmental variable mutex
	envLock.Unlock()
	if err != nil {
		return nil, err
	}
	deployCommand.Wait()

	err = ChangeDirectory(originalDir)
	if err != nil {
		return nil, err
	}

	outputBytes, err := ioutil.ReadAll(output)
	if err != nil {
		return nil, err
	}

	return outputBytes, nil
}
