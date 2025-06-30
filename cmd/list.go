/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	//"strings"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func isolateTimeStamp(tasks []string) {
	for i, task := range tasks[1:] {
		parts := strings.Split(task, ",")
		createdTime, err := time.ParseInLocation("2006-01-02 15:04:05", parts[3], time.Now().Location())
		if err != nil {
			log.Fatal("could not convert time: ", err)
		}
		createdTime = createdTime.In(time.Now().Location())
		created := timediff.TimeDiff(createdTime)
		parts[3] = created

		tasks[i+1] = strings.Join(parts, ",") // we add 1 to i since we are skipping the header line
	}
}

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
		isolateTimeStamp(content)
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
