package main

import (
	"testing"
)

func TestPrintSuccessEmoji(t *testing.T) {
	printSuccessEmoji()
	// Output: 😎
}

func TestPrintEndEmoji(t *testing.T) {
	printEndEmoji()
	// Output: 🙂
}

func TestIsGameFinished(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	if isGameFinished(warehouse) != true {
		t.Errorf("it does not detect the end of game")
	}
}
