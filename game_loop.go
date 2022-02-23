package main

import (
	"fmt"
	"strconv"
)

func play(turns int, graph *WarehouseSquareGraph) {
	if isWarehouseEmpty(graph) {
		emoji, err := strconv.Unquote(`"🚒"`)
		if err == nil {
			fmt.Println("\n", emoji)
		}
		return
	}
	for i := 0; i < turns; i++ {
		fmt.Printf("tour %d\n", i+1)
		transporterNodeList := getWharehouseNodeListWithObject(graph, TRANSPORTER)
		for _, transporterNode := range transporterNodeList {
			if transporterNode.transporter.isLoaded {
				moveTransporterTowardNearestTruck(graph, transporterNode)
			} else {
				moveTransporterTowardNearestBox(graph, transporterNode)
			}
		}
		truckNodeList := getWharehouseNodeListWithObject(graph, TRUCK)
		for _, truckNode := range truckNodeList {
			updateTruckStatus(graph, truckNode)
		}
		if isGameFinished(graph) && isTruckGone(graph) {
			emoji, err := strconv.Unquote(`"🚒"`)
			if err == nil {
				fmt.Println("\n", emoji)
			}
			return
		}
		fmt.Println()
	}
}

func getWharehouseNodeListWithObject(graph *WarehouseSquareGraph, object int) []Node {
	objectList := []Node{}

	for x := 0; x < graph.width; x++ {
		for y := 0; y < graph.height; y++ {
			if object == TRANSPORTER && graph.nodes[x+(y*graph.height)].transporter != nil {
				objectList = append(objectList, graph.nodes[x+(y*graph.height)])
			}
			if object == TRUCK && graph.nodes[x+(y*graph.height)].truck != nil {
				objectList = append(objectList, graph.nodes[x+(y*graph.height)])
			}
			if object == BOX && graph.nodes[x+(y*graph.height)].box != nil {
				objectList = append(objectList, graph.nodes[x+(y*graph.height)])
			}
		}
	}
	return objectList
}
