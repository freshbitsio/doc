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
	"doc/api"
	"doc/data"
	"fmt"
	"github.com/go-restit/lzjson"
	//"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"log"
	"github.com/ryanuber/columnize"
	"bytes"
)

// command flags
var abstract bool
var author string
var doi string
var extended string
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
		var query = make(map[string]string)
		if author != "" {
			query["author"] = author
		}
		if doi != "" {
			query["doi"] = doi
		}
		if extended != "" {
			query["extended"] = extended
		}
		if format != "" {
			query["format"] = format
		}
		res, err := search(query)
		if err != nil {
			log.Fatal(err)
		}
		prettyColumnPrint(res)
	},
}

// Initialize the module
func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().BoolVarP(&abstract, "abstract", "b", false, "Show abstract")
	searchCmd.PersistentFlags().StringVarP(&author,"author", "a", "","Author")
	searchCmd.PersistentFlags().StringVarP(&doi, "doi", "d", "","Digital object identifier")
	//searchCmd.PersistentFlags().BoolVarP(&extended, "extended", "e", false, "Show extended results list")
	searchCmd.PersistentFlags().StringVarP(&format, "format", "f", "bibjson", "Display format")
	searchCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Publication source or name")
}

// Pretty print search resultsObject to standard output
func prettyColumnPrint(res lzjson.Node) {
	//count := res.Get("count").String()
	//limit := res.Get("limit").String()
	//fmt.Println("\nShowing %s of %s documents", limit, count)
	fmt.Println("\nShowing 1 of 1 documents")
	docs := res.Get("docs")
	output := []string{
		"ID | TITLE | AUTHOR | DATE | PUBLICATION",
	}
	for i:=0; i < docs.Len(); i++ {
		d := docs.GetN(i)
		var line bytes.Buffer
		line.WriteString(d.Get("url").String())
		line.WriteString(" | ")
		line.WriteString( d.Get("title").String())
		line.WriteString(" | ")
		line.WriteString("Lead author")
		//line.WriteString(data.GetAuthorsAsString(d))
		line.WriteString(" | ")
		line.WriteString(d.Get("year").String())
		line.WriteString(" | ")
		line.WriteString("Publication Name")
		output = append(output, line.String())
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
func search (args map[string]string) (lzjson.Node, error) {
	return api.GetDocs(args)
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
