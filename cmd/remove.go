package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var directoryName string

var removeCmd = &cobra.Command{
	Use:   "rmv",
	Short: "Remove file or directory",

	Run: func(cmd *cobra.Command, args []string) {

		if directoryName != "" {
			os.RemoveAll(directoryName)
			return
		}

		arg := fmt.Sprint(strings.Join(args, " "))

		if _, err := os.Stat(arg); os.IsNotExist(err) {
			fmt.Println("file or directory does not exist")
			return
		}
		err := os.Remove(arg)
		if err != nil {
			fmt.Println("There was an error deleting the file or directory")
		}
	},
}


func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&directoryName, "all", "a", "", "remove directory and its content")
}
