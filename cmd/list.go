/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	//"strings"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "displays all current tasks to the user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		w := new(tabwriter.Writer)

		//lets define the format
		w.Init(os.Stdout, 0, 0, 3, ' ', 0)

		defer w.Flush()
		content, err := csvToArray("tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		}
		tabulate(content)

		for _, line := range content {
			fmt.Fprintln(w, line) // passes individual fields
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
