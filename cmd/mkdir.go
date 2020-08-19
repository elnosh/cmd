package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mkdirCmd represents the mkdir command
var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Make a new directory",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mkdir called")
	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)
}
