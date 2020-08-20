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

		if len(args) == 0 {
			fmt.Println("  specify file or directory name")
		}

		for _, arg := range args {
			if _, err := os.Stat(arg); !os.IsNotExist(err) {
				fmt.Println("file or directory already exists.")
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
