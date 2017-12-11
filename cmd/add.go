//-----------------------------------------------------------------------------
// Add command module
// Add publication by reference to the bib.json file contained in the root
// folder of the current project.
//-----------------------------------------------------------------------------
package cmd

import (
	"doc/api"
	"doc/bib"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"time"
	"os"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add publications to the current project",
	Long: `Add publication references to the current project metadata file so that they can be retrieved later.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one or more document identifiers")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if strings.Contains(args[0], "arxiv.org") {
			AddArxivResource(args[0])
		} else {
			var query = make(map[string]string)
			addResource(args[0], query)
		}
	},
}

// Add ARXIV resource.
func AddArxivResource (s string) {
	data, err := bib.Read()
	if err != nil {
		fmt.Println("Unable to read bib.json file")
		os.Exit(90)
	}
	jsonParsed, _ := gabs.ParseJSON([]byte(data))

	// add resource
	var path = strings.Replace(s, "http://", "", 1)
	path = strings.Replace(path, "arxiv.org", "arxiv", 1)
	var url = strings.Replace(s, "/abs", "/pdf", 1) + ".pdf"
	jsonParsed.S("resources").Set(url, path)

	// update the last modified time
	jsonParsed.Set(time.Now().Local().Format(time.RFC3339), "modified")

	// write the updated bib.json file
	err = bib.Write([]byte(jsonParsed.StringIndent("", "  ")))
	if err != nil {
		fmt.Println("Couldn't write updates to bib.json")
	} else {
		fmt.Println("Resource added")
	}
}

// Add resource to project bib.json
func addResource (urn string, args map[string]string) () {
	res, err := api.GetDoc(urn, args)
	if err != nil {
		fmt.Println("Resource not found")
		os.Exit(100)
	}
	resParsed, _ := gabs.ParseJSON(res)

	// read the bib.json file
	data, err := bib.Read()
	if err != nil {
		fmt.Println("Unable to read bib.json file")
		os.Exit(90)
	}
	jsonParsed, _ := gabs.ParseJSON([]byte(data))

	// append the record to the records field
	jsonParsed.ArrayAppend(resParsed.Data(), "records")

	// update the last modified time
	jsonParsed.Set(time.Now().Local().Format(time.RFC3339), "modified")

	// write the updated bib.json file
	err = bib.Write([]byte(jsonParsed.StringIndent("", "  ")))
	if err != nil {
		fmt.Println("Couldn't write updates to bib.json")
	} else {
		fmt.Println("Resource added")
	}
}

func init() {
	RootCmd.AddCommand(addCmd)
}
