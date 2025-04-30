package main

import (
	"github.com/joselws/go-utils/stack"
)

const maxDataSize uint16 = 50000

func ProcessBrainFuck(content []byte) error {
	dataPointers := [maxDataSize]uint8{}
	var currentPointer uint16
	loopPointers := stack.NewStack[uint16]()

	for _, char := range content {
		processCharacter(char, &dataPointers, &currentPointer, loopPointers)
	}
	return nil
}

func processCharacter(char byte, dataPointers *[50000]uint8, currentPointer *uint16, loopPointers *stack.Stack[uint16]) {
	var err error
	switch char {
	case '>':
		err = handleMoveRight(currentPointer)
	case '<':
		err = handleMoveLeft(currentPointer)
	case '+':
		handleIncrement(dataPointers, currentPointer)
	case '-':
		handleDecrement(dataPointers, currentPointer)
	// case '.':
	// 	handleOutput()
	// case ',':
	// 	return
	// case '[':
	// 	handleLoopStart()
	// case ']':
	// 	handleLoopEnd()
	default:
		return
	}
	panic(err)
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
