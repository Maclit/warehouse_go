package main

func play(map_info []int, warehouse *WarehouseSquareGraph) {
	warehouse.Print()
	for i := 0; i < map_info[2]; i++ {
		/*
			transporter_list := warehouse.GetTransporterList()
			for _, transporter := range transporter_list {
				if transporter.is_loaded {
					fmt.Println("Moving to truck.")
				} else {
					fmt.Println("Moving to nearest box.")
					warehouse.MoveTransporterToNearestBox(transporter.name)
				}
			}*/
	}
	warehouse.Print()
}

func getWharehouseTransporterList(graph *WarehouseSquareGraph) []Transporter {
	transporter_list := []Transporter{}

	for x := 0; x < graph.width; x++ {
		for y := 0; y < graph.height; y++ {
			if graph.nodes[x+(y*graph.height)].transporter != nil {
				transporter_list = append(transporter_list, *graph.nodes[x+(y*graph.height)].transporter)
			}
		}
	}
	return transporter_list
}
