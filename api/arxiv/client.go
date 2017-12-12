//-----------------------------------------------------------------------------
// Interface to the arXiv publication archive
//-----------------------------------------------------------------------------

package arxiv

import (
	"doc/api"
	"encoding/xml"
	"errors"
	"github.com/fatih/color"
	"github.com/ryanuber/columnize"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
	"strings"
)

var EndPoint = "http://export.arxiv.org/api/query"

type Name struct {
	Name string `xml:"name"`
}

type Link struct {
	Href string `xml:"href,attr"`
	Rel string `xml:"rel,attr"`
	Title string `xml:"title,attr"`
	Type string `xml:"type,attr"`
}

type Entry struct {
	Id string `xml:"id"`
	Updated string `xml:"updated"`
	Published string `xml:"published"`
	PrimaryCategory string `xml:"arxiv"`
	Title string `xml:"title"`
	Summary string `xml:"summary"`
	Author []Name `xml:"author"`
	Links []Link `xml:"link"`
}

type Feed struct {
	Id string `xml:"id"`
	Title string `xml:"title"`
	Entries []Entry `xml:"entry"`
}

type query struct {
	Args [][]string
	Id string
	Query string
}

// Parse XML response
func getBody (res *http.Response) Feed {
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		os.Exit(99)
	}

	var f Feed
	err := xml.Unmarshal(body, &f)
	if err != nil {
		os.Exit(99)
	}

	return f
}

// Retrieve publication or data resource.
func Get (id string) ([]byte, error) {
	return nil, nil
}

// Get resource metadata.
func Info (id string) ([]byte, error) {
	return nil, nil
}

func Print (feed Feed) {
	data := []string{"ID | TITLE | YEAR | AUTHOR | PUBLICATION"}
	for _, f := range feed.Entries {
		// TODO factor this out into a util module
		tf := strings.Split(f.Title, "\n")
		for i, s := range tf {
			tf[i] = strings.Trim(s, " ")
		}
		title := strings.Join(tf, " ")
		if len(title) > 64 {
			title = title[0:64-3] + "..."
		}
		authors := []string{}
		for _, a := range f.Author {
			authors = append(authors, a.Name)
		}
		author := strings.Join(authors, ", ")
		if len(author) > 16 {
			author = author[0:15-3] + "..."
		}
 		cols := []string{color.BlueString(f.Id), title, f.Published[0:4], author, "journal name"}
		row :=  strings.Join(cols, " | ")
		data = append(data, row)
	}
	result := columnize.SimpleFormat(data)
	fmt.Println("\n" + result)
}

// Search for resources.
func Search (query string, args map[string]string) (error) {
	var argz = make(map[string]string)
	argz["max_results"] = "100"
	argz["search_query"] = "all:neural network"
	argz["start"] = "0"

	res, err := api.Get(EndPoint, argz)
	if err != nil {
		return err
	}
	if res.StatusCode >= 300 {
		return errors.New(res.Status)
	}

	f := getBody(res)
	Print(f)

	os.Exit(90)

	return nil
}