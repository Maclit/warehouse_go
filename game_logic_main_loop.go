package main

import (
	"fmt"
	"strconv"
)

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

func play(turns int, graph *WarehouseSquareGraph) {
	for i := 0; i < turns; i++ {
		fmt.Printf("tour %d\n", i+1)
		transporterNodeList := graph.getWharehouseNodeListWithObject(TRANSPORTER)
		for _, transporterNode := range transporterNodeList {
			if transporterNode.transporter.isLoaded {
				err := graph.moveTransporterTowardNearestTruck(transporterNode)
				if err != nil {
					fmt.Println(err)
					emoji, emojiErr := strconv.Unquote(`"😱"`)
					if emojiErr == nil {
						fmt.Println("\n", emoji)
					}
					return
				}
			} else {
				err := graph.moveTransporterTowardNearestBox(transporterNode)
				if err != nil {
					fmt.Println(err)
					emoji, emojiErr := strconv.Unquote(`"😱"`)
					if emojiErr == nil {
						fmt.Println("\n", emoji)
					}
					return
				}
			}
		}
		truckNodeList := graph.getWharehouseNodeListWithObject(TRUCK)
		for _, truckNode := range truckNodeList {
			err := graph.updateTruckStatus(truckNode)
			if err != nil {
				fmt.Println(err)
				emoji, emojiErr := strconv.Unquote(`"😱"`)
				if emojiErr == nil {
					fmt.Println("\n", emoji)
				}
				return
			}
		}
		if isGameFinished(graph) && graph.areAllTrucksGone() {
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
