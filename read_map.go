package main

import (
	"bufio"
	"log"
	"os"
)

func readMap() (lines []string, err error) {
	var line string
	if len(os.Args) <= 1 || len(os.Args) > 2 {
		return []string{}, ArgumentsError("Incorrect arguments.")
	}
	mapfile := os.Args[1]
	f, err := os.Open(mapfile)
	if err != nil {
		return []string{}, ArgumentsError("Wrong file name.")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
	}
	f.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines, nil
}
