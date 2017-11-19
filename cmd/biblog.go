//-----------------------------------------------------------------------------
// bib/log command module
// Display the history of changes to this project folder.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// biblogCmd represents the biblog command
var biblogCmd = &cobra.Command{
	Use:   "log",
	Short: "Display the project change log",
	Long: `Print the project's Git repository log to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		printLog()
	},
}

// Print the repository log
func printLog () {
	// TODO should we first check to see if the current directory has a bib.json?
	out, err := exec.Command("git", "log").Output()
	if err != nil {
	}
	fmt.Printf("\n%s\n", out)
}

// Initialize the module.
func init() {
	bibCmd.AddCommand(biblogCmd)
}
