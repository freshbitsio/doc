//-----------------------------------------------------------------------------
// Manage settings for the Git repository remote.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
	"os"
)

// bibremoteCmd represents the bibremote command
var repoRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage settings for the remote repository",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			show()
		} else if len(args) == 1 {
			set(args[0])
		}
	},
}

// Add remote configuration.
func add (url string) {
	out, _ := exec.Command("git", "remote", "set-url", "origin", url).CombinedOutput()
	fmt.Println(string(out))
}

// Determine if the repo has an existing remote configuration.
func hasRemote () bool {
	out, err := exec.Command("git", "remote", "-v").CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		os.Exit(99)
	}
	if len(out) > 0 {
		lines := strings.Split(string(out),"\n")
		for _, line := range lines {
			if strings.Index(line, "origin") == 0 {
				return true
			}
		}
	}
	return false
}

// Initialize the module.
func init() {
	repoCmd.AddCommand(repoRemoteCmd)
}

// Show remote configuration.
func show() {
	out, _ := exec.Command("git", "remote", "-v").CombinedOutput()
	fmt.Printf("%s", out)
}

// Set or update remote configuration.
func set (url string) {
	if hasRemote() {
		out, _ := exec.Command("git", "remote", "set-url", "origin", url).CombinedOutput()
		fmt.Println(string(out))
	} else {
		out, _ := exec.Command("git", "remote", "add", "origin", url).CombinedOutput()
		fmt.Println(string(out))
	}
}
