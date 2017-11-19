//-----------------------------------------------------------------------------
// Add command module
// Add publication by reference to the bib.json file contained in the root
// folder of the current project.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
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
}

func resolve () {}
