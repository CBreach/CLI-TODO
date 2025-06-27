/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	//"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a new task to your list of tasks",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("task name parameter required")
			return
		}

		//lets first check if the file exists in the system
		if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
			fmt.Println("file does not exits, it'll be created")
		} else {
			fmt.Println("the file exits, it'll be opened")
		}
		file, err := openFile("edit", "tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		} else {
			fmt.Println("the file opened correctly")
			initializeCSV(file)
			addToList(file, args)
			defer file.Close()

		}

		fmt.Println("item:", strings.Join(args, " "), "added to the list.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
