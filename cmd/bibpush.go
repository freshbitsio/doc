//-----------------------------------------------------------------------------
// bib/push command module
// Push changes from the local repository to the remote.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os/exec"
)

// bibpushCmd represents the bibpush command
var bibpushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote project repository",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bibpush called")
	},
}

// Initialize the module.
func init() {
	bibCmd.AddCommand(bibpushCmd)
}

// Pull changes from the remote repository.
func pushChanges () {
	// TODO should we first check to see if the current directory has a bib.json?
	out, err := exec.Command("git", "push", "origin", "master").Output()
	if err != nil {
	}
	fmt.Printf("\n%s\n", out)
}
