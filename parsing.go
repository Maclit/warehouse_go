package main

import (
	"strconv"
	"strings"
)

func infoMap(text []string) (infoMap []int, err error) {
	for n := 0; n < len(text); n++ {
		intVar, err := strconv.Atoi(text[n])
		if err != nil || intVar < 0 {
			return []int{}, MapError(text[n])
		}
		infoMap = append(infoMap, intVar)
	}
	if len(infoMap) != 3 || (infoMap[2] < 10 || infoMap[2] > 10000) {
		return []int{}, MapError("Invalid number of turns")
	}
	return infoMap, nil
}

func infoColis(text []string) (infoNameColis string, infoColis []int, err error) {
	infoNameColis = text[0]

	for n := 1; n < len(text)-1; n++ {
		intVar, err := strconv.Atoi(text[n])
		if err != nil || intVar < 0 {
			ColisPosNull := []int{}
			ColisNull := ""
			return ColisNull, ColisPosNull, MapError(text[n])
		}
		infoColis = append(infoColis, intVar)
	}
	switch lower := strings.ToLower(text[len(text)-1]); lower {
	case "green":
		infoColis = append(infoColis, GREEN)
	case "yellow":
		infoColis = append(infoColis, YELLOW)
	case "blue":
		infoColis = append(infoColis, BLUE)
	default:
		ColisPosNull := []int{}
		ColisNull := ""
		return ColisNull, ColisPosNull, ColorError(lower)
	}
	return infoNameColis, infoColis, nil
}

func infoPaletteCamion(text []string) (infoColis string, infoPosColis []int, err error) {
	infoColis = text[0]

	for n := 1; n < len(text); n++ {
		intVar, err := strconv.Atoi(text[n])
		if err != nil || intVar < 0 {
			ColisPosNull := []int{}
			ColisNull := ""
			return ColisNull, ColisPosNull, MapError(text[n])
		}
		infoPosColis = append(infoPosColis, intVar)
	}
	return infoColis, infoPosColis, nil
}

func findObject(words []string, warehouse *WarehouseSquareGraph) error {
	switch size := len(words); size {
	case 4:
		name, posColor, infoErr := infoColis(words)
		if infoErr != nil {
			return infoErr
		}
		graphErr := warehouse.AddPackage(name, posColor[0], posColor[1], posColor[2])
		if graphErr != nil {
			return graphErr
		}
	case 3:
		name, pos, infoErr := infoPaletteCamion(words)
		if infoErr != nil {
			return infoErr
		}
		graphErr := warehouse.AddTransporter(name, pos[0], pos[1])
		if graphErr != nil {
			return graphErr
		}
	case 5:
		name, posSizeRound, infoErr := infoPaletteCamion(words)
		if infoErr != nil {
			return infoErr
		}
		graphErr := warehouse.AddTruck(name, posSizeRound[0], posSizeRound[1], posSizeRound[2], posSizeRound[3])
		if graphErr != nil {
			return graphErr
		}
	}
	return nil
}

func divide(text string) ([]string, error) {
	words := strings.Fields(text)
	if len(words) < 3 || len(words) > 5 {
		response := []string{}
		length := strconv.Itoa(len(words))
		return response, InputError(length)
	}
	return words, nil
}

func analyzeAllText() (warehouseGraph *WarehouseSquareGraph, round int, err error) {
	text, errRecupMap := readMap()

	if errRecupMap != nil {
		warehouseGraphErr := CreateWarehouseGraph(0, 0)
		return warehouseGraphErr, 0, errRecupMap
	}
	firstLine, err := divide(text[0])
	if err != nil {
		warehouseGraphErr := CreateWarehouseGraph(0, 0)
		return warehouseGraphErr, 0, err
	}
	mapInfo, errMap := infoMap(firstLine)
	if errMap != nil {
		warehouseGraphErr := CreateWarehouseGraph(0, 0)
		return warehouseGraphErr, 0, errMap
	}
	warehouseGraph = CreateWarehouseGraph(mapInfo[0], mapInfo[1])
	for n := 1; n < len(text); n++ {
		line, err := divide(text[n])
		if err != nil {
			warehouseGraphErr := CreateWarehouseGraph(0, 0)
			return warehouseGraphErr, 0, err
		}
		errLine := findObject(line, warehouseGraph)
		if errLine != nil {
			warehouseGraphErr := CreateWarehouseGraph(0, 0)
			return warehouseGraphErr, 0, errLine
		}
	}
	return warehouseGraph, mapInfo[2], nil
}
