package main

func main() {
	text1 := "5 5 30"
	text2 := "box_1 4 3 green"
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

	warehouse_graph := createWarehouseGraph(map_info[0], map_info[1])

	words_2, err := divide(text2)
	if err != nil {
		return
	}
	err2 := find_object(words_2, warehouse_graph)
	if err2 != nil {
		return
	}
	words_3, err := divide(text3)
	if err != nil {
		return
	}
	err3 := find_object(words_3, warehouse_graph)
	if err3 != nil {
		return
	}
	words_4, err := divide(text4)
	if err != nil {
		return
	}
	err4 := find_object(words_4, warehouse_graph)
	if err4 != nil {
		return
	}
	play(map_info[2], warehouse_graph)
}
