package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

// WriteFiles receives a 3D slice of strings and if successful creates
// as many files as there are elements in the first field of the slice,
// and writes the contents of field two and three of the slice.
// If there is an error, it will return the error.
func WriteFiles(records [][][]string) error {
	if records == nil {
		return errors.New("failed to add records in files")
	}
	for i := range records {
		fileName := fmt.Sprintf("data%v.csv", i)
		csvFile, err := os.Create(fileName)

		if err != nil {
			return err
		}
		csvWriter := csv.NewWriter(csvFile)

		for _, row := range records[i] {
			_ = csvWriter.Write(row)
		}
		csvWriter.Flush()
		err = csvFile.Close()
		if err != nil {
			return err
		}

	}
	return nil

}
