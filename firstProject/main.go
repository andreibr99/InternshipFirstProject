package main

import (
	"firstProject/services"
	"fmt"
)

func main() {
	err := services.Solver(services.ReadMyCSV, services.FilterEmptySpace, services.SplitIntoChunks, services.WriteFiles)
	if err != nil {
		fmt.Printf("unable to run program, reason: %v", err)
	}
	fmt.Printf("done")
}
