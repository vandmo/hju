package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var create bool

var switchCmd = &cobra.Command{
	Use:   "switch [-c|--create] <branch>",
	Short: "Switches to a branch in all managed repositories",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		branch := args[0]
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doSwitch(folder, branch, create)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	switchCmd.Flags().BoolVarP(&create, "create", "c", false, "create the branch if it doesn't exist")
	rootCmd.AddCommand(switchCmd)
}

func doSwitch(folder string, branch string, create bool) error {
	hasRef, err := git.HasRef(folder, branch)
	if err != nil {
		return err
	}

	if hasRef {
		fmt.Printf("--- \033[32mSwitching to branch %s in %s\033[0m\n", branch, folder)
		return git.Switch(folder, branch, false)
	} else if create {
		fmt.Printf("--- \033[32mCreating and switching to branch %s in %s\033[0m\n", branch, folder)
		return git.Switch(folder, branch, true)
	} else {
		fmt.Printf("--- \033[32mNOT creating NOR switching to branch %s in %s\033[0m\n", branch, folder)
		return nil
	}
}
