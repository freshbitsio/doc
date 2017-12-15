//-----------------------------------------------------------------------------
// bib/pull command module
// Retrieve changes from the Git remote and merge them into the local
// repository.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// bibpullCmd represents the bibpull command
var repoPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes from the remote project repository",
	Long: `Pull and merge changes from the remote Git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		pullChanges()
	},
}

// Initialize the module.
func init() {
	repoCmd.AddCommand(repoPullCmd)
}

// Pull changes from the remote repository.
func pullChanges () {
	out, _ := exec.Command("git", "pull").CombinedOutput()
	fmt.Println(string(out))
}
