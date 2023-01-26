package main

import (
	"firstProject/services"
	"fmt"
)

var linesPerChunk = 1000

func main() {
	records, err := services.ReadMyCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	filteredRecords, err := services.FilterEmptySpace(records)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The formated CSV record contains %v lines\n", len(filteredRecords))

	finalRecords, err := services.SplitIntoChunks(linesPerChunk, filteredRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = services.WriteFiles(finalRecords)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully added the records into separate files")

}
