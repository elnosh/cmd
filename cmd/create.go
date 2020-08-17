package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var file string
var directory string

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Make a new file or directory",
	Run: func(cmd *cobra.Command, args []string) {

		//fmt.Println("create: " + strings.Join(args, " "))
		fmt.Println("file name: ", file)
		//fmt.Println("directory name: ", directory)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)


	createCmd.Flags().StringVarP(&file, "file", "f", "", "set file name")
	createCmd.Flags().StringVarP(&directory, "directory", "d", "", "set directory name")

}

