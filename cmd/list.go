/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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
		w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)

		defer w.Flush()
		content, err := csvToArray("tasks.csv")
		if err != nil {
			log.Fatal("error: ", err)
		}
		tabulate(content)
		fmt.Fprintln(w, content)
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
