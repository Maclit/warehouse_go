package main

import "fmt"

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, startNode Node) {
	fmt.Print(startNode.transporter.name)
	if isWarehouseEmpty(graph) {
		if graph.doesNodeHasObject(startNode.point, TRUCK) {
			neighbor, err := getEmptyNeighbor(graph, startNode)
			if err != nil {
				fmt.Println(err.Error())
			}
			moveTransporterToNextPosition(graph, startNode.point, neighbor.point)
		} else {
			fmt.Print(" WAIT\n")
		}
		return
	}
	closest_box := findClosestObject(graph, startNode, BOX)
	if closest_box.box == nil {
		return
	}
	shortestPath := shortestPath(graph, startNode, closest_box, make([]Node, 0))
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
		moveTransporterToNextPosition(graph, startNode.point, shortestPath[1].point)
	}
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, startNode Node) {
	fmt.Print(startNode.transporter.name)
	if graph.doesNodeHasObject(startNode.point, TRUCK) {
		unloadTransporter(graph, startNode.point)
	} else {
		closestBox := findClosestObject(graph, startNode, TRUCK)
		if closestBox.truck == nil {
			fmt.Print(" WAIT\n")
			return
		}
		shortestPath := shortestPath(graph, startNode, closestBox, make([]Node, 0))
		moveTransporterToNextPosition(graph, startNode.point, shortestPath[1].point)
	}
}

func shortestPath(graph *WarehouseSquareGraph, start Node, end Node, path []Node) []Node {
	path = append(path, start)
	if isNodeInArray(path, end) {
		return path
	}
	shortest := make([]Node, 0)
	neighbors := getAllNeighborsNode(graph, start)
	for _, node := range neighbors {
		if !isNodeInArray(path, node) {
			newPath := shortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func findClosestObject(graph *WarehouseSquareGraph, startNode Node, toFind int) Node {
	if toFind == TRUCK && startNode.truck != nil {
		return startNode
	}
	if toFind == BOX && startNode.box != nil {
		return startNode
	}
	toVisit := getAllNeighborsNode(graph, startNode)
	visited := []Node{}
	for {
		if len(toVisit) == 0 {
			return Node{}
		}
		nextNode := toVisit[0]
		toVisit = toVisit[1:]
		visited = append(visited, nextNode)
		if toFind == TRUCK && nextNode.truck != nil {
			return nextNode
		}
		if toFind == BOX && nextNode.box != nil {
			return nextNode
		}
		neightborsList := getAllNeighborsNode(graph, nextNode)
		if len(neightborsList) > 0 {
			for _, node := range neightborsList {
				if !isNodeInArray(visited, node) {
					toVisit = append(toVisit, node)
				}
			}
		}
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
