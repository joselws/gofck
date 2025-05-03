package main

import (
	"testing"

	"github.com/joselws/go-utils/stack"
)

func TestHandleMoveRight(t *testing.T) {
	// Test case for moving the pointer to the right
	validPointer := uint16(0)
	invalidPointer := uint16(maxDataSize - 1)

	err := handleMoveRight(&validPointer)
	if validPointer != 1 || err != nil {
		t.Errorf("Expected pointer to move right to 1, got %d", validPointer)
	}
	err = handleMoveRight(&invalidPointer)
	if err != ErrInvalidPointer {
		t.Errorf("Expected pointer to move right to 1, got %d", validPointer)
	}
}

func TestHandleMoveLeft(t *testing.T) {
	// Test case for moving the pointer to the left
	validPointer := uint16(maxDataSize - 1)
	invalidPointer := uint16(0)

	err := handleMoveLeft(&validPointer)
	if validPointer == 0 || err != nil {
		t.Errorf("Expected pointer to move left to 1, got %d", validPointer)
	}
	err = handleMoveLeft(&invalidPointer)
	if err != ErrInvalidPointer {
		t.Errorf("Expected pointer to move left to 1, got %d", invalidPointer)
	}
}

func TestHandleIncrement(t *testing.T) {
	// Test case for incrementing the value at the current pointer
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)

	handleIncrement(&dataPointers, &currentPointer)
	if dataPointers[currentPointer] != 1 {
		t.Errorf("Expected value at pointer %d to be 1, got %d", currentPointer, dataPointers[currentPointer])
	}
}

func TestHandleDecrement(t *testing.T) {
	// Test case for incrementing the value at the current pointer
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)
	dataPointers[currentPointer] = 2

	handleDecrement(&dataPointers, &currentPointer)
	if dataPointers[currentPointer] != 1 {
		t.Errorf("Expected value at pointer %d to be 1, got %d", currentPointer, dataPointers[currentPointer])
	}
}

func TestHandleLoopStart(t *testing.T) {
	// Test case for starting a loop
	content := []byte("[>++<-]+")
	contentIndex := 0
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)
	dataPointers[currentPointer] = 1 // don't skip loop
	loopPointers := stack.NewStack[int]()

	err := handleLoopStart(content, &contentIndex, loopPointers, &dataPointers, &currentPointer)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if content[contentIndex] != '[' {
		t.Errorf("Expected content char to be at [, got %c", content[contentIndex])
	}
	if loopPointers.IsEmpty() {
		t.Errorf("Expected loop pointers to not be empty, got %d", loopPointers.Len())
	}
}

func TestHandleLoopStartSkipLoop(t *testing.T) {
	// Test case for starting a loop
	content := []byte("[>++<-]+")
	contentIndex := 0
	loopPointers := stack.NewStack[int]()
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)

	err := handleLoopStart(content, &contentIndex, loopPointers, &dataPointers, &currentPointer)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if content[contentIndex] != ']' {
		t.Errorf("Expected content index to be at the end (6), got %d", contentIndex)
	}
	if dataPointers[0] != 0 {
		t.Errorf("Expected value at pointer %d to be 0, got %d", currentPointer, dataPointers[currentPointer])
	}
	if currentPointer != 0 {
		t.Errorf("Expected current pointer to be 0, got %d", currentPointer)
	}
	if !loopPointers.IsEmpty() {
		t.Errorf("Expected loop pointers to be empty, got %d", loopPointers.Len())
	}
}

// Move the pointer to the matching [ index if the value at the current pointer is not zero
func TestHandleLoopEndLoopBack(t *testing.T) {
	// Test case for ending a loop
	// content := []byte("[>++<-]+")
	contentIndex := 6
	loopPointers := stack.NewStack[int]()
	loopPointers.Push(-1) // push the start of the loop
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)
	dataPointers[currentPointer] = 1 // don't move forward

	handleLoopEnd(&contentIndex, loopPointers, &dataPointers, &currentPointer)
	if !loopPointers.IsEmpty() {
		t.Errorf("Expected loop pointers to be empty, got %d", loopPointers.Len())
	}
	if contentIndex != -1 {
		t.Errorf("Expected content index to be at -1, not %d", contentIndex)
	}
}

// Move the pointer to the next character if the value at the current pointer is zero
func TestHandleLoopEndGoForward(t *testing.T) {
	// Test case for ending a loop
	content := []byte("[>++<-]+")
	contentIndex := 6
	loopPointers := stack.NewStack[int]()
	loopPointers.Push(-1) // push the start of the loop
	dataPointers := [maxDataSize]uint8{}
	currentPointer := uint16(0)
	dataPointers[currentPointer] = 0 // move forward

	handleLoopEnd(&contentIndex, loopPointers, &dataPointers, &currentPointer)
	if !loopPointers.IsEmpty() {
		t.Errorf("Expected loop pointers to be empty, got %d", loopPointers.Len())
	}
	if content[contentIndex] != ']' {
		t.Errorf("Expected content index to be at ], got %c", content[contentIndex])
	}
}
