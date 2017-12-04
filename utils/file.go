package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Ensure that the project directory exists.
func EnsureDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// Find the project root directory. Recursively traverse the parents of the
// current directory, until a directory containing a bib.json file is found or
// the file system root directory is encountered.
func GetProjectRootDirectory (dir string) (string, error) {
	return "", nil
}

// Initialize Git repository in the specified directory.
func InitGitRepo (dir string) (error) {
	// TODO revise so that we can specify the cwd for the command
	_, err := exec.Command("git", "init").Output()
	if err != nil {
		return err
	}
	return nil
}

// Determine if the specified directory is a Git repository
func IsGitRepo (dir string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == ".git" {
			return true, nil
		}
	}
	return false, nil
}
