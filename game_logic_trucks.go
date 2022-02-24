package main

import "fmt"

func updateTruckStatus(graph *WarehouseSquareGraph, truckNode Node) {
	fmt.Print(truckNode.truck.name)
	x := truckNode.point.x
	y := truckNode.point.y
	if graph.nodes[x+(y*graph.height)].truck.currentTimer == graph.nodes[x+(y*graph.height)].truck.maxTimer {
		graph.nodes[x+(y*graph.height)].truck.isGone = true
	}
	if graph.nodes[x+(y*graph.height)].truck.isGone {
		graph.nodes[x+(y*graph.height)].truck.currentTimer--
		fmt.Printf(" GONE %d/%d\n", graph.nodes[x+(y*graph.height)].truck.currentLoad, graph.nodes[x+(y*graph.height)].truck.maxLoad)
		if graph.nodes[x+(y*graph.height)].truck.currentTimer == 0 {
			graph.nodes[x+(y*graph.height)].truck.isGone = false
			graph.nodes[x+(y*graph.height)].truck.currentLoad = 0
		}
		return
	}
	if isGameFinished(graph) {
		graph.nodes[x+(y*graph.height)].truck.currentTimer = graph.nodes[x+(y*graph.height)].truck.maxTimer
	}
	fmt.Printf(" WAITING %d/%d\n", graph.nodes[x+(y*graph.height)].truck.currentLoad, graph.nodes[x+(y*graph.height)].truck.maxLoad)
	if graph.nodes[x+(y*graph.height)].truck.currentLoad == graph.nodes[x+(y*graph.height)].truck.maxLoad {
		graph.nodes[x+(y*graph.height)].truck.currentTimer = graph.nodes[x+(y*graph.height)].truck.maxTimer
	}
}
