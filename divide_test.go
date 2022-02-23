package main

import (
	"testing"
)

// TestDivide verify if Divide create an array corresponding to the original string
func TestDivide(t *testing.T) {
	inputString := "5 test 100"
	expectedString := []string{"5", "test", "100"}
	resultString, err := divide(inputString)
	if err != nil {
		t.Errorf("there is an error in inputString")
	}
	if len(expectedString) != len(resultString) {
		t.Errorf("size is different between the two different array expectedString: %d resultString: %d", len(expectedString), len(resultString))
	}
	for n := 0; n < len(resultString); n++ {
		if resultString[n] != expectedString[n] {
			t.Errorf("the result are different %s and %s", resultString[n], expectedString[n])
		}
	}
}
