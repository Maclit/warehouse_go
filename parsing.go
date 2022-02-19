package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrInput []string

func (e ErrInput) Error() string {
	return fmt.Sprintf("Can not accept more than 5 parameters and less than 3. number of parameters: %d", len(e))
}

type ErrMap string

func (e ErrMap) Error() string {
	return fmt.Sprintf("The input for the map is not a number or not greater than zero, string: %s ", e)
}

type ErrColor string

func (e ErrColor) Error() string {
	return fmt.Sprintf("This input does not correspond to an color accepted by the program, string: %s ", e)
}

func info_map(text []string) ([]int, error) {
	infoMap := []int{}

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

func info_colis(text []string) (string, []int, error) {
	infoColis := []int{}
	infoNameColis := text[0]

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

func info_palette_camion(text []string) (string, []int, error) {
	infoPosColis := []int{}
	infoColis := text[0]

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

func find_object(words []string, warehouse *WarehouseSquareGraph) error {
	switch size := len(words); size {
	case 4:
		name, pos_color, err := info_colis(words)
		if err != nil {
			return err
		}
		warehouse.AddPackage(name, pos_color[0], pos_color[1], pos_color[2])
	case 3:
		name, pos, err := info_palette_camion(words)
		if err != nil {
			return err
		}
		warehouse.AddTransporter(name, pos[0], pos[1])
	case 5:
		name, pos_size_round, err := info_palette_camion(words)
		if err != nil {
			return err
		}
		warehouse.AddTruck(name, pos_size_round[0], pos_size_round[1], pos_size_round[2], pos_size_round[3])
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
