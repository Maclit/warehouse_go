package main

import (
	"fmt"
	"strconv"
)

func main() {
	warehouseGraph, nbTurn, fileErr := analyzeAllText()
	if fileErr != nil {
		fmt.Println(fileErr)
		emoji, emojiErr := strconv.Unquote(`"😱"`)
		if emojiErr == nil {
			fmt.Println("\n", emoji)
		}
		return
	}
	graphErr := warehouseGraph.validate()
	if graphErr != nil {
		fmt.Println(graphErr)
		emoji, emojiErr := strconv.Unquote(`"😱"`)
		if emojiErr == nil {
			fmt.Println("\n", emoji)
		}
		return
	}
	play(nbTurn, warehouseGraph)
}
