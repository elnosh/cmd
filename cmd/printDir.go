package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// printDirectoryCmd represents the printDirectory command
var printDirectoryCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Print current directory",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("arguments not necessary")
			return
		}

		dirName, err := os.Getwd()
		if err != nil {
			fmt.Println("There was an error getting the working directory")
			return
		}
		fmt.Println(dirName)

	},
}

func init() {
	rootCmd.AddCommand(printDirectoryCmd)
}
