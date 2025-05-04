package main

import (
	"fmt"

	"github.com/joselws/go-utils/stack"
)

const maxDataSize uint16 = 30000

func ProcessBrainFuck(content []byte) (string, error) {
	dataPointers := [maxDataSize]byte{}
	var currentPointer uint16
	loopPointers := stack.NewStack[int]()
	contentSize := len(content) - 1
	contentIndex := -1
	var output string

	for contentIndex < contentSize {
		contentIndex++
		char := content[contentIndex]
		var err error
		switch char {
		case '>':
			err = handleMoveRight(&currentPointer)
		case '<':
			err = handleMoveLeft(&currentPointer)
		case '+':
			handleIncrement(&dataPointers, &currentPointer)
		case '-':
			handleDecrement(&dataPointers, &currentPointer)
		case '.':
			output = handleOutput(&dataPointers, &currentPointer, output)
		// case ',':
		// 	handleInput(&dataPointers, &currentPointer)
		case '[':
			err = handleLoopStart(content, &contentIndex, loopPointers, &dataPointers, &currentPointer)
		case ']':
			err = handleLoopEnd(&contentIndex, loopPointers, &dataPointers, &currentPointer)
		}
		if err != nil {
			return output, err
		}
	}
	return output, nil
}

func handleMoveRight(currentPointer *uint16) error {
	*currentPointer++
	if *currentPointer >= maxDataSize {
		return ErrInvalidPointer
	}
	return nil
}

func handleMoveLeft(currentPointer *uint16) error {
	if *currentPointer == 0 {
		return ErrInvalidPointer
	}
	*currentPointer--
	return nil
}

func handleIncrement(dataPointers *[maxDataSize]byte, currentPointer *uint16) {
	dataPointers[*currentPointer]++
}

func handleDecrement(dataPointers *[maxDataSize]byte, currentPointer *uint16) {
	dataPointers[*currentPointer]--
}

func handleOutput(dataPointers *[maxDataSize]byte, currentPointer *uint16, output string) string {
	return output + fmt.Sprintf("%c", dataPointers[*currentPointer])
}

// If the byte at the data pointer is zero, skip to the char after the matching ]
func handleLoopStart(content []byte, contentIndex *int, loopPointers *stack.Stack[int], dataPointers *[maxDataSize]byte, currentPointer *uint16) error {
	loopPointers.Push(*contentIndex - 1) // deduct the step you get at the start of each loop
	// skip find the matching ]
	if dataPointers[*currentPointer] == 0 {
		for !loopPointers.IsEmpty() {
			*contentIndex++
			if content[*contentIndex] == '[' {
				loopPointers.Push(*contentIndex)
			}
			if content[*contentIndex] == ']' {
				// pop the index of the matching [ and continue
				_, err := loopPointers.Pop()
				if err != nil {
					return ErrInvalidBFSyntax
				}
			}
		}
	}
	return nil
}

// Only move the content index forward if the value is zero
func handleLoopEnd(contentIndex *int, loopPointers *stack.Stack[int], dataPointers *[maxDataSize]byte, currentPointer *uint16) error {
	startLoopIndex, err := loopPointers.Pop()
	if err != nil {
		return ErrInvalidBFSyntax
	}
	if dataPointers[*currentPointer] != 0 {
		*contentIndex = startLoopIndex
	}
	return nil
}
