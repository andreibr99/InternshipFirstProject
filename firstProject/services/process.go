package services

import (
	"errors"
	"fmt"
)

// FilterEmptySpace filters a 2D slice of strings, deleting lines that
// contain at least one empty field and returns a 2D slice of strings.
// If there is an error, it will return the error.
func FilterEmptySpace(records [][]string) ([][]string, error) {
	if records == nil {
		return nil, errors.New("failed to format file: file is empty")
	}
	var filteredRecords [][]string
	flag := false
	for i, v := range records {
		for j := range v {
			if records[i][j] == "" {
				flag = true
			}
		}
		if !flag {
			filteredRecords = append(filteredRecords, records[i])
		}
		flag = false
	}
	return filteredRecords, nil
}

func computeNoOfChunks(x int, records [][]string) (noOfChunks int) {
	if len(records)%x == 0 {
		noOfChunks = len(records) / x
	} else {
		noOfChunks = len(records)/x + 1
	}
	return noOfChunks
}

// SplitIntoChunks splits a 2D slice of strings into chunks of x lines and
// returns a 3D slice of strings. First field of the 3D slice will be the index of
// the chunk and every chunk will have the header as the first line.
func SplitIntoChunks(x int, records [][]string) ([][][]string, error) {
	records, _ = FilterEmptySpace(records)
	if x > len(records) || x <= 0 || records == nil {
		return nil, fmt.Errorf("failed to split into chunks")
	}
	header := records[0]
	records = records[1:]
	line := x
	finalRecords := make([][][]string, computeNoOfChunks(x, records))
	for i := range finalRecords {
		if i == 0 {
			finalRecords[i] = records[:line]
		} else if i == len(finalRecords)-1 {
			finalRecords[i] = records[line:]
		} else {
			finalRecords[i] = records[line : line+x]
			line += x
		}
		finalRecords[i] = append([][]string{header}, finalRecords[i]...)
	}
	return finalRecords, nil
}
