/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a task from the list of TODO's",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("error, task ID not provided")
		}
		taskIdStr := args[0]
		id, err := strconv.Atoi(taskIdStr)
		if err != nil {
			log.Fatal("Task ID must be a valid, positive Integer")
		}
		content, err := csvToArray("tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		}
		newList, err := removeTask(content, id)
		if err != nil {
			log.Fatal("could not remove task from list, probably invalid ID")
		}
		updateID(newList)
		//now we can simply overwrite the csv file
		file, err := openFile("overwrite", "tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		}
		file.Close()
		file, err = openFile("edit", "tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		}
		defer file.Close()
		_, err = file.Write([]byte{}) // test write
		if err != nil {
			log.Fatal("Can't write to file:", err)
		}

		initializeCSV(file)
		addToList(file, newList)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
