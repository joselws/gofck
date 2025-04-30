package main

import (
	"github.com/joselws/go-utils/stack"
)

func processBrainFuck(content []byte) error {
	const maxDataSize uint16 = 50000
	dataPointers := [maxDataSize]uint8{}
	var currentPointer uint16
	loopPointers := stack.NewStack[uint16]()

	for _, char := range content {
		processCharacter(char, dataPointers, &currentPointer, loopPointers)
	}
	return nil
}

func processCharacter(char byte, dataPointers [50000]uint8, currentPointer *uint16, loopPointers *stack.Stack[uint16]) {
	switch char {
	// case '>':
	// 	handleMoveRight()
	// case '<':
	// 	handleMoveLeft()
	// case '+':
	// 	handleIncrement()
	// case '-':
	// 	handleDecrement()
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
}
