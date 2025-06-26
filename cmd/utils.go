package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile(mode string, file string) (*os.File, error) {
	switch mode {
	case "edit":
		return os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	case "overwrite":
		return os.OpenFile(file, os.O_RDWR|os.O_WRONLY|os.O_TRUNC, 0644)
	default:
		return nil, fmt.Errorf("invalid mode parameter: %s", mode)
	}
}
func initializeCSV(file *os.File) {
	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatal("Could not seek to beginning: ", err)
	}
	info, err := os.Stat(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("file, size: ", info.Size())
	if info.Size() == 0 {
		headers := []string{"ID", "TASK"}
		//fmt.Println("Writable?", file.Fd()) // should be > 0
		_, err = file.Write([]byte{})
		if err != nil {
			log.Fatal("Can't write to file:", err)
		}
		w := csv.NewWriter(file)
		if err := w.Write(headers); err != nil {
			log.Fatal("error writing to the csv: ", err)
		}
		w.Flush()
	}

}
func getTaskId(file *os.File) int {
	_, err := file.Seek(0, 0)
	if err != nil {
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

func updateID(content []string) {
	for i, line := range content {
		if line != "" {
			_, updatedEntry, found := strings.Cut(line, ",")
			if found {
				content[i] = updatedEntry
			}
		}
	}
}
func removeTask(records []string, ID int) ([]string, error) {
	if ID <= 0 || ID >= len(records) {
		return records, fmt.Errorf("ID does not exist: %d", ID)
	}
	half1 := records[1:ID]
	half2 := records[ID+1:]
	return append(half1, half2...), nil

}
func csvToArray(file string) ([]string, error) {

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file")
	}
	//convert raw data into content str
	content := strings.Split(string(data), "\n")
	content = content[:len(content)-1] //gets rid of the annoying white empty index
	return content, nil
}
func tabulate(content []string) {
	for i, entry := range content {
		//fmt.Println("this is entry at this point: ", entry)
		content[i] = strings.ReplaceAll(entry, ",", "\t")
		content[i] += "\t" //adding a trailing tab char to fix allignment issues
		//fmt.Printf("%q\n", content[i])
	}
}
