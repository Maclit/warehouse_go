package main

import (
	"fmt"
	"strconv"
)

func play(turns int, graph *WarehouseSquareGraph) {
	if graph.isEmpty() {
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
		if isGameFinished(graph) && graph.areAllTrucksGone() {
			emoji, err := strconv.Unquote(`"🚒"`)
			if err == nil {
				fmt.Println("\n", emoji)
			}
			return
		}
		fmt.Println()
	}
}

func main() {
	warehouseGraph, nbTurn, err := analyzeAllText()
	if err != nil {
		return
	}
	play(nbTurn, warehouseGraph)
}
