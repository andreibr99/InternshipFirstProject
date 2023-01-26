package services

import (
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
			{"2", "Fan", "", "fgilvear1@people.com.cn", "Female", "125.219.253.132"},
			{"3", "Gerri", "Choffin", "gchoffin2@ning.com", "", "9.254.198.50"},
			{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
			{"5", "Benoite", "Jaffray", "bjaffray4@github.com", "Female", ""},
			{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
			{"7", "", "Mc Dermid", "emcdermid6@plala.or.jp", "", "72.200.10.99"},
			{"8", "Andrea", "", "amckeran7@example.com", "", ""},
			{""}},
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
			{"32", "Berenice", "Ilymanov", "", "Genderqueer", "112.90.4.181"},
			{"33", "Leyla", "St Angel", "lstangelw@kickstarter.com", "Female", "72.28.90.48"},
			{"", "Keenan", "Beranek", "kberanekx@whitehouse.gov", "Female", "163.190.193.184"},
			{"35", "Deirdre", "Mallan", "dmallany@live.com", "Male", "116.23.23.17"},
			{"36", "Bria", "Curtain", "bcurtainz@yolasite.com", "Female", "91.170.234.184"},
			{"37", "Tilly", "Dredge", "tdredge10@wsj.com", "Male", "213.231.29.86"},
			{"38", "Keefer", "York", "kyork11@diigo.com", "Female", "230.52.23.219"}},
	},
	{
		name: "test 3",
		records: [][]string{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"45", "", "Gegay", "", "Polygender", ""},
			{"46", "Binny", "", "bcannan19@mediafire.com", "Male", "113.171.207.158"},
			{"47", "", "", "ceburne1a@blogtalkradio.com", "Female", "44.218.114.65"},
			{"48", "Lanie", "Jeppensen", "ljeppensen1b@discuz.net", "", "11.5.103.224"},
			{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"}},
	},
}

func TestFilterEmptySpace(t *testing.T) {
	expectedOutput := [][][]string{
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"1", "Mavra", "Malec", "mmalec0@usa.gov", "Female", "229.215.245.102"},
			{"4", "Tremayne", "Loosemore", "tloosemore3@cnn.com", "Male", "167.249.115.222"},
			{"6", "Clint", "Oliphard", "coliphard5@ft.com", "Genderfluid", "39.69.123.72"},
		},
		{
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
			{"38", "Keefer", "York", "kyork11@diigo.com", "Female", "230.52.23.219"},
		},
		{
			{"id", "first_name", "last_name", "email", "gender", "ip_address"},
			{"49", "Errick", "Kealy", "ekealy1c@cisco.com", "Male", "255.219.69.164"},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FilterEmptySpace(tt.records)
			if len(got) != len(expectedOutput[i]) || err != nil {
				t.Fatal("The lengths do not match")
				return
			}
			for j := range got {
				for k, v := range got[j] {
					if v != expectedOutput[i][j][k] {
						t.Fatal("The result does not match the expected output")
					}
				}
			}
		})
	}
}

func TestFilterEmptySpace2(t *testing.T) {
	var records [][]string
	got, err := FilterEmptySpace(records)
	if got != nil || err == nil {
		t.Fatal("for empty records the function fails the test")
	}
}

func TestComputeNoOfChunks(t *testing.T) {
	numberOfLines := []int{
		4,
		2,
		1,
	}

	expectedOutput := []int{
		3,
		8,
		6,
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeNoOfChunks(numberOfLines[i], tt.records)
			if got != expectedOutput[i] {
				t.Fatalf("The result: %v does not match the expected output: %v", got, expectedOutput[i])
			}
		})
	}
}

func TestSplitIntoChunks(t *testing.T) {
	numberOfLines := []int{
		2,
		4,
		1,
	}
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
			got, err := SplitIntoChunks(numberOfLines[i], tt.records)
			if err != nil {
				t.Fatal("err is not nil")
				return
			}
			for j := range got {
				for k := range got[j] {
					for l, v := range got[j][k] {
						if v != expectedOutput[i][j][k][l] {
							t.Fatal("the result does not match the expected output")
						}
					}
				}
			}
		})
	}
}
