package cmd

import (
    "fmt"

	"github.com/spf13/cobra"
	"github.com/vandmo/hju/core"
)

var foldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "Prints the managed folders",
	Args:  cobra.ExactValidArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		hjuFile, parseErr := core.ParseHjuFile()
		if parseErr != nil {
			return parseErr
		}
		for _, folder := range hjuFile.Folders {
            fmt.Println(folder)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(foldersCmd)
}
