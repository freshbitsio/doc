package bib

// Determine if the bib.json file exists in the specified directory.
func Exists (dir string) (bool, error) {
	return true, nil
}

// Find the bib.json file in the project root directory.
func FindRootBib () (string, error) {
	return "", nil
}

func Read () {}

func Write () {}
