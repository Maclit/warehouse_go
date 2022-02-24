package main

import "fmt"

func (graph *WarehouseSquareGraph) unloadTransporter(transporterPoint Point) {
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

func (graph *WarehouseSquareGraph) moveTransporterToNextPosition(start, end Point) {
	newX := end.x
	newY := end.y
	oldX := start.x
	oldY := start.y
	graph.nodes[newX+(newY*graph.height)].transporter = graph.nodes[oldX+(oldY*graph.height)].transporter
	graph.nodes[oldX+(oldY*graph.height)].transporter = nil
	fmt.Printf(" GO [%d,%d]\n", newX, newY)
}

func (graph *WarehouseSquareGraph) loadTransporter(transporter, box Point) {
	boxX := box.x
	boxY := box.y
	transporterX := transporter.x
	transporterY := transporter.y
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
}
