package main

import "fmt"

const (
	GREEN  int = 0
	YELLOW int = 1
	BLUE   int = 2
)

type Box struct {
	name  string
	color int
}
type Transporter struct {
	name      string
	is_loaded bool
}

type Truck struct {
	name         string
	max_load     int
	waiting_time int
}

type Node struct {
	box         *Box
	transporter *Transporter
	truck       *Truck
	x           int
	y           int
}

type WarehouseSquareGraph struct {
	nodes  []Node
	width  int
	height int
}

func (graph *WarehouseSquareGraph) initializeNodes() {
	graph.nodes = make([]Node, graph.width*graph.height)

	for i := 0; i < graph.width*graph.height; i++ {
		graph.nodes[i] = Node{x: i % graph.width, y: i / graph.height, box: nil, transporter: nil, truck: nil}
	}
}

func (graph *WarehouseSquareGraph) addPackage(name string, x, y, color int) {
	graph.nodes[x+(y*graph.height)].box = &Box{name: name, color: color}
}

func (graph *WarehouseSquareGraph) addTruck(name string, x, y, max_load, waiting_time int) {
	graph.nodes[x+(y*graph.height)].truck = &Truck{name: name, max_load: max_load, waiting_time: waiting_time}
}

func (graph *WarehouseSquareGraph) addTransporter(name string, x, y int) {
	graph.nodes[x+(y*graph.height)].transporter = &Transporter{name: name, is_loaded: false}
}

func (graph *WarehouseSquareGraph) getTransporterList() []Transporter {
	transporter_list := []Transporter{}

	for x := 0; x < graph.width; x++ {
		for y := 0; y < graph.height; y++ {
			if graph.nodes[x+(y*graph.height)].transporter != nil {
				transporter_list = append(transporter_list, *graph.nodes[x+(y*graph.height)].transporter)
			}
		}
	}
	return transporter_list
}

func (graph *WarehouseSquareGraph) moveTransporterToNearestBox(transporter_name string) {
}

func (graph *WarehouseSquareGraph) print() {
	for i := 0; i < graph.width*graph.height; i++ {
		if i != 0 && i%graph.height == 0 {
			fmt.Print("\n")
		}
		node := graph.nodes[i]
		if node.box != nil {
			fmt.Print("[x]")
		} else if node.transporter != nil {
			fmt.Print("[t]")
		} else if node.truck != nil {
			fmt.Print("[g]")
		} else {
			fmt.Print("[ ]")
		}
	}
	fmt.Print("\n")
}
