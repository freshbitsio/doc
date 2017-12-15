//-----------------------------------------------------------------------------
// bib/log command module
// Display the history of changes to this project folder.
//-----------------------------------------------------------------------------
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// biblogCmd represents the biblog command
var repoCloneCmd = &cobra.Command{
	Use:   "clone [url]",
	Short: "Clone an existing repository",
	Long: `Clone an existing Git project repository to the local folder.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires URL of remote repository to be cloned")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		clone(args[0])
	},
}

// Clone the remote repository to the local folder.
func clone (url string) {
	out, _ := exec.Command("git", "clone", url).CombinedOutput()
	fmt.Println(string(out))
}

// Initialize the module.
func init() {
	repoCmd.AddCommand(repoCloneCmd)
}
