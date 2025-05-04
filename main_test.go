package main

import "testing"

func TestGetFileName(t *testing.T) {
	tests := []struct {
		args           []string
		expectedResult string
		expectedError  error
	}{
		{[]string{"arg1", "arg2"}, "arg2", nil},
		{[]string{"arg1"}, "", ErrInvalidCLIArgsLength},
		{[]string{}, "", ErrInvalidCLIArgsLength},
	}

	for _, test := range tests {
		result, err := getFileName(test.args)
		if err != test.expectedError {
			t.Errorf("getParameters(%v) = %v; expected %v", test.args, err, test.expectedError)
		}
		if len(result) != len(test.expectedResult) {
			t.Errorf("getParameters(%v) = %v; expected %v", test.args, result, test.expectedResult)
		}
	}
}

func TestFileDoesNotExist(t *testing.T) {
	tests := []struct {
		filename      string
		expectedError error
	}{
		{"nonexistentfile.txt", ErrFileDoesNotExist},
		{"existingfile.txt", ErrBadFileFormat},
		{"example.bf", nil},
		{"example2.b", nil},
	}

	for _, test := range tests {
		_, err := getFileContents(test.filename)
		if err != test.expectedError {
			t.Errorf("Got error: %v on %v; expected: %v", err, test.filename, test.expectedError)
		}
	}
}
