package bib

import (
	"io/ioutil"
	"os"
)

func AddRecord (bib, id, title, year string) {

}

// Determine if the bib.json file exists in the specified directory.
func Exists (dir string) (bool, error) {
	return true, nil
}

// Find the bib.json file in the project root directory.
func FindRootBib () (string, error) {
	return "", nil
}

// Read the bib.json file
func Read () ([]byte, error) {
	return ioutil.ReadFile("bib.json")
}

// Write the bib.json file
func Write (data []byte) error {
	return ioutil.WriteFile("bib.json", data, os.FileMode(0666))
}
