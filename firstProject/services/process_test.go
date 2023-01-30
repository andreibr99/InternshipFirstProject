package services

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type testCase struct {
	name    string
	records [][]string
}

var tests = []testCase{
	{
		name: "test 1",
		records: [][]string{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
			{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
			{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		},
	},
	{
		name: "test 2",
		records: [][]string{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"25", "May", "Cumes", "mcumeso@mac.com", "Male", "28.113.141.184"},
			{"26", "Lucie", "Ager", "lagerp@slideshare.net", "Female", "88.112.235.215"},
			{"27", "Hayden", "Gioani", "hgioaniq@narod.ru", "Female", "26.46.19.2"},
			{"28", "Alano", "Abrahamsson", "aabrahamssonr@kickstarter.com", "Non-binary", "177.171.154.186"},
			{"29", "Casey", "Fossord", "cfossords@europa.eu", "Female", "60.203.145.197"},
			{"30", "Geoff", "Alliban", "gallibant@dedecms.com", "Bigender", "212.150.89.224"},
			{"31", "Concordia", "Ginie", "cginieu@webnode.com", "Male", "3.229.203.224"},
			{"33", "Leyla", "St Angel", "lstangelw@kickstarter.com", "Female", "72.28.90.48"},
			{"35", "Deirdre", "Mallan", "dmallany@live.com", "Male", "116.23.23.17"},
			{"36", "Bria", "Curtain", "bcurtainz@yolasite.com", "Female", "91.170.234.184"},
			{"37", "Tilly", "Dredge", "tdredge10@wsj.com", "Male", "213.231.29.86"},
			{"38", "Keefer", "York", "kyork11@diigo.com", "Female", "230.52.23.219"}},
	},
	{
		name: "test 3",
		records: [][]string{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"}},
	},
}

func TestFilterEmptySpace(t *testing.T) {
	type filterTestCase struct {
		name  string
		input [][]string
		want  [][]string
	}
	var tests = []filterTestCase{
		{
			name: "input contains invalid data, output contains only valid data",
			input: [][]string{
				{"25", "May", "Cumes", "mcumeso@mac.com", "Male", "28.113.141.184"},
				{"45", "", "Gegay", "", "Polygender", ""},
				{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"},
				{"46", "Binny", "", "bcannan19@mediafire.com", "Male", "113.171.207.158"},
				{""}},
			want: [][]string{
				{"25", "May", "Cumes", "mcumeso@mac.com", "Male", "28.113.141.184"},
				{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"}},
		},
		{
			name:  "input is nil, output is nil",
			input: nil,
			want:  nil,
		},
		{
			name: "all records are invalid, output is nil",
			input: [][]string{
				{"", "May", "", "mcumeso@mac.com", "", "28.113.141.184"},
				{"45", "", "Gegay", "", "Polygender", ""},
				{"49", "", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"},
				{"46", "Binny", "", "bcannan19@mediafire.com", "Male", "113.171.207.158"},
				{""}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := FilterEmptySpace(tt.input)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Function result: %v, expected result: %v", got, tt.want)
			}
		})
	}
}

func TestFilterEmptySpace2(t *testing.T) {
	var records [][]string
	got, err := FilterEmptySpace(records)
	if got != nil || err == nil {
		t.Fatal("for empty file the err should not be nil")
	}
}

func TestComputeNoOfChunks(t *testing.T) {
	numberOfLines := []int{4, 2, 1}
	expectedOutput := []int{1, 7, 1}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeNoOfChunks(numberOfLines[i], tt.records)
			if got != expectedOutput[i] {
				t.Errorf("The result: %v does not match the expected output: %v", got, expectedOutput[i])
			}
		})
	}
}

func TestSplitIntoChunks(t *testing.T) {
	numberOfLines := []int{2, 4, 1}
	var expectedOutput [3][][][]string
	expectedOutput[0] = [][][]string{
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
			{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		},
	}
	expectedOutput[1] = [][][]string{
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"25", "May", "Cumes", "mcumeso@mac.com", "Male", "28.113.141.184"},
			{"26", "Lucie", "Ager", "lagerp@slideshare.net", "Female", "88.112.235.215"},
			{"27", "Hayden", "Gioani", "hgioaniq@narod.ru", "Female", "26.46.19.2"},
			{"28", "Alano", "Abrahamsson", "aabrahamssonr@kickstarter.com", "Non-binary", "177.171.154.186"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"29", "Casey", "Fossord", "cfossords@europa.eu", "Female", "60.203.145.197"},
			{"30", "Geoff", "Alliban", "gallibant@dedecms.com", "Bigender", "212.150.89.224"},
			{"31", "Concordia", "Ginie", "cginieu@webnode.com", "Male", "3.229.203.224"},
			{"33", "Leyla", "St Angel", "lstangelw@kickstarter.com", "Female", "72.28.90.48"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"35", "Deirdre", "Mallan", "dmallany@live.com", "Male", "116.23.23.17"},
			{"36", "Bria", "Curtain", "bcurtainz@yolasite.com", "Female", "91.170.234.184"},
			{"37", "Tilly", "Dredge", "tdredge10@wsj.com", "Male", "213.231.29.86"},
			{"38", "Keefer", "York", "kyork11@diigo.com", "Female", "230.52.23.219"},
		},
	}
	expectedOutput[2] = [][][]string{
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := SplitIntoChunks(numberOfLines[i], tt.records)
			if diff := cmp.Diff(got, expectedOutput[i]); diff != "" {
				t.Errorf("Function result: %v, expected result: %v", got, expectedOutput[i])
			}
		})
	}
}
