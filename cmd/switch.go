package cmd

import (
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
			gitErr := git.Switch(folder, branch, create)
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
