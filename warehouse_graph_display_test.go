package main

import (
	"testing"
)

func TestPrint(t *testing.T) {
	warehouse := CreateWarehouseGraph(2, 2)

	warehouse.Print()
	// Output:
	// [][][]
	// [][][]
	// [][][]
}
