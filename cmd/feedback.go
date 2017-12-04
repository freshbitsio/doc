//-----------------------------------------------------------------------------
// feedback command module
// Send quality and crash feedback to the project developers.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// feedbackCmd represents the feedback command
var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Send comments, request features, suggest new publication sources",
	Long: `Ok, we know you're smart since you're using A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("feedback called")
	},
}

func init() {
	RootCmd.AddCommand(feedbackCmd)
}
