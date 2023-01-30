package services

import (
	"errors"
	"testing"
)

var errReader = errors.New("fake reader")
var errWriter = errors.New("fake writer")

func fakeReader(dataLocation string) ([][]string, error) {
	return nil, errReader
}

func fakeWriter(records [][][]string) error {
	return errWriter
}

func TestSolver(t *testing.T) {
	type testCaseSolver struct {
		name  string
		input error
		want  error
	}

	var testsSolver = []testCaseSolver{
		{
			name:  "fake reader",
			input: Solver(fakeReader, FilterEmptySpace, SplitIntoChunks, WriteFiles),
			want:  errReader,
		},
		{
			name:  "fake writer",
			input: Solver(ReadMyCSV, FilterEmptySpace, SplitIntoChunks, fakeWriter),
			want:  errWriter,
		},
	}

	for _, tt := range testsSolver {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input
			if got != tt.want {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}
