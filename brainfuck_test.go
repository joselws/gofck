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
