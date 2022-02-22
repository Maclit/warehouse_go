package main

import (
	"strconv"
	"strings"
)

func infoMap(text []string) (infoMap []int, err error) {
	for n := 0; n < len(text); n++ {
		intVar, err := strconv.Atoi(text[n])
		if err != nil || intVar < 0 {
			MapNull := []int{}
			return MapNull, ErrMap(text[n])
		}
		infoMap = append(infoMap, intVar)
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
			return ColisNull, ColisPosNull, ErrMap(text[n])
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
		return ColisNull, ColisPosNull, ErrColor(lower)
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
			return ColisNull, ColisPosNull, ErrMap(text[n])
		}
		infoPosColis = append(infoPosColis, intVar)
	}
	return infoColis, infoPosColis, nil
}

func findObject(words []string, warehouse *WarehouseSquareGraph) error {
	switch size := len(words); size {
	case 4:
		name, posColor, err := infoColis(words)
		if err != nil {
			return err
		}
		warehouse.AddPackage(name, posColor[0], posColor[1], posColor[2])
	case 3:
		name, pos, err := infoPaletteCamion(words)
		if err != nil {
			return err
		}
		warehouse.AddTransporter(name, pos[0], pos[1])
	case 5:
		name, posSizeRound, err := infoPaletteCamion(words)
		if err != nil {
			return err
		}
		warehouse.AddTruck(name, posSizeRound[0], posSizeRound[1], posSizeRound[2], posSizeRound[3])
	}
	return nil
}

func divide(text string) ([]string, error) {
	words := strings.Fields(text)
	if len(words) < 3 || len(words) > 5 {
		response := []string{}
		return response, ErrInput(words)
	}
	return words, nil
}

// AnalyzeAllText use the array string return by readMap to create the graph WarehouseSquareGraph and the maximum number of turn
func AnalyzeAllText() (warehouseGraph *WarehouseSquareGraph, round int, err error) {
	text := readMap()

	firstLine, err := divide(text[0])
	if err != nil {
		warehouseGraphErr := createWarehouseGraph(0, 0)
		return warehouseGraphErr, 0, err
	}
	mapInfo, errMap := infoMap(firstLine)
	if errMap != nil {
		warehouseGraphErr := createWarehouseGraph(0, 0)
		return warehouseGraphErr, 0, errMap
	}
	warehouseGraph = createWarehouseGraph(mapInfo[0], mapInfo[1])
	for n := 1; n < len(text); n++ {
		line, err := divide(text[n])
		if err != nil {
			warehouseGraphErr := createWarehouseGraph(0, 0)
			return warehouseGraphErr, 0, err
		}
		errLine := findObject(line, warehouseGraph)
		if errLine != nil {
			warehouseGraphErr := createWarehouseGraph(0, 0)
			return warehouseGraphErr, 0, errLine
		}
	}
	return warehouseGraph, mapInfo[2], nil
}
