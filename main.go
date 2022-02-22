package main

func main() {
	warehouseGraph, nbTurn, err := AnalyzeAllText()
	if err != nil {
		return
	}
	play(nbTurn, warehouseGraph)
}
