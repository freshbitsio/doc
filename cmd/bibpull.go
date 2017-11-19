//-----------------------------------------------------------------------------
// bib/pull command module
// Retrieve changes from the Git remote and merge them into the local
// repository.
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

// bibpullCmd represents the bibpull command
var bibpullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes from the remote project repository",
	Long: `Pull and merge changes from the remote Git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		pullChanges()
	},
}

// Initialize the module.
func init() {
	bibCmd.AddCommand(bibpullCmd)
}

// Pull changes from the remote repository.
func pullChanges () {
	// TODO should we first check to see if the current directory has a bib.json?
	out, err := exec.Command("git", "log").Output()
	if err != nil {
	}
	fmt.Printf("\n%s\n", out)
}
