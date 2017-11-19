package data

// Result document data structure
type Doc struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Uri string `json:"uri"`
	Date string `json:"date"`
	Keywords []string `json:"keywords"`
}

// JSON links data structure
type Links struct {
	Self string `json:"self"`
	Prev string `json:"prev"`
	Next string `json:"next"`
}

// Search results data structure
type ResultsObject struct {
	Links Links `json:"_links"`
	Count int `json:"count"`
	Docs []Doc `json:"docs"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type DocSearchResults struct {
	Comments string 		  `json:"comments"`
	Links  SearchResultsLinks `json:"_links"`
	Count  uint8              `json:"count"`
	Limit  uint8              `json:"limit"`
	Offset uint8              `json:"offset"`
	Docs   []Record           `json:"docs"`
}

type SearchResultsLinks struct {
	TermsOfUse string `json:"termsofuse"`
	Docs string `json:"documentation"`
	Self string `json:"self"`
	Next string `json:"next"`
	Prev string `json:"prev"`
}
