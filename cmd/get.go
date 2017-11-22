//-----------------------------------------------------------------------------
// get command module
// Retrieve remote resources to the local storage.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"github.com/spf13/cobra"
	"time"
	"os"
)

var depth uint8
//var queue []string
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

// Get publication metadata
func getMetadata () {}

func getResource(urn string) {
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
	getCmd.PersistentFlags().BoolVarP(&save,"save",  "s", false, "Save document to project")
	getCmd.PersistentFlags().Uint8VarP(&depth,"depth",  "d", 0, "Retrieve cited references to specified depth. Maximum depth of three")
}

// Print the download status to the console
func printStatus () {

}

func saveFile () (error) {
	return nil
}