package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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
func addToList(file *os.File, task string) {
	currId := getTaskId(file)
	record := []string{strconv.Itoa(currId), task}
	w := csv.NewWriter(file)
	defer w.Flush()
	if err := w.Write(record); err != nil {
		log.Fatalln(err)
	}
}
func removeTask(records []string, ID int){
	for i, record := range records{
		
	}
}
