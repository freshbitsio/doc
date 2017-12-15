//-----------------------------------------------------------------------------
// bib/log command module
// Display the history of changes to this project folder.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// biblogCmd represents the biblog command
var repoLogCmd = &cobra.Command{
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
	out, _ := exec.Command("git", "log").CombinedOutput()
	fmt.Println(string(out))
}

// Initialize the module.
func init() {
	repoCmd.AddCommand(repoLogCmd)
}
