package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getMap() string {
	fmt.Print("Enter map file path: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

func readMap() []string {
	lines := []string{}
	var line string
	mapfile := getMap()
	f, err := os.Open(mapfile)
	if err != nil {
		log.Fatal(err)
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
	return lines
}
