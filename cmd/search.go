//-----------------------------------------------------------------------------
// search command module
// Search for resources via the API.
//-----------------------------------------------------------------------------
package cmd

import (
	"bytes"
	"doc/api"
	arxiv "doc/api/arxiv"
	"doc/utils"
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"strings"
	"os"
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
	Use:   "search [source] [search terms]",
	Short: "Search for publications",
	Long:  `Search for publications by title, keyword, author, doi, and source.`,
	Example: `  doc search dblp deep learning
  doc search dblp --author=hinton neural networks
  doc search arxiv --author="geoffrey hinton" neural networks
  doc search sem neural networks
  doc search arvix neural networks --type=conference`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Requires one or more search terms")
		} else {
			options := []string{"arxiv","dblp","plos"}
			if !utils.ContainsStr(options, args[0]) {
				return errors.New("Unsupported source. Please specify one of arxiv, dblp, or plos.")
			}
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// build the search query
		var query = make(map[string]string)
		if author != "" {
			query["author"] = author
		}
		if extended != "" {
			query["extended"] = extended
		}
		if format != "" {
			query["format"] = format
		}
		// execute the request
		switch args[0] {
		case "arxiv":
			arxiv.Search("network", query)
		case "dblp":
			fmt.Println("dblp")
		case "plos":
			fmt.Println("plos")
		default:
			fmt.Println("default (arxiv)")
			// default search service
			res, err := api.GetDocs(query)
			if err != nil {
				fmt.Println("Unable to retrieve search results")
				os.Exit(100)
			}

			// extract the raw string values that we want to present
			docs := getFields(res)

			// clean up the strings
			docs = utils.CleanStrings(docs)

			fmt.Println(docs)
		}
	},
}

// Extract fields from json records.
func getFields(data []byte) [][]string {
	jsonParsed, _ := gabs.ParseJSON([]byte(data))
	docs, _ := jsonParsed.S("docs").Children()
	var lines [][]string

	for _, doc := range docs {
		var line []string
		urn := doc.S("id").String()
		title := doc.S("title").String()
		var author bytes.Buffer
		authors, _ := doc.S("authors").Children()
		for _, a := range authors {
			ath := strings.Trim(a.S("name").String(), "\"")
			author.WriteString(ath + " ")
		}
		date := doc.S("year").String()
		publication := doc.S("journal", "name").String()

		line = append(line, urn)
		line = append(line, title)
		line = append(line, author.String())
		line = append(line, date)
		line = append(line, publication)

		lines = append(lines, line)
	}
	return lines
}

// Get the maximum column width for each field.
func getMaximumColumnWidths(data [][]string) ([]int) {
	urn := ""
	title := ""
	author := ""
	year := ""
	publication := ""

	for i := 0; i < len(data); i++ {
		line := data[i]
		urn = utils.Longest(line[0], urn)
		title = utils.Longest(line[1], title)
		author = utils.Longest(line[2], author)
		year = utils.Longest(line[3], year)
		publication = utils.Longest(line[4], publication)
	}

	// set fixed maximum widths

	return []int{len(urn), len(title), len(author), len(year), len(publication)}
}

// Initialize the module
func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().BoolVarP(&abstract, "abstract", "b", false, "Show abstract")
	searchCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "Author")
	searchCmd.PersistentFlags().StringVarP(&doi, "doi", "d", "", "Digital object identifier")
	//searchCmd.PersistentFlags().BoolVarP(&extended, "extended", "e", false, "Show extended results list")
	searchCmd.PersistentFlags().StringVarP(&format, "format", "f", "bibjson", "Display format")
	searchCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Publication source or name")
}
