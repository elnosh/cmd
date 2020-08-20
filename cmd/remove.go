package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rmvAll *bool

var removeCmd = &cobra.Command{
	Use:   "rmv",
	Short: "Remove file or directory",

	Run: func(cmd *cobra.Command, args []string) {

		// TODO: check if args is empty

		if *rmvAll {
			for _, arg := range args {
				if _, err := os.Stat(arg); os.IsNotExist(err) {
					fmt.Println("File or directory does not exist.")
				} else {
					err := os.RemoveAll(arg)
					if err != nil {
						fmt.Println("There was an error deleting the file or directory")
					}
				}
			}
		} else {
			for _, arg := range args {
				if _, err := os.Stat(arg); os.IsNotExist(err) {
					fmt.Println("File or directory does not exist")
				} else {
					err := os.Remove(arg)
					if err != nil {
						fmt.Println("There was an error deleting the file or directory")
					}
				}
			}
		}
	},
}


func init() {
	rootCmd.AddCommand(removeCmd)
	rmvAll = removeCmd.Flags().BoolP("all", "a", false, "remove directory with its content")
}

