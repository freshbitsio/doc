//-----------------------------------------------------------------------------
// version command module
// Display version information for this application.
//-----------------------------------------------------------------------------
package cmd

import (
	"freshbits.io/doc/api"
	"freshbits.io/doc/data"
	"fmt"
	"github.com/spf13/cobra"
)

// Version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Long: `Print the application version identifier to the console`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doc version " + data.VersionIdentifier + " " + data.PlatformIdentifier)
	},
}

// Initialize the module.
func init() {
	RootCmd.AddCommand(versionCmd)
}

// Determine if a new version of the application is available.
func updateIsAvailable(args []string) bool {
	versions, err := api.GetClientVersions()
	if err != nil {
		panic(err)
	}

	// TODO compare application version to latest
	if data.VersionIdentifier == versions.Latest {
		return false
	} else {
		return true
	}
}
