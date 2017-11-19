//-----------------------------------------------------------------------------
// Init command module
// A project is a file system directory that contains a BibJSON file named
// bib.json in the top level directory. This module is used to initialize a
// project by creating or resetting that top level file.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------

package cmd

import (
	"bufio"
	"doc/data"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/spf13/cobra"
	"io/ioutil"
	//"log"
	"os"
	"os/user"
	//"path/filepath"
	"time"
	//"path"
	"os/exec"
	//"strings"
	//"bytes"
)

var ProjectPath string

// User configuration file
type userConfiguration struct {
	fullname string
	email string
	location string
	organization string
	citationStyle string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty project bibliography or reinitialize an existing one",
	Long: `
  Create an empty project bibliography or reinitialize an existing one.
  The bibliography will be stored in BibJSON format. For more information
  about the BibJSON format, see http://okfnlabs.org/bibjson

    Collection - The collection name
    Created - The date the bibliography was created
    Description - A description of the bibliography
    From -
    Id - a unique identifier for the bibliography
    Label -
    Modified - the date at which the file was last modified
    Owner -
    Records - an array of citations or collections
    Size - the number of records in the
    Source -`,
	Example: `  doc init
  doc init --path=/path/to/project/folder`,
	Run: func(cmd *cobra.Command, args []string) {
		initGitRepo("/tmp/test")
		//metadata, metadataErr := promptForMetadata()
		//if metadataErr != nil {
		//	log.Fatal(metadataErr)
		//	panic(metadataErr)
		//}
		//
		//dir, metadataErr := filepath.Abs(filepath.Dir(os.Args[0]))
		//if metadataErr != nil {
		//	log.Fatal(metadataErr)
		//}
		//
		//ensureErr := ensureProjectDirectory(dir)
		//if ensureErr != nil {
		//	log.Fatal(ensureErr)
		//	panic("Unable to create project directory")
		//}
		//
		//filePath := path.Join(dir, "bib.json")
		//writeErr := writeBibliography(filePath, metadata)
		//if writeErr != nil {
		//	fmt.Println("failed to write project bibliography")
		//	fmt.Println(writeErr)
		//}
		//fmt.Println("\n  Wrote", filePath, " \n ")
	},
}

// Ensure that the project directory exists.
func ensureProjectDirectory (dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func initGitRepo (dirPath string) (error) {
	out, err := exec.Command("git", "log").Output()
	if err != nil {
		return err
	}
	fmt.Printf("The date is %s\n", out)
	return nil
}

// Determine if the specified directory is a Git repository
func isGitRepo (dirPath string) (bool, error) {
	files, err := ioutil.ReadDir(dirPath)
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

// Initialize the module.
func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&ProjectPath, "path", "p", ".", "Path to project folder")
}

// Prompt user to provide project metadata
func promptForMetadata () (data.CollectionMetadata, error) {
	fmt.Println(`
This utility will create a new bib.json file in the current directory and
then ask you to provide values for a minimum number of fields. We'll try to
use some sensible defaults for the remainder.

See "doc init --help" for more details on the contents of the bib.json.

Use "doc get" afterwards to retrieve all the documents identified in the
bibliography.

Press ^C at any time to quit.
`)

	reader := bufio.NewReader(os.Stdin)
	meta := data.CollectionMetadata{}

	usr, userErr := user.Current()
	if userErr != nil {
		panic(userErr)
	}

	host, hostErr := os.Hostname()
	if hostErr != nil {
		panic(hostErr)
	}

	uid, uuidErr := uuid.NewV4()
	if uuidErr != nil {
		panic(uuidErr)
	}

	// default values
	meta.Collection = "Bibliography"
	meta.Description = "Project bibliography."
	meta.Id = uid.String()
	meta.Owner = usr.Username + "@" + host
	meta.Created = time.Now().UTC().Format(time.RFC3339)
	meta.Modified = time.Now().UTC().Format(time.RFC3339)

	// override defaults with user specified values
	fmt.Print("  Project or collection name: (Bibliography) ")
	name, _ := reader.ReadString('\n')

	fmt.Print("  Description: (Project bibliography.) ")
	desc, _ := reader.ReadString('\n')

	fmt.Print("  Owner: (" + usr.Username + "@" + host + ") ")
	owner, _ := reader.ReadString('\n')

	meta.Collection = name
	meta.Description = desc
	meta.Owner = owner

	return meta, nil
}

// Save user metadata into a local configuration file
func promptForUserData () {}

// Write bibliography file.
func writeBibliography (path string, metadata data.CollectionMetadata) error {
	data, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		panic("Couldn't convert object to JSON")
	}
	return ioutil.WriteFile("bib.json", data, os.FileMode(0666))
}

func writeGitIgnore () {}

// Write user configuration file.
func writeUserConfiguration (path string, config string) error {
	return nil
}