package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
	"github.com/vandmo/hju/git"
)

var force bool
var recurse bool

var cleanCmd = &cobra.Command{
	Use:   "clean [-d|--recurse] [-f|--force]",
	Short: "cleans all managed repositories",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
			gitErr := doClean(folder, force, recurse)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
  cleanCmd.Flags().BoolVarP(&force, "force", "f", false, "force")
	cleanCmd.Flags().BoolVarP(&recurse, "recurse", "d", false, "recurse")
	rootCmd.AddCommand(cleanCmd)
}

func doClean(folder string, force bool, recurse bool) error {
	fmt.Printf("--- \033[32mCleaning %s\033[0m\n", folder)
	return git.Clean(folder, force, recurse)
}
