// See the LICENSE file for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Long: `Print the application version identifier to the console`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doc version " + VersionIdentifier + " " + PlatformIdentifier)
	},
}

// Initialize the module.
func init() {
	RootCmd.AddCommand(versionCmd)
}
