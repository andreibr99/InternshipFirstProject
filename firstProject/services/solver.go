package services

import "fmt"

func Solver(
	reader func(recordsLocation string) ([][]string, error),
	filter func(records [][]string) ([][]string, error),
	splitter func(chunkLines int, records [][]string) ([][][]string, error),
	writer func(records [][][]string) error,
) error {
	recordsLocation := "input.csv"
	linesPerChunk := 150

	records, err := reader(recordsLocation)
	if err != nil {
		fmt.Printf("unable to read the data, reason: %v\n", err)
		return err
	}
	fmt.Println("successfully read CSV file")

	filteredRecords, err := filter(records)
	if err != nil {
		fmt.Printf("unable to filter the data, reason: %v\n", err)
		return err
	}
	if len(filteredRecords) == 0 {
		fmt.Println("all data is invalid")
		return err
	}
	fmt.Printf("the formated CSV record contains %v lines\n", len(filteredRecords))

	finalRecords, err := splitter(linesPerChunk, filteredRecords)
	if err != nil {
		fmt.Printf("unable to split the data, reason: %v\n", err)
		return err
	}

	err = writer(finalRecords)
	if err != nil {
		fmt.Printf("unable to write the data, reason: %v\n", err)
		return err
	}
	fmt.Println("successfully added the records into separate files")
	return nil
}
