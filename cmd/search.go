//-----------------------------------------------------------------------------
// search command module
// Search for resources via the API.
//-----------------------------------------------------------------------------
package cmd

import (
	"bytes"
	"doc/api"
	"errors"
	"fmt"
	term "github.com/buger/goterm"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"log"
	"strings"
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
	Long:  `Search for publications by title, keyword, author, doi, and source.`,
	Example: `  doc search deep learning
  doc search --author=hinton neural networks
  doc search --author="geoffrey hinton" neural networks
  doc search --doi=10.1038/nature14539 // doesn't make sense
  doc search --source=arxiv neural networks
  doc search --source="conference machine learning" neural networks`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires one or more search terms")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// build the search query
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
		// execute the request
		res, err := search(query)
		if err != nil {
			log.Fatal(err)
		}
		// extract the raw string values that we want to present
		docs := getFields(res)
		// clean up the strings
		docs = cleanStrings(docs)
		// determine the maximum column widths
		maxwidths := getMaximumColumnWidths(docs)
		// set fixed maximum widths
		maxwidths[2] = 24
		maxwidths[3] = 6
		maxwidths[4] = 24
		// print the documents to the terminal
		print(docs,0,0, maxwidths)
	},
}

// Remove whitespace and punctuation from start and end of strings.
func cleanStrings(data [][]string) [][]string {
	for i := 0; i < len(data); i++ {
		line := data[i]
		for j := 0; j < len(line); j++ {
			line[j] = strings.Trim(line[j], "\"")
		}
	}
	return data
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
		urn = longest(line[0], urn)
		title = longest(line[1], title)
		author = longest(line[2], author)
		year = longest(line[3], year)
		publication = longest(line[4], publication)
	}

	// set fixed maximum widths

	return []int{len(urn), len(title), len(author), len(year), len(publication)}
}

// Extract fields from json records.
func getFields(data []byte) [][]string {
	jsonParsed, _ := gabs.ParseJSON([]byte(data))
	docs, _ := jsonParsed.S("docs").Children()
	var lines [][]string

	for _, doc := range docs {
		var line []string
		urn := doc.S("url").String()
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

// Return the longest string.
func longest(s1 string, s2 string) string {
	if len(s1) > len (s2) {
		return s1
	} else {
		return s2
	}
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

func padRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

// Pretty print search resultsObject to standard output
func print(docs [][]string, count int, limit int, widths []int) {
	// terminal width
	termWidth := term.Width()
	if termWidth == 0 {
		termWidth = 200
	}

	// determine the column widths
	// give any remaining space to the title column
	widths[1] = termWidth - (widths[0] + widths[2] + widths[3] + widths[4] + 8)

	fmt.Printf("\nShowing %v of %v documents\n", limit, count)

	printColumns("URN", "TITLE", "AUTHOR", "DATE", "PUBLICATION", widths)

	for i := 0; i< len(docs); i++ {
		printColumns(docs[i][0], docs[i][1], docs[i][2], docs[i][3], docs[i][4], widths)
	}
}

func printColumns(urn string, title string, author string,
	year string, publication string, widths []int) {
	var buf bytes.Buffer
	buf.WriteString(padOrTruncateString(urn, widths[0]))
	buf.WriteString("  ")
	buf.WriteString(padOrTruncateString(title, widths[1]))
	buf.WriteString("  ")
	buf.WriteString(padOrTruncateString(author, widths[2]))
	buf.WriteString("  ")
	buf.WriteString(padOrTruncateString(year, widths[3]))
	buf.WriteString("  ")
	buf.WriteString(padOrTruncateString(publication, widths[4]))
	fmt.Println(buf.String())
}

// Execute the search request against the API
func search(args map[string]string) ([]byte, error) {
	return api.GetDocs(args)
}

// Padding or truncate the string to ensure that it is the specified length
func padOrTruncateString(str string, l int) string {
	if l < 0 {
		return str
	} else if len(str) >= l {
		return str[0:l-3] + "..."
	} else {
		return padRight(str, " ", l)
	}
}
