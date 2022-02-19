package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func get_map() string {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Replace(input, "\r\n", "", -1)
	return input
}

func read_map() []string {

	lines := []string{}
	var line string
	mapfile := get_map()
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
