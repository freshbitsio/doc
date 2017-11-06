// See the LICENSE file for license information.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"runtime"
)

var cfgFile string
var PlatformIdentifier = runtime.GOOS + "/" + runtime.GOARCH
var VersionIdentifier = "v0.1.0" // TODO release task to update this automatically

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "doc",
	Short: "The research publication manager",
	Long: `      _
     | |
   __| | ___   ___
  / _  |/ _ \ / __|
 | (_| | (_) | (__
  \__,_|\___/ \___|

  doc is the research publication manager.

  Discover, download, and share collections of research publications
  easily. Make your research reproducible with trivial effort.`,
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
