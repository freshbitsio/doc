//-----------------------------------------------------------------------------
// Add command module
// This module adds publications by reference to the bib.json file contained in
// the root folder of the current project.
//
// See the LICENSE file for license information.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add publications to the current project",
	Long: `Add publication references to the current project metadata file so that they can be retrieved later.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

// Find the root project folder. The root folder is the folder that contains the
// bib.json file.
func findProjectRoot () {

}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func resolve () {}
