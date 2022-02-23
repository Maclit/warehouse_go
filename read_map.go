package main

import (
	"bufio"
	"log"
	"os"
	// "strings"
)

// func getMap() string {
// 	fmt.Print("Enter map file path: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	input = strings.ReplaceAll(input, "\n", "")
// 	input = strings.ReplaceAll(input, "\r", "")
// 	return input
// }

func readMap() (lines []string, err error) {
	// lines := []string{}
	var line string
	if len(os.Args) <= 1 {
		argsNull := []string{}
		return argsNull, ArgumentsError(os.Args[0])
	}
	mapfile := os.Args[1]
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
	return lines, nil
}
