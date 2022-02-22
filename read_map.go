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
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)
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

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)

	}

	println(line)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
