//-----------------------------------------------------------------------------
// root command module
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"os"
	"github.com/mgutz/ansi"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var longDescription = `      _
     | |
   __| | ___   ___
  / _  |/ _ \ / __|
 | (_| | (_) | (__
  \__,_|\___/ \___|

  The research package manager.

  Discover, download, and share collections of research publications and
  datasets easily. Make your research reproducible with trivial effort.`

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "doc",
	Short: "The research publication manager",
	Long: ansi.Color(longDescription, "blue"),
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize the module
func init() { 
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".doc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".doc")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
