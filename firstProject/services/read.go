package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadMyCSV reads a CSV input file, it prints a message if
// the file was successfully read and returns a 2D slice of strings with
// the records from the csv file. If there is an error, it will return the error.
func ReadMyCSV() ([][]string, error) {
	var err error
	file, err := os.Open("input.csv")
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, err
}
