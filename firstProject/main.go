package main

import (
	"firstProject/services"
	"fmt"
)

var X = 100

func main() {
	records, err := services.ReadMyCSV()
	if err != nil {
		fmt.Println(err)
	}

	filteredRecords, err := services.FilterEmptySpace(records)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The formated CSV record contains %v lines\n", len(filteredRecords))
	}

	finalRecords, err := services.SplitIntoChunks(X, filteredRecords)
	if err != nil {
		fmt.Println(err)
	}

	err = services.WriteFiles(finalRecords)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully added the records into separate files")
	}
}
