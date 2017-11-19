//-----------------------------------------------------------------------------
// search command module
// Search for resources via the API.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"bytes"
	"doc/api"
	"doc/data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-restit/lzjson"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"log"
)

// command flags
var abstract bool
var author string
var doi string
var extended bool
var format string
var source string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [options] [search terms]",
	Short: "Search for publications",
	Long: `Search for publications by title, keyword, author, doi, and source.`,
  	Example: `  doc search deep learning
  doc search --author=hinton neural networks
  doc search --author="geoffrey hinton" neural networks
  doc search --doi=10.1038/nature14539
  doc search --source=arxiv neural networks
  doc search --source="conference machine learning" neural networks`,
	Args: func(cmd *cobra.Command, args []string) error {
	  //if len(args) < 1 {
		//  return errors.New("requires at least one arg")
	  //}
	  return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := search(args)
		if err != nil {
			log.Fatal(err)
		}
		prettyColumnPrint(r)
	},
}

// Initialize the module
func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().BoolVarP(&abstract, "abstract", "b", false, "Show abstract")
	searchCmd.PersistentFlags().StringVarP(&author,"author", "a", "","Author")
	searchCmd.PersistentFlags().StringVarP(&doi, "doi", "d", "","Digital object identifier")
	searchCmd.PersistentFlags().BoolVarP(&extended, "extended", "e", false, "Show extended results list")
	searchCmd.PersistentFlags().StringVarP(&format, "format", "f", "apa", "Display format")
	searchCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Publication source or name")
}

// Pretty print search resultsObject to standard output
func prettyColumnPrint(r data.SearchResults) {
	output := []string{
		"ID | TITLE | AUTHOR | DATE | PUBLICATION",
	}
	for i := 0; i < len(r.Docs); i++ {
		d := r.Docs[i]
		var line bytes.Buffer
		line.WriteString(data.GetUri(d))
		line.WriteString(" | ")
		line.WriteString( d.Title)
		line.WriteString(" | ")
		line.WriteString(data.GetAuthorsAsString(d))
		line.WriteString(" | ")
		line.WriteString(d.Year)
		line.WriteString(" | ")
		line.WriteString("Publication Name")
		//for j := 0; j < len(d.Keywords); j++ {
		//	line.WriteString(d.Keywords[j])
		//	if (j + 1) < len(d.Keywords) {
		//		line.WriteString(" ")
		//	}
		//}
		//output = append(output, line.String())
	}
	c := columnize.SimpleFormat(output)
	fmt.Println(c)
}

func prettyTabPrint (r data.ResultsObject) {
	//w := new(tabwriter.Writer)
	//w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	//fmt.Println("\nShowing", len(r.Docs), "of", r.Count, "matches\n")
	//fmt.Fprintln(w, "URN\tTITLE\tAUTHOR\tDATE\tKEYWORDS")
	//for i := 0; i < len(r.Docs); i++ {
	//	d := r.Docs[i]
	//	fmt.Fprintln(w, d.Uri, "\t", d.Title, "\t", d.Author, "\t", d.Date, "\t", d.Keywords)
	//}
	//w.Flush()
}

func printCompactResultsListing () {}

func printFullResultsListing () {}

// Execute the search request against the API
func search (args []string) (data.SearchResults, error) {
	// search results template
	results := data.SearchResults{}

	client, req, err := api.GetHttpClient(data.ServiceApiEndpoint + "/docs")
	if err != nil {
		return results, err
	}

	// execute the request
	res, getErr := client.Do(req)
	if getErr != nil {
		return results, getErr
	}

	// TODO check the response code
	// get the response body
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return results, readErr
	}

	// ensure that the body contains valid json
	isValid := json.Valid(body)
	if isValid == false {
		return results, errors.New("Invalid server response")
	}

	// marshall the json into our struct
	unmarshalErr := json.Unmarshal(body, &results)
	if unmarshalErr != nil {
		return results, unmarshalErr
	}

	fmt.Println("Found", len(results.Docs), "documents")

	// return search results
	return results, nil
}

// Execute search against remote API.
func searchJson (args []string) (data.SearchResults, error) {
	resp, _ := http.Get("http://foobarapi.com/things")
	json := lzjson.Decode(resp.Body)
	fmt.Println(json)

	results := data.SearchResults{}
	return results, nil
}

// Truncate string to maximum length
func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
