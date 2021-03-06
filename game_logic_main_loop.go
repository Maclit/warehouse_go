package main

import (
	"fmt"
	"strconv"
)

func printSuccessEmoji() {
	emoji, err := strconv.Unquote(`"😎"`)
	if err == nil {
		fmt.Println("\n", emoji)
	}
}

func printEndEmoji() {
	emoji, err := strconv.Unquote(`"🙂"`)
	if err == nil {
		fmt.Println("\n", emoji)
	}
}

func isGameFinished(graph *WarehouseSquareGraph) bool {
	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.box != nil {
			return false
		}
		if node.transporter != nil && node.transporter.isLoaded {
			return false
		}
	}
	return true
}

func processTransporters(graph *WarehouseSquareGraph) error {
	transporterNodeList := graph.GetWharehouseNodeListWithObject(TRANSPORTER)
	for _, transporterNode := range transporterNodeList {
		if transporterNode.transporter.isLoaded {
			err := graph.MoveTransporterTowardNearestTruck(transporterNode)
			if err != nil {
				return err
			}
		} else {
			err := graph.MoveTransporterTowardNearestBox(transporterNode)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func processTrucks(graph *WarehouseSquareGraph) error {
	truckNodeList := graph.GetWharehouseNodeListWithObject(TRUCK)
	for _, truckNode := range truckNodeList {
		err := graph.UpdateTruckStatus(truckNode)
		if err != nil {
			return err
		}
	}
	return nil
}

func play(turns int, graph *WarehouseSquareGraph) error {
	for i := 0; i < turns; i++ {
		fmt.Printf("tour %d\n", i+1)
		transporterErr := processTransporters(graph)
		if transporterErr != nil {
			return transporterErr
		}
		truckErr := processTrucks(graph)
		if truckErr != nil {
			return truckErr
		}
		if isGameFinished(graph) && graph.AreAllTrucksGone() {
			printSuccessEmoji()
			return nil
		}
		fmt.Println()
	}
	printEndEmoji()
	return nil
}
