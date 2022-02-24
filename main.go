package main

import (
	"fmt"
	"strconv"
)

func play(turns int, graph *WarehouseSquareGraph) {
	for i := 0; i < turns; i++ {
		fmt.Printf("tour %d\n", i+1)
		transporterNodeList := graph.getWharehouseNodeListWithObject(TRANSPORTER)
		for _, transporterNode := range transporterNodeList {
			if transporterNode.transporter.isLoaded {
				graph.moveTransporterTowardNearestTruck(transporterNode)
			} else {
				graph.moveTransporterTowardNearestBox(transporterNode)
			}
		}
		truckNodeList := graph.getWharehouseNodeListWithObject(TRUCK)
		for _, truckNode := range truckNodeList {
			graph.updateTruckStatus(truckNode)
		}
		if graph.isGameFinished() && graph.areAllTrucksGone() {
			emoji, err := strconv.Unquote(`"😎"`)
			if err == nil {
				fmt.Println("\n", emoji)
			}
			return
		}
		fmt.Println()
	}
	emoji, err := strconv.Unquote(`"🙂"`)
	if err == nil {
		fmt.Println("\n", emoji)
	}
}

func main() {
	warehouseGraph, nbTurn, fileErr := analyzeAllText()
	if fileErr != nil {
		fmt.Println(fileErr)
		emoji, emojiErr := strconv.Unquote(`"😱"`)
		if emojiErr == nil {
			fmt.Println("\n", emoji)
		}
		return
	}
	graphErr := warehouseGraph.validate()
	if graphErr != nil {
		fmt.Println(graphErr)
		emoji, emojiErr := strconv.Unquote(`"😱"`)
		if emojiErr == nil {
			fmt.Println("\n", emoji)
		}
		return
	}
	play(nbTurn, warehouseGraph)
}
