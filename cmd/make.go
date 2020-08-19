package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var file string
var directory string

var makeCmd = &cobra.Command{
	Use: "mk",
	Short: "Make a new file or directory",
	Run: func(cmd *cobra.Command, args []string) {

		if file != "" {
			f, err := os.Create(file)
			if err != nil {
				fmt.Println("There was an error creating file.")
			}
			defer f.Close()
		}

		if directory != "" {
			err := os.Mkdir(directory, 0755)
			if err != nil {
				fmt.Println("There was an error creating the directory")
			}
		}

		//fmt.Println("create: " + strings.Join(args, " "))
		fmt.Println("file name: ", file)
		fmt.Println("directory name: ", directory)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	makeCmd.Flags().StringVarP(&file, "file", "f", "", "set file name")
	makeCmd.Flags().StringVarP(&directory, "directory", "d", "", "set directory name")
}

