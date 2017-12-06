package api

import (
	"doc/data"
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

// Publication and data source
type Source interface {
	Get (id string) ([]byte, error)
	Info (id string) ([]byte, error)
	Search (query string, args [][]string) ([]byte, error)
}

var ClientTimeout = time.Second * 30

// Build the query portion of the request.
func buildQuery (req *http.Request, args map[string]string) (*http.Request) {
	q := req.URL.Query()
	// build query from map
	for key, _ := range args {
		if key == "q" {
			parts := strings.Split(args[key]," ")
			q.Add("q", strings.Join(parts, "+"))
		} else {
			q.Add(key, args[key])
		}
	}
	req.URL.RawQuery = q.Encode()
	return req
}

// Execute HTTP GET request.
func Get (url string, args map[string]string) (*http.Response, error) {
	// create a client
	client, req, _ := GetHttpClient(url)
	// build the query
	if len(args) > 0 {
		req = buildQuery(req, args)
	}
	// execute the request
	//fmt.Println(req.URL.String())
	return client.Do(req)
}

// Get the response body.
func getBody (res *http.Response) ([]byte, error) {
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
	// return result
	return body, nil
}

// Get the list of supported API clients.
func GetClientVersions () (data.ClientVersions, error) {
	var args map[string]string
	var versions = data.ClientVersions{}
	// execute request
	res, err := Get(apiClientVersions, args)
	if err != nil {
		return versions, err
	}
	// get the response body
	body, err := getBody(res)
	if err != nil {
		return versions, err
	}
	// marshall the json into our struct
	unmarshalErr := json.Unmarshal(body, &versions)
	if unmarshalErr != nil {
		return versions, unmarshalErr
	}
	// return result
	return versions, nil
}

// Get extended document metadata.
func GetDoc (urn string, args map[string]string) ([]byte, error) {
	res, err := Get(apiDocs + "/" + urn, args)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, errors.New(res.Status)
	}
	return getBody(res)
}

// Search for documents matching specified criteria.
func GetDocs (args map[string]string) ([]byte, error) {
	res, err := Get(apiDocs, args)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, errors.New(res.Status)
	}
	return getBody(res)
}

// Search for sources matching specified criteria.
func GetDocsSources (args map[string]string) ([]byte, error) {
	res, err := Get(apiDocSources, args)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, errors.New(res.Status)
	}
	return getBody(res)
}

// Get HTTP client
func GetHttpClient (url string) (http.Client, *http.Request, error) {
	// create a new http client and request object
	client := http.Client{
		Timeout: ClientTimeout,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return client, req, err
	}
	// identify the client to the search api
	req.Header.Set("User-Agent", "doc-client-" + data.VersionIdentifier)

	return client, req, nil
}
