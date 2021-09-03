package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hju",
	Short: "Taking git to the next level",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
