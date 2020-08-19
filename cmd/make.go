package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fileFlag *bool
var dirFlag *bool

var makeCmd = &cobra.Command{
	Use: "mk",
	Short: "Make a new file",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(args)

		//if file != "" {
		//	f, err := os.Create(file)
		//	if err != nil {
		//		fmt.Println("There was an error creating file.")
		//	}
		//	defer f.Close()
		//}
		//
		//if directory != "" {
		//	err := os.Mkdir(directory, 0755)
		//	if err != nil {
		//		fmt.Println("There was an error creating the directory")
		//	}
		//}
		//
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

