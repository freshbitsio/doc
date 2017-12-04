//-----------------------------------------------------------------------------
// bib/remote command module
// Manage settings for the Git remote.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os/exec"
)

// bibremoteCmd represents the bibremote command
var bibremoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage settings for the remote repository",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		printRemote()
	},
}

// Add remote configuration.
func add () {}

// Initialize the module.
func init() {
	bibCmd.AddCommand(bibremoteCmd)
}

// Print remote configuration.
func printRemote () {
	// TODO should we first check to see if the current directory has a bib.json?
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
	}
	fmt.Printf("\n%s\n", out)
}

// Remove remote configuration.
func remove () {}

// Add or update remote configuration.
func update () {}