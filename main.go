package main

func main() {
	//width := 5
	//height := 5

	text1 := "5 5 1000"
	text2 := "box_1 1 1 green"
	text3 := "transporter_1 2 0"
	text4 := "truck 3 4 10 50"

	words_1, err := divide(text1)
	if err != nil {
		return
	}
	words_2, err := divide(text2)
	if err != nil {
		return
	}
	words_3, err := divide(text3)
	if err != nil {
		return
	}
	words_4, err := divide(text4)
	if err != nil {
		return
	}
	size, err := info_map(words_1)
	if err != nil {
		return
	}
	warehouse_graph := WarehouseSquareGraph{width: size[0], height: size[1]}
	warehouse_graph.initializeNodes()
	err2 := find_object(words_2, &warehouse_graph)
	if err2 != nil {
		return
	}
	err3 := find_object(words_3, &warehouse_graph)
	if err3 != nil {
		return
	}
	err4 := find_object(words_4, &warehouse_graph)
	if err4 != nil {
		return
	}
	/*warehouse_graph.addPackage("box_1", 1, 1, GREEN)
	warehouse_graph.addTransporter("transporter_1", 2, 0)
	warehouse_graph.addTruck("truck_1", 3, 4, 10, 50)*/
	warehouse_graph.print()
}
