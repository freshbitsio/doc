//-----------------------------------------------------------------------------
// Init command module
// A project is a file system directory that contains a BibJSON file named
// bib.json in the top level directory. This module is used to initialize a
// project by creating or resetting that top level file.
//-----------------------------------------------------------------------------

package cmd

import (
	"bufio"
	"doc/data"
	"doc/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/mitchellh/go-homedir"
	"github.com/nu7hatch/gouuid"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
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

  Press ^C at any time to quit.
`)

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		cfg, usr, err := promptForUserInput()
		if err != nil {
			fmt.Print("Failed to capture input")
			os.Exit(1)
		}

		fmt.Println("")

		err = utils.InitGitRepo(dir)
		if err != nil {
			fmt.Println("  \u2717 Failed to initialize Git repository")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Initialized Git repository")
		}

		err = writeGitignore(dir)
		if err != nil {
			fmt.Println("  \u2717 Failed to write .gitignore")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Wrote .gitignore")
		}

		err = initBibJson(dir, cfg)
		if err != nil {
			fmt.Println("  \u2717 Failed to write bib.json")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Wrote bib.json")
		}

		err = writeReadme(dir)
		if err != nil {
			fmt.Println("  \u2717 Failed to write README.md")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Wrote README.md")
		}

		err = writeLicense(dir)
		if err != nil {
			fmt.Println("  \u2717 Failed to write LICENSE")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Wrote LICENSE")
		}

		// TODO ensure that user preferences folder exists
		err = writeUserConfiguration(usr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("  \u2717 Failed to write user preferences file")
			os.Exit(1)
		} else {
			fmt.Println("  \u2713 Wrote preferences to ~/.docrc/preferences.json")
		}

		// TODO do first commit?

		fmt.Println("\n  \u270C Success! \n")
	},
}

// Initialize bib.json file in the specified directory.
func initBibJson (dir string, metadata data.CollectionMetadata) error {
	err := utils.EnsureDirectory(dir)
	if err != nil {
		return errors.New("Unable to create project directory")
		os.Exit(100)
	}

	filePath := path.Join(dir, "bib.json")
	err = writeBibJSON(filePath, metadata)
	if err != nil {
		fmt.Println("failed to write project bibliography")
		os.Exit(100)
	}

	return nil
}

// Initialize the module.
func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&ProjectPath, "path", "p", ".", "Path to project folder")
}

// Prompt user for project metadata and user profile information
func promptForUserInput() (data.CollectionMetadata, data.UserPreferences, error) {
	// TODO try to get the user name and email from the git config first before asking for it
	reader := bufio.NewReader(os.Stdin)
	meta := data.CollectionMetadata{}
	u := data.UserPreferences{}

	username, err := user.Current()
	if err != nil {
		fmt.Print("Could not determine current user")
		os.Exit(1)
	}

	host, hostErr := os.Hostname()
	if hostErr != nil {
		fmt.Print("Could not determine hostname")
		os.Exit(1)
	}

	uid, uuidErr := uuid.NewV4()
	if uuidErr != nil {
		fmt.Print("Could not generate project identifier")
		os.Exit(1)
	}

	// default values
	meta.Collection = "Bibliography"
	meta.Description = "Project bibliography."
	meta.Id = uid.String()
	meta.Owner = username.Username + "@" + host
	meta.Created = time.Now().UTC().Format(time.RFC3339)
	meta.Modified = time.Now().UTC().Format(time.RFC3339)
	meta.Records = []data.Record{}
	meta.Resources = make(map[string]string)

	// override defaults with user specified values
	fmt.Print(ansi.Color("  Project or collection name: (" + meta.Collection + ") ", "blue"))
	name, _ := reader.ReadString('\n')
	if name == "\n" {
		name = "Bibliography"
	}

	fmt.Print(ansi.Color("  Description: (" + meta.Description + ") ", "blue"))
	desc, _ := reader.ReadString('\n')
	if desc == "\n" {
		desc = "Project bibliography."
	}

	fmt.Print(ansi.Color("  Owner: (" + username.Username +  ") ", "blue"))
	owner, _ := reader.ReadString('\n')
	if owner == "\n" {
		owner = username.Username
	}

	meta.Collection = name
	meta.Description = desc
	meta.Owner = owner

	fmt.Print(ansi.Color("  Email: (" + username.Username + "@" + host + ") ", "blue"))
	email, _ := reader.ReadString('\n')
	if email == "\n" {
		email = username.Username + "@" + host
	}

	u.Fullname = strings.Trim(owner, "\n")
	u.Email = strings.Trim(email, "\n")

	return meta, u, nil
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
	bib, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		fmt.Println("Couldn't convert object to JSON")
		os.Exit(1)
	}
	return ioutil.WriteFile("bib.json", bib, os.FileMode(0666))
}

// Write default .gitignore file
func writeGitignore(dir string) error {
	f := path.Join(dir, ".gitignore")
	return ioutil.WriteFile(f, []byte(data.DefaultGitignore), os.FileMode(0666))
}

// Write default LICENSE file
func writeLicense (dir string) error {
	f := path.Join(dir, "LICENSE")
	return ioutil.WriteFile(f, []byte(data.DefaultLicense), os.FileMode(0666))
}

// Write default readme.md file
func writeReadme (dir string) error {
	f := path.Join(dir, "README.md")
	return ioutil.WriteFile(f, []byte(data.DefaultReadme), os.FileMode(0666))
}

// Write user configuration file.
func writeUserConfiguration (config data.UserPreferences) error {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Print("Could not determine user home")
		os.Exit(1)
	}
	// ensure the user config folder exists
	f := path.Join(home, ".docrc")
	err = utils.EnsureDirectory(f)
	if err != nil {
		fmt.Print("Could not create user configuration directory")
		os.Exit(1)
	}
	// write configuration
	f = path.Join(f, "preferences.json")
	cfg, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Couldn't convert user preferences to JSON")
		os.Exit(1)
	}
	return ioutil.WriteFile(f, cfg, os.FileMode(0666))
}