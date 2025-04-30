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

	_, err = getFileContents(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// err := processBrainFuck(brainFuckCode)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Output written to", outFile)
}

// func parseCharacter(char uint8) {
// 	switch char {
// 	case '>':
// 		fmt.Println("Move right")
// 	case '<':
// 		fmt.Println("Move left")
// 	case '+':
// 		fmt.Println("Increment value")
// 	case '-':
// 		fmt.Println("Decrement value")
// 	case '.':
// 		fmt.Println("Output value")
// 	case ',':
// 		fmt.Println("Input value")
// 	case '[':
// 		fmt.Println("Start loop")
// 	case ']':
// 		fmt.Println("End loop")
// 	default:
// 		return
// 	}
// }

func getFileName(args []string) (string, error) {
	if len(args) != 2 {
		return "", ErrInvalidCLIArgsLength
	}
	return args[1], nil
}

func getFileContents(filename string) (string, error) {
	// Simulate file reading
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", ErrFileDoesNotExist
	}
	if !strings.HasSuffix(filename, ".bf") {
		return "", ErrBadFileFormat
	}
	return string(data), nil
}
