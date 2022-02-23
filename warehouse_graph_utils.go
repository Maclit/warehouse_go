package main

import "fmt"

func unloadTransporter(graph *WarehouseSquareGraph, transporterPoint Point) {
	x := transporterPoint.x
	y := transporterPoint.y
	if graph.nodes[x+(y*graph.height)].truck.isGone {
		fmt.Print(" WAIT\n")
		return
	}
	newWeight := graph.nodes[x+(y*graph.height)].truck.currentLoad + graph.nodes[x+(y*graph.height)].transporter.box.color
	if newWeight > graph.nodes[x+(y*graph.height)].truck.maxLoad {
		fmt.Print(" WAIT\n")
		graph.nodes[x+(y*graph.height)].truck.isGone = true
		graph.nodes[x+(y*graph.height)].truck.currentTimer = graph.nodes[x+(y*graph.height)].truck.maxTimer
		return
	}
	graph.nodes[x+(y*graph.height)].truck.currentLoad = newWeight
	fmt.Printf(" LEAVE %s ", graph.nodes[x+(y*graph.height)].transporter.box.name)
	switch graph.nodes[x+(y*graph.height)].transporter.box.color {
	case YELLOW:
		fmt.Print("YELLOW\n")
	case GREEN:
		fmt.Print("GREEN\n")
	case BLUE:
		fmt.Print("BLUE\n")
	}
	graph.nodes[x+(y*graph.height)].transporter.box = nil
	graph.nodes[x+(y*graph.height)].transporter.isLoaded = false
}

func (graph *WarehouseSquareGraph) doesNodeHasObject(coord Point, object int) bool {
	x := coord.x
	y := coord.y
	if object == TRUCK && graph.nodes[x+(y*graph.height)].truck != nil {
		return true
	}
	if object == BOX && graph.nodes[x+(y*graph.height)].box != nil {
		return true
	}
	if object == TRANSPORTER && graph.nodes[x+(y*graph.height)].transporter != nil {
		return true
	}
	return false
}

func moveTransporterToNextPosition(graph *WarehouseSquareGraph, start, end Point) {
	newX := end.x
	newY := end.y
	oldX := start.x
	oldY := start.y
	graph.nodes[newX+(newY*graph.height)].transporter = graph.nodes[oldX+(oldY*graph.height)].transporter
	graph.nodes[oldX+(oldY*graph.height)].transporter = nil
	fmt.Printf(" GO [%d,%d]\n", newX, newY)
}

func getEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	var x, y int

	if node.point.x >= 0 && node.point.x < graph.width-1 {
		x = node.point.x + 1
		y = node.point.y
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			return graph.nodes[x+(y*graph.height)], nil
		}
	}
	if node.point.x > 0 && node.point.x <= graph.width-1 {
		x = node.point.x - 1
		y = node.point.y
		return graph.nodes[x+(y*graph.height)], nil
	}
	if node.point.y >= 0 && node.point.y < graph.height-1 {
		x = node.point.x
		y = node.point.y + 1
		return graph.nodes[x+(y*graph.height)], nil
	}
	if node.point.y > 0 && node.point.y <= graph.height-1 {
		x = node.point.x
		y = node.point.y - 1
		return graph.nodes[x+(y*graph.height)], nil
	}
	return Node{}, StuckTransporterError(node.transporter.name)
}

func getAllNeighborsNode(graph *WarehouseSquareGraph, node Node) []Node {
	var x, y int

	neighbors := []Node{}
	if node.point.x >= 0 && node.point.x < graph.width-1 {
		x = node.point.x + 1
		y = node.point.y
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			neighbors = append(neighbors, graph.nodes[x+(y*graph.height)])
		}
	}
	if node.point.x > 0 && node.point.x <= graph.width-1 {
		x = node.point.x - 1
		y = node.point.y
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			neighbors = append(neighbors, graph.nodes[x+(y*graph.height)])
		}
	}
	if node.point.y >= 0 && node.point.y < graph.height-1 {
		x = node.point.x
		y = node.point.y + 1
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			neighbors = append(neighbors, graph.nodes[x+(y*graph.height)])
		}
	}
	if node.point.y > 0 && node.point.y <= graph.height-1 {
		x = node.point.x
		y = node.point.y - 1
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			neighbors = append(neighbors, graph.nodes[x+(y*graph.height)])
		}
	}
	return neighbors
}

func isNodeInArray(array []Node, node Node) bool {
	for _, element := range array {
		if element.point.x == node.point.x && element.point.y == node.point.y {
			return true
		}
	}
	return false
}

func isWarehouseEmpty(graph *WarehouseSquareGraph) bool {
	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.box != nil {
			return false
		}
	}
	return true
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

func isTruckGone(graph *WarehouseSquareGraph) bool {
	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.truck != nil && !node.truck.isGone {
			return false
		}
	}
	return true
}
