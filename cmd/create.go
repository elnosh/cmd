package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	log "github.com/sirupsen/logrus"
)

var file string
var directory string

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Make a new file or directory",
	Run: func(cmd *cobra.Command, args []string) {

		if file != "" {
			fmt.Println("file name: ", file)

			f, err := os.Create(file)
			if err != nil {
				log.Error("There was an error while creating file")
			}
			defer f.Close()
		}

		//fmt.Println("create: " + strings.Join(args, " "))
		//fmt.Println("directory name: ", directory)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)


	createCmd.Flags().StringVarP(&file, "file", "f", "", "set file name")
	createCmd.Flags().StringVarP(&directory, "directory", "d", "", "set directory name")

}

