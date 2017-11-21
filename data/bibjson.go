package data

import (
	"strings"
)

type Author struct {
	Name string `json:"name"`
}

// Collection
type Collection struct {
	Metadata CollectionMetadata `json:"metadata"`
	Records  []Record           `json:"records"`
}

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
type CollectionMetadata struct {
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

// Identifier record
type Identifier struct {
	Type string `json:"type"`
	ID string `json:"id"`
}

// Link record
type Link struct {
	Download string `json:"download"`
	Href string `json:"href"`
	Url string `json:"url"`
}

//{
//	"title": "Open Bibliography for Science, Technology and Medicine",
//	"author":[{"name": "Richard Jones"},],
//	"type": "article",
//	"year": "2011",
//	"journal": {"name": "Journal of Cheminformatics"},
//	"link": [{"url":"http://www.jcheminf.com/content/3/1/47"}],
//	"identifier": [{"type":"doi","id":"10.1186/1758-2946-3-47"}]
//}
type Record struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Author []Author `json:"author"`
	Type string `json:"type"`
	Year string `json:"year"`
	Journal string `json:"journal"`
	Link []Link `json:"link"`
	Identifier []Identifier`json:"identifier"`
}

func GetAuthorsAsString (r Record) string {
	var names []string
	for i := 0; i< len(r.Author); i++ {
		author := r.Author[i]
		names = append(names, author.Name)
	}
	return strings.Join(names, ", ")
}

func GetDoi (r Record) string {
	return ""
}

func GetDownloadLink (r Record) string {
	for i := 0; i< len(r.Link); i++ {
		link := r.Link[i]
		if link.Download != "" {
			return link.Download
		}
	}
	return ""
}

func GetUri (r Record) string {
	return r.Id
}

func RemoveRecord (r Record) (bool, error) {
	return false, nil
}