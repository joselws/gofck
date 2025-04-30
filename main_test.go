package main

import "testing"

func TestGetParameters(t *testing.T) {
	tests := []struct {
		args           []string
		expectedResult []string
		expectedError  error
	}{
		{[]string{"arg1", "arg2"}, []string{"arg1", "arg2"}, nil},
		{[]string{"arg1"}, []string{}, ErrInvalidCLIArgsLength},
	}

	for _, test := range tests {
		result, err := getParameters(test.args)
		if err != test.expectedError {
			t.Errorf("getParameters(%v) = %v; expected %v", test.args, err, test.expectedError)
		}
		if len(result) != len(test.expectedResult) {
			t.Errorf("getParameters(%v) = %v; expected %v", test.args, result, test.expectedResult)
		}
	}
}
