package main

import (
	"fmt"
)

const maxDataSize uint16 = 50000

func ProcessBrainFuck(content []byte) error {
	dataPointers := [maxDataSize]uint8{}
	var currentPointer uint16
	// loopPointers := stack.NewStack[int]()
	contentSize := len(content)
	contentIndex := 0

	for contentIndex < contentSize {
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
			handleOutput(&dataPointers, &currentPointer)
			// case '[':
			// 	err = handleLoopStart(content, &contentIndex, &loopPointers, &dataPointers, &currentPointer)
			// case ']':
			// 	handleLoopEnd()
			contentIndex++
		}
		if err != nil {
			return err
		}
	}
	return nil
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

func handleIncrement(dataPointers *[maxDataSize]uint8, currentPointer *uint16) {
	dataPointers[*currentPointer]++
}

func handleDecrement(dataPointers *[maxDataSize]uint8, currentPointer *uint16) {
	dataPointers[*currentPointer]--
}

func handleOutput(dataPointers *[maxDataSize]uint8, currentPointer *uint16) {
	fmt.Print(string(dataPointers[*currentPointer]))
}
