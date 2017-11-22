//-----------------------------------------------------------------------------
// bib/save command module
// Create a Git commit that saves the current state of the repository.
//
// Copyright (c) 2017 Davis Marques <dmarques@freshbits.io> and
// Hossein Pursultani <hossein@freshbits.io> See the LICENSE file for license
// information.
//-----------------------------------------------------------------------------
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"bufio"
	"os"
	"strings"
)

// bibsaveCmd represents the bibsave command
var bibsaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a snapshot of the current bibliography state",
	Long: `Create a commit of all changes in the current directory.
This command assumes that you are using a single branch development approach
and that all changes are made on the master branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		//showStatus()
		addAllFiles()
		commit()
	},
}

func addAllFiles() {
	// TODO add only the set of target files specified in configuration
	out, err := exec.Command("git", "add", "--all").Output()
	if err != nil {
		fmt.Println("add files")
		panic(err)
	}
	if VerboseOutput == true {
		fmt.Printf("\n%s\n", out)
	}
}

func commit() {
	// prompt user for commit message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Describe the changes you are saving: ")
	msg, _ := reader.ReadString('\n')
	msg = "\"" + strings.TrimSpace(msg) + "\""
	// save commit
	out, err := exec.Command("git", "commit", "-m", msg).Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", out)
	fmt.Println("Commit saved")
}

// Initialize the module.
func init() {
	bibCmd.AddCommand(bibsaveCmd)
}

// Show status.
func showStatus () {
	out, _ := exec.Command("git", "status").Output()
	fmt.Printf("\n%s\n", out)
}