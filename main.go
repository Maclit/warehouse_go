package main

func main() {
	warehouse_graph, nb_turn, err := AnalyzeAllText()
	if err != nil {
		return
	}
	play(nb_turn, warehouse_graph)
}
