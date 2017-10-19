package main

import "testing"

func TestcountSwears(t *testing.T) {
	/*cases := []struct {
		name           string
		input1          *KnownSwears
		input2			*[]string
		expectedOutput *SwearCounts
	}{
		{
			name:           "empty file",
			input:          [{"fuck", false}, {"shit"}],
			expectedOutput: 0,
		},
		{
			name:           "single character",
			input:          "a",
			expectedOutput: "a",
		},
		{
			name:           "two characters",
			input:          "ab",
			expectedOutput: "ba",
		},
		{
			name:           "stressed",
			input:          "stressed",
			expectedOutput: "desserts",
		},
		{
			name:           "high unicode",
			input:          "Hello, 世界",
			expectedOutput: "界世 ,olleH",
		},
	}

	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}*/
}

func TestloadKnownSwears(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput *KnownSwears
	}{
		{
			name:           "empty file",
			input:          "",
			expectedOutput: &map[string]bool,
		},
	}
	for _, c := range cases {
		if output := Reverse(c.input); output != c.expectedOutput {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

func TestloadWords(t *testing.T) {

}

func TestopenFile(t *testing.T) {
	// hint os.Args = []string{"cmd", "path/to/file.txt"}
}
