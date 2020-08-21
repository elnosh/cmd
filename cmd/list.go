package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files and directories inside of path provided",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			readDir(".")
		}

		for _, arg := range args {
			readDir(arg)
		}

	},
}

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("path specified is not a directory")
	}
	for _, file := range files {
		fmt.Printf("%s    ", file.Name())
	}
	return
}

func init() {
	rootCmd.AddCommand(listCmd)
}
