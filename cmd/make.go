package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var makeCmd = &cobra.Command{
	Use: "mk",
	Short: "Make a new file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("  specify file or directory name")
		}

		for _, arg := range args {
			if _, err := os.Stat(arg); err == nil {
				fmt.Println("File already exists")
			} else {
				f, err := os.Create(arg)
				if err != nil {
					fmt.Println("There was an error creating the file")
				}
				defer f.Close()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}