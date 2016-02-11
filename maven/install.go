package maven

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

func Install(localCache, projectDirectory string, envLock *sync.Mutex) ([]byte, error) {
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
	)
	output, err := installCommand.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// obtain environmental variable mutex
	envLock.Lock()
	if localCache != "" {
		mavenOpts := fmt.Sprintf(
			"-Dmaven.repo.local=%s", localCache)
		// fmt.Printf(mavenOpts) // debug
		os.Setenv("MAVEN_OPTS", mavenOpts)
	}
	err = installCommand.Start()
	// release mutex after command has started

	envLock.Unlock()
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
