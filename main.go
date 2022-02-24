package main

import (
	"fmt"
	"strconv"
)

func printErrorEmoji() {
	emoji, emojiErr := strconv.Unquote(`"😱"`)
	if emojiErr == nil {
		fmt.Println("\n", emoji)
	}
}

func main() {
	warehouseGraph, nbTurn, fileErr := analyzeAllText()
	if fileErr != nil {
		fmt.Println(fileErr)
		printErrorEmoji()
		return
	}
	graphErr := warehouseGraph.Validate()
	if graphErr != nil {
		fmt.Println(graphErr)
		printErrorEmoji()
		return
	}
	playErr := play(nbTurn, warehouseGraph)
	if playErr != nil {
		fmt.Println(playErr)
		printErrorEmoji()
	}
}
