//-----------------------------------------------------------------------------
// Init command module
// A project is a file system directory that contains a BibJSON file named
// bib.json in the top level directory. This module is used to initialize a
// project by creating or resetting that top level file.
//
// See the LICENSE file for license information.
//-----------------------------------------------------------------------------
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
	"path"
)

var ProjectPath string

//"metadata": {
//	"collection": "my_collection",
//	"label": "My collection of records",
//	"description" "a great collection",
//	"id": "long_complex_uuid",
//	"owner": "test",
//	"created": "2011-10-31T16:05:23.055882",
//	"modified": "2011-10-31T16:05:23.055882",
//	"source": "http://webaddress.com/collection.bib",
//	"records": 1594,
//	"from": 0,
//	"size": 2,
//}
type metadata struct {
	Collection string `json:"collection"`
	Created string `json:"created"`
	Description string `json:"description"`
	From uint `json:"from"`
	Id string `json:"id"`
	Label string `json:"label"`
	Modified string `json:"modified"`
	Owner string `json:"owner"`
	Records uint `json:"records"`
	Size uint `json:"size"`
	Source string `json:"source"`
}

type record struct {

}

type bibliography struct {
	Metadata metadata `json:"metadata"`
	Records []record `json:"records"`
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
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		metadata, err := promptForMetadata()
		if err != nil {
			fmt.Println(err)
		}

		filePath := path.Join(dir, "bib.json")
		fmt.Println(filePath)

		writeErr := writeBibliography(filePath, metadata)
		if writeErr != nil {
			fmt.Println("failed to write project bibliography")
			fmt.Println(writeErr)
		}
	},
}

// Ensure that the project directory exists.
func ensureProjectDirectory (path string) error {
	return nil
}

// Initialize the module.
func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&ProjectPath, "path", "p", ".", "Path to project folder")
}

// Prompt user to provide project metadata
func promptForMetadata () (metadata, error) {
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
	meta := metadata{}

	usr, userErr := user.Current()
	host, hostErr := os.Hostname()
	if userErr != nil || hostErr != nil {
		panic(userErr)
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
	fmt.Print("Project or collection name: (Bibliography) ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Description: (Project bibliography.) ")
	desc, _ := reader.ReadString('\n')

	fmt.Print("Owner: (" + usr.Username + "@" + host + ") ")
	owner, _ := reader.ReadString('\n')

	meta.Collection = name
	meta.Description = desc
	meta.Owner = owner

	return meta, nil
}

// Write initial project bibliography file.
func writeBibliography (path string, metadata metadata) error {
	data, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		panic("Couldn't convert object to JSON")
	}
	ioutil.WriteFile("bib.json", data, os.FileMode(0666))
	return nil
}