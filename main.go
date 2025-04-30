package main

import (
	"fmt"
)

func main() {
	// dataArray := [5000]uint8{}
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints a message to the console.")
	fmt.Println("Goodbye!")
	// var arrayPointer uint16
	// args := os.Args[1:]
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

func getParameters(args []string) ([]string, error) {
	if len(args) != 2 {
		return []string{}, ErrInvalidCLIArgsLength
	}
	return args, nil
}
