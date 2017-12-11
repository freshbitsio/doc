//-----------------------------------------------------------------------------
// get command module
// Retrieve remote resources to the local storage.
//-----------------------------------------------------------------------------
package cmd

import (
	"doc/bib"
	"doc/utils"
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var depth uint8
var downloadurl string
//var queue []string
var save bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve individual and collections of publications",
	Long: `Download individual publications, and collections of publications to the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		// download resources from project file
		resources := getProjectResources()
		for path, url := range resources {
			url = strings.Trim(url, "\"")
			download(url, path)
		}
		fmt.Printf("\nDone")
	},
}

// Download the resource to the specified path.
func download(url string, p string) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(99)
	}

	dir := path.Join(cwd, "resources", filepath.Dir(p))
	f := path.Join(cwd, "resources", p + ".pdf")

	err = utils.EnsureDirectory(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(99)
	}

	// remove existing file if present
	err = os.RemoveAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(99)
	}

	// create client
	client := grab.NewClient()
	req, err := grab.NewRequest(f, url)
	if err != nil {
		fmt.Println(err)
		os.Exit(99)
	}

	// start download
	fmt.Println("000 Downloading", req.URL())
	resp := client.Do(req)

	// display retrieval progress
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	Loop:
		for {
			select {
				case <- t.C:
					fmt.Printf("\r000 Transferred %v/%v bytes (%.2f%%)",
						resp.BytesComplete(),
						resp.Size,
						100 * resp.Progress())
				case <- resp.Done:
					break Loop // download complete
			}
		}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ERR Download failed:", err)
		os.Exit(99)
	} else {
		fmt.Println("000 Saved", resp.Filename)
	}

	return nil
}

// Get project resources
func getProjectResources () map[string]string {
	resources := make(map[string]string)

	data, err := bib.Read()
	if err != nil {
		fmt.Println("Unable to read bib.json file")
		os.Exit(90)
	}
	jsonParsed, _ := gabs.ParseJSON([]byte(data))

	children, _ := jsonParsed.S("resources").ChildrenMap()
	for key, child := range children {
		resources[key] = child.String()
	}

	return resources
}

// Initialize the module.
func init() {
	RootCmd.AddCommand(getCmd)
	//getCmd.SetUsageTemplate("This is the template")
	//getCmd.PersistentFlags().BoolVarP(&save,"save",  "s", false, "Save document to project")
	//getCmd.PersistentFlags().Uint8VarP(&depth,"depth",  "d", 0, "Retrieve cited references to specified depth. Maximum depth of three")
}

// Print the download status to the console
func printStatus () {

}

func saveFile () (error) {
	return nil
}