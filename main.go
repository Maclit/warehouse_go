package main

func main() {
	warehouseGraph, nbTurn, err := analyzeAllText()
	if err != nil {
		return
	}
	play(nbTurn, warehouseGraph)
}
