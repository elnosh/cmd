package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var mkdirCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "Make a new directory",

	Run: func(cmd *cobra.Command, args []string) {

		for _, arg := range args {
			if _, err := os.Stat(arg); !os.IsNotExist(err) {
				fmt.Println("File or directory already exists.")
			} else {
				err := os.Mkdir(arg, 0755)
				if err != nil {
					fmt.Println("There was an error creating the directory.")
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(mkdirCmd)
}
