package main

import "testing"

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
