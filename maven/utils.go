package maven

import (
	"fmt"
	"os"
)

func MoveToProjectDirectory(projectDirectory string) (string, error) {
	// first get the current working directory,
	// so that we can return there when done
	originalDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf(
			"couldn't get current directory:\n%s\n",
			err.Error(),
		)
		return "", err
	}

	// then check if the supplied project directory
	// exists.
	if _, err := os.Stat(projectDirectory); os.IsNotExist(err) {
		fmt.Printf(
			"project directory does not exist:\n%s\n",
			err.Error(),
		)
		return "", err
	}

	// since we verified that it does exist,
	// now change to that directory
	err = os.Chdir(projectDirectory)
	if err != nil {
		fmt.Printf(
			"couldn't change to the project directory:\n%s\n",
			err.Error(),
		)
		return "", err
	}

	return originalDirectory, nil
}

func ChangeDirectory(directory string) error {
	// since we verified that it does exist,
	// now change to that directory
	err := os.Chdir(directory)
	if err != nil {
		fmt.Printf(
			"couldn't change back to the current directory:\n%s\n",
			err.Error(),
		)
		return err
	}
	return nil
}
