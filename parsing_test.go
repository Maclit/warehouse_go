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
		return
	}
	if len(expectedString) != len(resultString) {
		t.Errorf("size is different between the two different array expectedString: %d resultString: %d", len(expectedString), len(resultString))
		return
	}
	for n := 0; n < len(resultString); n++ {
		if resultString[n] != expectedString[n] {
			t.Errorf("the result are different %s and %s", resultString[n], expectedString[n])
			return
		}
	}
}

// Verify if the the function catch the error
func TestDivideError(t *testing.T) {
	inputString := "5 test 100 5 6 7 "
	resultString, err := divide(inputString)
	if err == nil || len(resultString) > 0 {
		t.Errorf("there is an error in inputString")
		return
	}
}

// TestInfoMap verify if the size of the map in parameter is good
func TestInfoMap(t *testing.T) {
	infoMap, err := infoMap([]string{"5", "5", "1000"})
	expectedResult := []int{5, 5, 1000}
	if err != nil {
		t.Errorf("there is an error in inputString")
		return
	}
	if len(infoMap) != len(expectedResult) {
		t.Errorf("size is different between the two different array expectedString: %d resultString: %d", len(expectedResult), len(infoMap))
		return
	}
	for n := 0; n < len(infoMap); n++ {
		if infoMap[n] != expectedResult[n] {
			t.Errorf("the result are different %d and %d", infoMap[n], expectedResult[n])
			return
		}
	}
}

// TestInfoMap verify if the programm catchthe error
func TestInfoMapError(t *testing.T) {
	infoMap, err := infoMap([]string{"-5", "5", "1000"})
	if err == nil || len(infoMap) > 0 {
		t.Errorf("there is an error in inputString")
		return
	}
}

// Test if InfoColis works with GREEN
func TestInfoColisGreen(t *testing.T) {
	infoNameColis, infoColis, err := infoColis([]string{"name", "5", "1000", "gReEn"})
	expectedResult := []int{5, 1000, GREEN}
	expectedNameColis := "name"
	if err != nil {
		t.Errorf("there is an error in inputString")
		return
	}
	if infoNameColis != expectedNameColis {
		t.Errorf("the names are different%s and %s", infoNameColis, expectedNameColis)
		return
	}
	if len(infoColis) != len(expectedResult) {
		t.Errorf("the size for the position of box are different %d and %d", len(infoColis), len(expectedResult))
		return
	}
	for n := 0; n < len(infoColis); n++ {
		if infoColis[n] != expectedResult[n] {
			t.Errorf("the result are different %d and %d", infoColis[n], expectedResult[n])
			return
		}
	}
}

// Test if InfoColis works with YELLOW
func TestInfoColisYellow(t *testing.T) {
	infoNameColis, infoColis, err := infoColis([]string{"name", "5", "1000", "YELLOW"})
	expectedResult := []int{5, 1000, YELLOW}
	expectedNameColis := "name"
	if err != nil {
		t.Errorf("there is an error in inputString")
		return
	}
	if infoNameColis != expectedNameColis {
		t.Errorf("the names are different%s and %s", infoNameColis, expectedNameColis)
		return
	}
	if len(infoColis) != len(expectedResult) {
		t.Errorf("the size for the position of box are different %d and %d", len(infoColis), len(expectedResult))
		return
	}
	for n := 0; n < len(infoColis); n++ {
		if infoColis[n] != expectedResult[n] {
			t.Errorf("the result are different %d and %d", infoColis[n], expectedResult[n])
			return
		}
	}
}

// Test if InfoColis works with BLUE
func TestInfoColisBLUE(t *testing.T) {
	infoNameColis, infoColis, err := infoColis([]string{"name", "5", "1000", "BLue"})
	expectedResult := []int{5, 1000, BLUE}
	expectedNameColis := "name"
	if err != nil {
		t.Errorf("there is an error in inputString")
		return
	}
	if infoNameColis != expectedNameColis {
		t.Errorf("the names are different%s and %s", infoNameColis, expectedNameColis)
		return
	}
	if len(infoColis) != len(expectedResult) {
		t.Errorf("the size for the position of box are different %d and %d", len(infoColis), len(expectedResult))
		return
	}
	for n := 0; n < len(infoColis); n++ {
		if infoColis[n] != expectedResult[n] {
			t.Errorf("the result are different %d and %d", infoColis[n], expectedResult[n])
			return
		}
	}
}

// Test if Color does not match
func TestInfosColisErrorColor(t *testing.T) {
	infoNameColis, infoColis, err := infoColis([]string{"name", "5", "1000", "Red"})
	if err == nil || len(infoColis) > 0 || infoNameColis != "" {
		t.Errorf("the error red is not caught")
	}
}

// Test if pos is not correct

func TestInfosColisErrorPos(t *testing.T) {
	infoNameColis, infoColis, err := infoColis([]string{"name", "-5", "1000", "Red"})
	if err == nil || len(infoColis) > 0 || infoNameColis != "" {
		t.Errorf("the error negative number is not caught")
	}
}

// Test InfoPaletteCamion
func TestInfoPaletteCamion(t *testing.T) {
	infoName, infoPos, err := infoPaletteCamion([]string{"name", "5", "5"})
	expectedName := "name"
	expectedInfo := []int{5, 5}

	if err != nil {
		t.Errorf("an error occur in functionInfoPaletteCamion")
		return
	}
	if expectedName != infoName {
		t.Errorf("the name are different %s and %s", expectedName, infoName)
		return
	}
	if len(expectedInfo) != len(infoPos) {
		t.Errorf("the size of pos are different %d and %d", len(expectedInfo), len(infoPos))
		return
	}
	for n := 0; n < len(expectedInfo); n++ {
		if expectedInfo[n] != infoPos[n] {
			t.Errorf("the position are not the same %d and %d", expectedInfo[n], infoPos[n])
			return
		}
	}
}

func TestInfoPaletteCamionErrorPos(t *testing.T) {
	infoName, infoPos, err := infoPaletteCamion([]string{"name", "-5", "5"})
	if err == nil || len(infoPos) > 0 || infoName != "" {
		t.Errorf("the error negative number is not caught")
	}
}
