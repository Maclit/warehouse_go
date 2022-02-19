package main

import "fmt"

func play(map_info []int, warehouse *WarehouseSquareGraph) {
	warehouse.print()
	for i := 0; i < map_info[2]; i++ {
		transporter_list := warehouse.getTransporterList()
		for _, transporter := range transporter_list {
			if transporter.is_loaded {
				fmt.Println("Moving to truck.")
			} else {
				fmt.Println("Moving to nearest box.")
				warehouse.moveTransporterToNearestBox(transporter.name)
			}
		}
	}
	warehouse.print()
}

func main() {
	text1 := "5 5 1"
	text2 := "box_1 1 1 green"
	text3 := "transporter_1 2 0"
	text4 := "truck 3 4 10 50"

	words_1, err := divide(text1)
	if err != nil {
		return
	}
	map_info, err := info_map(words_1)
	if err != nil {
		return
	}
	warehouse_graph := WarehouseSquareGraph{width: map_info[0], height: map_info[1]}
	warehouse_graph.initializeNodes()

	words_2, err := divide(text2)
	if err != nil {
		return
	}
	err2 := find_object(words_2, &warehouse_graph)
	if err2 != nil {
		return
	}
	words_3, err := divide(text3)
	if err != nil {
		return
	}
	err3 := find_object(words_3, &warehouse_graph)
	if err3 != nil {
		return
	}
	words_4, err := divide(text4)
	if err != nil {
		return
	}
	err4 := find_object(words_4, &warehouse_graph)
	if err4 != nil {
		return
	}
	play(map_info, &warehouse_graph)
}
