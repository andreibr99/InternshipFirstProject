package services

import (
	"errors"
)

// FilterEmptySpace filters a 2D slice of strings, deleting lines that
// contain at least one empty field and returns a 2D slice of strings.
// If there is an error, it will return the error.
func FilterEmptySpace(records [][]string) ([][]string, error) {
	if len(records) == 0 {
		return nil, errors.New("file is empty")
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

func ComputeNoOfChunks(chunkLines int, records [][]string) (noOfChunks int) {
	if len(records)%chunkLines == 0 {
		noOfChunks = len(records) / chunkLines
	} else {
		noOfChunks = len(records)/chunkLines + 1
	}
	if len(records) == 2 {
		noOfChunks = 1
	} //if the data contains just the header and one record
	return noOfChunks
}

// SplitIntoChunks splits a 2D slice of strings into chunks of x lines and
// returns a 3D slice of strings. First field of the 3D slice will be the index of
// the chunk and every chunk will have the header as the first line.
func SplitIntoChunks(chunkLines int, records [][]string) ([][][]string, error) {
	if chunkLines <= 0 {
		return nil, errors.New("invalid number of lines per chunk")
	}
	header := records[0]
	records = records[1:]
	line := chunkLines
	finalRecords := make([][][]string, ComputeNoOfChunks(chunkLines, records))
	for i := range finalRecords {
		if i == 0 {
			finalRecords[i] = records[:line]
		} else if i == len(finalRecords)-1 {
			finalRecords[i] = records[line:]
		} else {
			finalRecords[i] = records[line : line+chunkLines]
			line += chunkLines
		}
		finalRecords[i] = append([][]string{header}, finalRecords[i]...)
	}
	return finalRecords, nil
}
