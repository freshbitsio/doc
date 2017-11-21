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
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/nu7hatch/gouuid"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"time"
	"strings"
)

var ProjectPath string

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
		// TODO is it an existing doc repo?
		// TODO is it an existing git repo?
		// TODO are you sure you want to init?
		fmt.Println(`
  This utility will create a new bib.json file in the current directory and
  then ask you to provide values for a minimum number of fields. We'll try to
  use some sensible defaults for the remainder.

  See "doc init --help" for more details on the contents of the bib.json.

  Use "doc get" afterwards to retrieve all the documents identified in the
  bibliography.

  Press ^C at any time to quit.`)

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		err = initGitRepo(dir)
		if err != nil {
			fmt.Println("Failed to initialize Git repository")
			os.Exit(1)
		} else {
			fmt.Println("  Initialized empty Git repository\n")
		}

		err = writeGitIgnore(dir)
		if err != nil {
			fmt.Println("Failed to write .gitignore")
			os.Exit(1)
		}

		err = initBibJson(dir)
		if err != nil {
			fmt.Println("Failed to write bib.json")
			os.Exit(1)
		} else {
			fmt.Println("  \nWrote bib.json file\n")
		}

		cfg, err := promptForUserData()
		if err != nil {
			fmt.Println("Failed to collect user configuration")
			os.Exit(1)
		}

		err = writeUserConfiguration(dir, cfg)
		if err != nil {
			fmt.Println("Failed to write user preferences file")
			os.Exit(1)
		} else {
			fmt.Println("  Wrote preferences to ~/.docrc")
		}
		fmt.Println("\n  Done")
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

// Initialize bib.json file in the specified directory.
func initBibJson (dir string) error {
	metadata, err := promptForMetadata()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = ensureProjectDirectory(dir)
	if err != nil {
		return errors.New("Unable to create project directory")
	}

	filePath := path.Join(dir, "bib.json")
	err = writeBibJSON(filePath, metadata)
	if err != nil {
		fmt.Println("failed to write project bibliography")
		fmt.Println(err)
	}

	return nil
}

// Initialize Git repository in the specified directory.
func initGitRepo (dirPath string) (error) {
	// TODO revise so that we can specify the cwd for the command
	_, err := exec.Command("git", "init").Output()
	if err != nil {
		return err
	}
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
	fmt.Print(ansi.Color("  Project or collection name: (Bibliography) ", "blue"))
	name, _ := reader.ReadString('\n')

	fmt.Print(ansi.Color("  Description: (Project bibliography.) ", "blue"))
	desc, _ := reader.ReadString('\n')

	fmt.Print(ansi.Color("  Owner: (" + usr.Username + "@" + host + ") ", "blue"))
	owner, _ := reader.ReadString('\n')

	meta.Collection = name
	meta.Description = desc
	meta.Owner = owner

	return meta, nil
}

// Save user metadata into a local configuration file
func promptForUserData () (data.UserPreferences, error) {
	// TODO try to get the user name and email from the git config first before asking for it
	reader := bufio.NewReader(os.Stdin)
	meta := data.UserPreferences{}

	fmt.Print(ansi.Color("  Fullname: ", "blue"))
	fullname, err := reader.ReadString('\n')
	if err != nil {
		return meta, err
	}

	fmt.Print(ansi.Color("  Email: ", "blue"))
	email, err := reader.ReadString('\n')
	if err != nil {
		return meta, err
	}

	meta.Fullname = strings.Trim(fullname, "\n")
	meta.Email = strings.Trim(email, "\n")

	return meta, nil
}

// Write bib.json bibliography file.
func writeBibJSON(path string, metadata data.CollectionMetadata) error {
	data, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		panic("Couldn't convert object to JSON")
	}
	return ioutil.WriteFile("bib.json", data, os.FileMode(0666))
}

func writeGitIgnore (dir string) error {
	var gitignore = `*~
lib/**/*
tmp`
	return ioutil.WriteFile(".gitignore", []byte(gitignore), os.FileMode(0666))
}

// Write user configuration file.
func writeUserConfiguration (dir string, config data.UserPreferences) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		panic("Couldn't convert user preferences to JSON")
	}
	return ioutil.WriteFile("~/.doc/preferences.json", data, os.FileMode(0666))
}