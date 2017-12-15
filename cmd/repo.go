//-----------------------------------------------------------------------------
// Repo command module
// This module manages changes to the project repository.
//-----------------------------------------------------------------------------
package cmd

import (
	"github.com/spf13/cobra"
)

// bibCmd represents the bib command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage the project repository",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("bib called")
	//},
}

func init() {
	RootCmd.AddCommand(repoCmd)
}
