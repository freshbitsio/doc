//-----------------------------------------------------------------------------
// Add command module
// Add publication by reference to the bib.json file contained in the root
// folder of the current project.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"doc/api"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"time"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add publications to the current project",
	Long: `Add publication references to the current project metadata file so that they can be retrieved later.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			var query = make(map[string]string)
			addResource(args[0], query)
		} else {
			// error
		}
	},
}

// Add resource to project bib.json
func addResource (urn string, args map[string]string) () {
	res, err := api.GetDoc(urn, args)
	if err != nil {
		fmt.Println("Client timed out while retrieving the resource")
		panic(err)
	}
	resParsed, _ := gabs.ParseJSON(res)

	// read the bib.json file
	data, err := ioutil.ReadFile("bib.json")
	if err != nil {
		fmt.Println("Unable to read bib.json file")
		panic(err)
	}
	jsonParsed, _ := gabs.ParseJSON([]byte(data))

	// create a new record
	id := strings.Trim(resParsed.S("id").String(), "\"")
	title := strings.Trim(resParsed.S("title").String(), "\"")
	year := strings.Trim(resParsed.S("year").String(), "\"")
	record := gabs.New()
	record.Set(id, "urn")
	record.Set(title, "title")
	record.Set(year, "year")

	// append the record to the records field
	jsonParsed.ArrayAppend(record.Data(), "records")

	// update the last modified time
	jsonParsed.Set(time.Now().Local().Format(time.RFC3339), "modified")

	// write the updated bib.json file
	//fmt.Println("%s", jsonParsed.StringIndent("", "  "))
	ioutil.WriteFile("bib.json", []byte(jsonParsed.StringIndent("", "  ")), os.FileMode(0666))

	fmt.Println("Resource added")
}

func init() {
	RootCmd.AddCommand(addCmd)
}
