package main

import "fmt"

func play(turns int, warehouse *WarehouseSquareGraph) {
	if isWarehouseEmpty(warehouse) {
		return // TODO finish
	}
	for i := 0; i < turns; i++ {
		fmt.Printf("tour %d\n", i+1)
		transporterNodeList := getWharehouseTransporterNodeList(warehouse)
		for _, node := range transporterNodeList {
			if node.transporter.is_loaded {
				moveTransporterTowardNearestTruck(warehouse, node)
			} else {
				moveTransporterTowardNearestBox(warehouse, node)
			}
		}
		if isGameFinished(warehouse) && isTruckGone(warehouse) {
			return // TODO finish
		}
		warehouse.Print()
		fmt.Println()
	}
}

func getWharehouseTransporterNodeList(graph *WarehouseSquareGraph) []Node {
	transporterNodeList := []Node{}

	for x := 0; x < graph.width; x++ {
		for y := 0; y < graph.height; y++ {
			if graph.nodes[x+(y*graph.height)].transporter != nil {
				transporterNodeList = append(transporterNodeList, graph.nodes[x+(y*graph.height)])
			}
		}
	}
	return transporterNodeList
}
