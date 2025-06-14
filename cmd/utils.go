package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"io"
)
func openFile(mode string, file string) (*os.File,error){
	switch mode {
    case "edit":
        return os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    case "overwrite":
        return os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    default:
        return nil, fmt.Errorf("invalid mode parameter: %s", mode)
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
	_, err := file.Seek(0,0)
	if err != nil{
		log.Fatal("Could not seek to beginning: ", err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return len(records)
}
func addToList(file *os.File, tasks []string) {
	// Ensure we're writing at the end of the file
	_, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		log.Fatal("Could not seek to end: ", err)
	}

	// Count current lines (ID base)
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal("Could not seek to beginning: ", err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	currId := len(records)

	// Back to end for writing
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		log.Fatal("Could not seek to end (again): ", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, task := range tasks {
		record := []string{strconv.Itoa(currId), task}
		if err := w.Write(record); err != nil {
			log.Fatalln(err)
		}
		currId++
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
