//-----------------------------------------------------------------------------
// get command module
// Retrieve remote resources to the local storage.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/spf13/cobra"
	"os"
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
		// if a URL is not provided as an arg then get the list of resources from the bib file and download them
		getResource("")
	},
}

// Get publication metadata
func getMetadata () {}

// Download the resource.
func getResource(url string) {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", "http://www.golang-book.com/public/pdf/gobook.pdf")

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size,
					100*resp.Progress())

			case <-resp.Done:
				// download is complete
				break Loop
			}
		}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}

func getResources (urns []string) {}

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