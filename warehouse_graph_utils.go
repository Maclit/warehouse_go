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

func moveTransporterToNextPosition(graph *WarehouseSquareGraph, start, end Point) {
	newX := end.x
	newY := end.y
	oldX := start.x
	oldY := start.y
	graph.nodes[newX+(newY*graph.height)].transporter = graph.nodes[oldX+(oldY*graph.height)].transporter
	graph.nodes[oldX+(oldY*graph.height)].transporter = nil
	fmt.Printf(" GO [%d,%d]\n", newX, newY)
}

func isNodeInArray(array []Node, node Node) bool {
	for _, element := range array {
		if element.point.x == node.point.x && element.point.y == node.point.y {
			return true
		}
	}
	return false
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

func (graph *WarehouseSquareGraph) getWharehouseNodeListWithObject(object int) []Node {
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

func (graph *WarehouseSquareGraph) isEmpty() bool {
	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.box != nil {
			return false
		}
	}
	return true
}

func (graph *WarehouseSquareGraph) areAllTrucksGone() bool {
	status := true

	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.truck != nil && !node.truck.isGone {
			status = false
		}
	}
	return status
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

func (graph *WarehouseSquareGraph) validate() error {
	if len(graph.getWharehouseNodeListWithObject(TRUCK)) == 0 {
		return GraphError("No trucks")
	}
	if len(graph.getWharehouseNodeListWithObject(TRANSPORTER)) == 0 {
		return GraphError("No transporters")
	}
	return nil
}
