package main

import "fmt"

func (graph *WarehouseSquareGraph) moveTransporterTowardNearestBox(startNode Node) {
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

func (graph *WarehouseSquareGraph) moveTransporterTowardNearestTruck(startNode Node) {
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

func (graph *WarehouseSquareGraph) shortestPath(start Node, end Node, path []Node) []Node {
	path = append(path, start)
	if graph.isNodeInArray(path, end) {
		return path
	}
	shortest := make([]Node, 0)
	neighbors := getAllNeighborsNode(graph, start)
	for _, node := range neighbors {
		if !graph.isNodeInArray(path, node) {
			newPath := graph.shortestPath(node, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func (graph *WarehouseSquareGraph) findClosestObject(startNode Node, toFind int) Node {
	toVisit := getAllNeighborsNode(graph, startNode)
	visited := []Node{}
	for len(toVisit) > 0 {
		nextNode := toVisit[0]
		toVisit = toVisit[1:]
		visited = append(visited, nextNode)
		if (toFind == TRUCK && nextNode.truck != nil) || (toFind == BOX && nextNode.box != nil) {
			return nextNode
		}
		neightborsList := getAllNeighborsNode(graph, nextNode)
		for _, node := range neightborsList {
			if !graph.isNodeInArray(visited, node) {
				toVisit = append(toVisit, node)
			}
		}
	}
	return Node{}
}

func (graph *WarehouseSquareGraph) updateTruckStatus(truckNode Node) {
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
	if graph.isGameFinished() {
		graph.nodes[x+(y*graph.height)].truck.currentTimer = graph.nodes[x+(y*graph.height)].truck.maxTimer
	}
	fmt.Printf(" WAITING %d/%d\n", graph.nodes[x+(y*graph.height)].truck.currentLoad, graph.nodes[x+(y*graph.height)].truck.maxLoad)
	if graph.nodes[x+(y*graph.height)].truck.currentLoad == graph.nodes[x+(y*graph.height)].truck.maxLoad {
		graph.nodes[x+(y*graph.height)].truck.currentTimer = graph.nodes[x+(y*graph.height)].truck.maxTimer
	}
}
