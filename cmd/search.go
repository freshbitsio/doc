// See the LICENSE file for license information.

package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	//"text/tabwriter"
	"strings"
)

// command flags
var author string
var doi string
var source string

// URL to the search API endpoint
var ServiceApiEndpoint = "http://davismarques.com/priv/results.json"

// Result document data structure
type doc struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Uri string `json:"uri"`
	Date string `json:"date"`
	Keywords []string `json:"keywords"`
}

// JSON links data structure
type links struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

// Search resultsObject data structure
type resultsObject struct {
	Links links `json:"_links"`
	Count int `json:"count"`
	Docs []doc `json:"docs"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}


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
		fmt.Println(author)
		fmt.Println(doi)
		fmt.Println(source)
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
		searchCmd.PersistentFlags().StringVarP(&author,"author", "a", "","Author")
		searchCmd.PersistentFlags().StringVarP(&doi, "doi", "d", "","Digital object identifier")
		searchCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Publication source or name")
	}

	// Pretty print search resultsObject to standard output
	func prettyColumnPrint(r resultsObject) {
		output := []string{
			"URN | TITLE | AUTHOR | DATE | KEYWORDS",
		}
		for i := 0; i < len(r.Docs); i++ {
			d := r.Docs[i]
			var line bytes.Buffer
			line.WriteString(d.Uri)
			line.WriteString(" | ")
			line.WriteString( d.Title)
			line.WriteString(" | ")
			line.WriteString(d.Author)
			line.WriteString(" | ")
			line.WriteString(d.Date)
			line.WriteString(" | ")
			for j := 0; j < len(d.Keywords); j++ {
				line.WriteString(d.Keywords[j])
				if (j + 1) < len(d.Keywords) {
					line.WriteString(" ")
				}
			}
			output = append(output, line.String())
		}
		c := columnize.SimpleFormat(output)
		fmt.Println(c)
	}

	func prettyTabPrint (r resultsObject) {
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

	// Execute the search request against the API
	func search (args []string) (resultsObject, error) {
		// search results template
		results := resultsObject{}

		// create a new http client and request object
		client := http.Client{
			Timeout: time.Second * 2,
		}
		req, err := http.NewRequest(http.MethodGet, ServiceApiEndpoint, nil)
		if err != nil {
			return results, err
		}

		// identify the client to the search api
		req.Header.Set("User-Agent", "doc-client-" + VersionIdentifier)

		// build the search query
		q := req.URL.Query()
		if author != "" {
			q.Add("author", author)
		}
		if doi != "" {
			q.Add("doi", doi)
		}
		if source != "" {
			q.Add("source", source)
		}
		if len(args) > 0 {
			q.Add("q", strings.Join(args, "+"))
		}
		q.Add("offset", "0")
		q.Add("limit", "100")
		q.Add("sort", "date,title")
		req.URL.RawQuery = q.Encode()

		// execute the request
		// fmt.Println(req.URL.String())
		res, getErr := client.Do(req)
		if getErr != nil {
			return results, getErr
		}

		// get the response body
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return results, readErr
		}

		// ensure that the body contains valid json
		isValid := json.Valid(body)
		if isValid == false {
			return results, errors.New("Invalid JSON")
		}

		// marshall the json into our struct
		unmarshalErr := json.Unmarshal(body, &results)
		if unmarshalErr != nil {
			return results, unmarshalErr
		}

		// return search results
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
