package main

import "fmt"

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, startNode Node) {
	fmt.Print(startNode.transporter.name)
	if graph.isEmpty() {
		if graph.doesNodeHasObject(startNode.point, TRUCK) {
			neighbor, err := getEmptyNeighbor(graph, startNode)
			if err != nil {
				fmt.Println(err.Error())
			}
			graph.moveTransporterToNextPosition(startNode.point, neighbor.point)
		} else {
			fmt.Print(" WAIT\n")
		}
		return
	}
	closestBox := graph.findClosestObject(startNode, BOX)
	if closestBox.box == nil {
		return
	}
	shortestPath := graph.shortestPath(startNode, closestBox, make([]Node, 0))
	if len(shortestPath) == 2 {
		boxX := shortestPath[1].point.x
		boxY := shortestPath[1].point.y
		transporterX := startNode.point.x
		transporterY := startNode.point.y
		graph.nodes[transporterX+(transporterY*graph.height)].transporter.isLoaded = true
		graph.nodes[transporterX+(transporterY*graph.height)].transporter.box = graph.nodes[boxX+(boxY*graph.height)].box
		fmt.Printf(" TAKE %s ", graph.nodes[boxX+(boxY*graph.height)].box.name)
		switch graph.nodes[boxX+(boxY*graph.height)].box.color {
		case YELLOW:
			fmt.Print("YELLOW\n")
		case GREEN:
			fmt.Print("GREEN\n")
		case BLUE:
			fmt.Print("BLUE\n")
		}
		graph.nodes[boxX+(boxY*graph.height)].box = nil
	} else {
		graph.moveTransporterToNextPosition(startNode.point, shortestPath[1].point)
	}
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, startNode Node) {
	fmt.Print(startNode.transporter.name)
	if graph.doesNodeHasObject(startNode.point, TRUCK) {
		graph.unloadTransporter(startNode.point)
	} else {
		closestBox := graph.findClosestObject(startNode, TRUCK)
		if closestBox.truck == nil {
			fmt.Print(" WAIT\n")
			return
		}
		shortestPath := graph.shortestPath(startNode, closestBox, make([]Node, 0))
		graph.moveTransporterToNextPosition(startNode.point, shortestPath[1].point)
	}
}

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
