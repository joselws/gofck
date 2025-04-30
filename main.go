package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inFile, err := getFileName(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	brainFuckCode, err := getFileContents(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ProcessBrainFuck(brainFuckCode)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFileName(args []string) (string, error) {
	if len(args) != 2 {
		return "", ErrInvalidCLIArgsLength
	}
	return args[1], nil
}

func getFileContents(filename string) ([]byte, error) {
	// Simulate file reading
	data, err := os.ReadFile(filename)
	if err != nil {
		return data, ErrFileDoesNotExist
	}
	if !strings.HasSuffix(filename, ".bf") {
		return data, ErrBadFileFormat
	}
	return data, nil
}
