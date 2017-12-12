//-----------------------------------------------------------------------------
// remove command module
// Remove resource references from the bib.json.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove publication from the current project",
	Long: `Remove publication reference from the current project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
