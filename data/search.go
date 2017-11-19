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

// Search resultsObject data structure
type ResultsObject struct {
	Links Links `json:"_links"`
	Count int `json:"count"`
	Docs []Doc `json:"docs"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}