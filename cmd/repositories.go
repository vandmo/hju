package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
)

var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Prints the managed repositories",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, repository := range hjuFile.Repositories {
			fmt.Println(repository)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
}
