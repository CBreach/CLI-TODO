package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func initializeCSV(file *os.File) {
	fmt.Println("we are here yall")
	info, err := os.Stat(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	if info.Size() == 0 {
		headers := []string{"ID", "TASK"}
		w := csv.NewWriter(file)
		if err := w.Write(headers); err != nil {
			log.Fatal("error writing to the csv: ", err)
		}
		w.Flush()
	}

}
func getTaskId(file *os.File) int {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return len(records)
}
func addToList(file *os.File, item []string) {
	currId := getTaskId(file)
	w := csv.NewWriter(file)

	if err := w.Write(item); err != nil {

	}
}
