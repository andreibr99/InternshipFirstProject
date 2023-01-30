package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadMyCSV reads a CSV input file, it prints a message if
// the file was successfully read and returns a 2D slice of strings with
// the records from the csv file. If there is an error, it will return the error.
func ReadMyCSV(recordsLocation string) ([][]string, error) {
	file, err := os.Open(recordsLocation)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if err = file.Close(); err != nil {
		fmt.Println(err)
	}

	return records, err
}
