package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var rmvAll *bool

var removeCmd = &cobra.Command{
	Use:   "rmv",
	Short: "Remove file or directory",

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("  specify file or directory name")
		}

		// remove directory along with content if specified in flag
		if *rmvAll {
			for _, arg := range args {
				if _, err := os.Stat(arg); os.IsNotExist(err) {
					fmt.Println(err)
				} else {
					err := os.RemoveAll(arg)
					checkErr(err)
				}
			}
		} else {
			for _, arg := range args {

				// stat to check if its file or directory
				fi, statErr := os.Stat(arg)
				if statErr != nil {
					fmt.Println(statErr)
					return
				}
				if fi.Mode().IsRegular() {
					err := os.Remove(arg)
					checkErr(err)
				} else if fi.Mode().IsDir() {
					// read directory to check if its empty
					files, err := ioutil.ReadDir(arg)
					if err != nil {
						fmt.Println(err)
						return
					}
					// if directory is not empty. Use flag -a to delete directory along with content.
					if len(files) > 0 {
						fmt.Println("rmv:   Use -a to delete directory along with its content.")
						return
					}

					if os.IsNotExist(statErr) {
						fmt.Println(statErr)
					} else {
						err := os.Remove(arg)
						checkErr(err)
					}
					
				}
			}
		}
	},
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("There was an error deleting the file or directory")
		return
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
	rmvAll = removeCmd.Flags().BoolP("all", "a", false, "remove directory with its content")
}

