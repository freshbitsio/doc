//-----------------------------------------------------------------------------
// Bib command module
// This module manages changes to the bibliography file.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"github.com/spf13/cobra"
)

// bibCmd represents the bib command
var bibCmd = &cobra.Command{
	Use:   "bib",
	Short: "A brief description of your command",
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
	RootCmd.AddCommand(bibCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bibCmd.PersistentFlags().String("import", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bibCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
