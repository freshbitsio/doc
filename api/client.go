package api

import (
	"doc/data"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	//"strings"
	"time"
)

// Build the query portion of the request.
func buildQuery (req *http.Request, args map[string]string) (*http.Request) {
	// build the request query
	q := req.URL.Query()
	//if author != "" {
	//	q.Add("author", author)
	//}
	//if doi != "" {
	//	q.Add("doi", doi)
	//}
	//if source != "" {
	//	q.Add("source", source)
	//}
	//if len(args) > 0 {
	//	q.Add("q", strings.Join(args, "+"))
	//}
	q.Add("offset", "0")
	q.Add("limit", "100")
	q.Add("sort", "date,title")
	req.URL.RawQuery = q.Encode()

	return req
}

// Execute HTTP GET request.
// var m map[string]string
// make(map[string]Vertex)
//m["Bell Labs"] = Vertex{
//40.68433, -74.39967,
//}
//fmt.Println(m["Bell Labs"])
func Get (url string, args map[string]string) ([]byte, error) {
	// create a client
	var body []byte
	client, req, err := GetHttpClient(url)
	if err != nil {
		return body, err
	}
	// build the query
	if len(args) > 0 {
		req = buildQuery(req, args)
	}
	// execute the request
	res, getErr := client.Do(req)
	if getErr != nil {
		return body, getErr
	}
	// get the response body
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return body, readErr
	}
	// ensure that the body contains valid json
	isValid := json.Valid(body)
	if isValid == false {
		return body, errors.New("Invalid server response")
	}
	return body, nil
}

// Get client versions
func GetClientVersions () (data.ClientVersions, error) {
	var args map[string]string
	var versions = data.ClientVersions{}
	// get client versions
	body, err := Get(apiClientVersions, args)
	if err != nil {
		return versions, err
	}
	// marshall the json into our struct
	unmarshalErr := json.Unmarshal(body, &versions)
	if unmarshalErr != nil {
		return versions, unmarshalErr
	}
	return versions, nil
}

func GetDocs (args map[string]string) {

}

func GetDocsSources (args map[string]string) {

}

// Get HTTP client
func GetHttpClient (url string) (http.Client, *http.Request, error) {
	// create a new http client and request object
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return client, req, err
	}
	// identify the client to the search api
	req.Header.Set("User-Agent", "doc-client-" + data.VersionIdentifier)

	return client, req, nil
}
