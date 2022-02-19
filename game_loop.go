package main

import "fmt"

func play(turns int, warehouse *WarehouseSquareGraph) {
	warehouse.Print()
	for i := 0; i < turns; i++ {
		transporter_node_list := getWharehouseTransporterNodeList(warehouse)
		for _, node := range transporter_node_list {
			if node.transporter.is_loaded {
				fmt.Println("Moving to truck.")
			} else {
				fmt.Println("Moving to nearest box.")
				moveTransporterTowardNearestBox(warehouse, node.x, node.y)
			}
		}
	}
	warehouse.Print()
}

func getWharehouseTransporterNodeList(graph *WarehouseSquareGraph) []Node {
	transporter_node_list := []Node{}

	for x := 0; x < graph.width; x++ {
		for y := 0; y < graph.height; y++ {
			if graph.nodes[x+(y*graph.height)].transporter != nil {
				transporter_node_list = append(transporter_node_list, graph.nodes[x+(y*graph.height)])
			}
		}
	}
	return transporter_node_list
}
