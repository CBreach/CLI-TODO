package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func openFile(mode string, file string) (*os.File,error){
	if mode == "edit"{
		return nil
	}else if mode == "overwrite"{
		f, err := os.OpenFile("task.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) 
		return nil
	}
	else{
		return nil,fmt.Errorf("invalid mode parameter")
	}
}
func initializeCSV(file *os.File) {
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

func updateID(content []string){
	for i, line := range(content){
		if i > 0{
			_, updatedEntry, found := strings.Cut(line, ",")
			if found {
				content[i] = updatedEntry
			}
		}
	}
}
func removeTask(records []string, ID int) ([]string, error){
	fmt.Println("this is the length of the arr", len(records))
	if ID <= 0 || ID >= len(records){
		return records, fmt.Errorf("ID does not exist: %d", ID)
	}
	half1 := records[:ID]
	half2 := records [ID+1:]
	return append(half1, half2...), nil
	

}
