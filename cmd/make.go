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

		for _, arg := range args {
			if _, err := os.Stat(arg); os.IsExist(err) {
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

