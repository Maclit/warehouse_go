package main

import . "fmt"

const (
	GREEN  int = 0
	YELLOW int = 1
	BLUE   int = 2
)

type Box struct {
	name  string
	color int
	x     int
	y     int
}
type Transporter struct {
	name string
	x    int
	y    int
}

type Truck struct {
	name         string
	x            int
	y            int
	max_load     int
	waiting_time int
}

type Node struct {
	box         Box
	transporter Transporter
	truck       Truck
	x           int
	y           int
}

type WarehouseSquareGraph struct {
	nodes  []Node
	width  int
	height int
}

func (graph *WarehouseSquareGraph) initializeNodes() {
	for i := 0; i < graph.width*graph.height; i++ {
		graph.nodes[i] = Node{x: i % graph.width, y: i / graph.height}
	}
}

func (graph *WarehouseSquareGraph) addPackage(name string, x, y, color int) {
	graph.nodes[x+(y*graph.height)].box.name = name
	graph.nodes[x+(y*graph.height)].box.x = x
	graph.nodes[x+(y*graph.height)].box.y = y
	graph.nodes[x+(y*graph.height)].box.color = color
}

func (graph *WarehouseSquareGraph) addTruck(name string, x, y, max_load, waiting_time int) {
	graph.nodes[x+(y*graph.height)].truck.name = name
	graph.nodes[x+(y*graph.height)].truck.x = x
	graph.nodes[x+(y*graph.height)].truck.y = y
	graph.nodes[x+(y*graph.height)].truck.max_load = max_load
	graph.nodes[x+(y*graph.height)].truck.waiting_time = waiting_time
}

func (graph *WarehouseSquareGraph) addTransporter(name string, x, y int) {
	graph.nodes[x+(y*graph.height)].transporter.name = name
	graph.nodes[x+(y*graph.height)].transporter.x = x
	graph.nodes[x+(y*graph.height)].transporter.y = y
}

func (graph *WarehouseSquareGraph) print() {
	for i := 0; i < graph.width*graph.height; i++ {
		if i != 0 && i%graph.height == 0 {
			Print("\n")
		}
		node := graph.nodes[i]
		if node.box.name != "" {
			Print("[x]")
		} else if node.transporter.name != "" {
			Print("[t]")
		} else if node.truck.name != "" {
			Print("[g]")
		} else {
			Print("[ ]")
		}
	}
}

func main() {
	width := 5
	height := 5
	length := width * height
	nodes := make([]Node, length)

	warehouse_graph := WarehouseSquareGraph{
		nodes:  nodes,
		width:  width,
		height: height,
	}
	warehouse_graph.initializeNodes()
	warehouse_graph.addPackage("box_1", 1, 1, GREEN)
	warehouse_graph.addTransporter("transporter_1", 2, 0)
	warehouse_graph.addTruck("truck_1", 3, 4, 10, 50)
	warehouse_graph.print()
}
