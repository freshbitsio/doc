// See the LICENSE file for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
)

var depth uint8
var queue []string
var save bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve individual and collections of publications",
	Long: `Download individual publications, and collections of publications to the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

// Enqueue download task
func enqueue () {}

// Get publication file
func getFile (uri string) ([]byte, error) {
	// create a new http client and request object
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, ServiceApiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// identify the client to the search api
	req.Header.Set("User-Agent", "doc-client-" + VersionIdentifier)

	// build the search query
	q := req.URL.Query()
	q.Add("uri", uri)
	req.URL.RawQuery = q.Encode()

	// execute the request
	// fmt.Println(req.URL.String())
	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	// get the response body
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	// TODO return mimetype
	return body, nil
}

// Get publication metadata
func getMetadata () {

}

func init() {
	RootCmd.AddCommand(getCmd)
	//getCmd.SetUsageTemplate("This is the template")
	getCmd.PersistentFlags().BoolVarP(&save,"save",  "s", false, "Save document to project")
	getCmd.PersistentFlags().Uint8VarP(&depth,"depth",  "d", 0, "Retrieve cited references to specified depth. Maximum depth of three")
}

// Print the download status to the console
func printStatus () {

}

func saveFile () (error) {
	return nil
}