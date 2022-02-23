package main

func createWarehouseGraph(height, width int) *WarehouseSquareGraph {
	warehouseGraph := WarehouseSquareGraph{width: width, height: height}

	warehouseGraph.initializeNodes()
	return &warehouseGraph
}

func (graph *WarehouseSquareGraph) initializeNodes() {
	graph.nodes = make([]Node, graph.width*graph.height)

	for i := 0; i < graph.width*graph.height; i++ {
		graph.nodes[i] = Node{point: Point{x: i % graph.width, y: i / graph.height}, box: nil, transporter: nil, truck: nil}
	}
}

// AddPackage Add a package object to the graph
func (graph *WarehouseSquareGraph) AddPackage(name string, x, y, color int) {
	box := Box{
		name:  name,
		color: color,
	}
	graph.nodes[x+(y*graph.height)].box = &box
}

// AddTruck Add a truck object to the graph
func (graph *WarehouseSquareGraph) AddTruck(name string, x, y, max_load, waiting_time int) {
	truck := Truck{
		name:         name,
		maxLoad:      max_load,
		currentLoad:  0,
		maxTimer:     waiting_time,
		currentTimer: 0,
		isGone:       false,
	}
	graph.nodes[x+(y*graph.height)].truck = &truck
}

// AddTransporter Add a transporter object to the graph
func (graph *WarehouseSquareGraph) AddTransporter(name string, x, y int) {
	transporter := Transporter{
		name:     name,
		isLoaded: false,
		box:      nil,
	}
	graph.nodes[x+(y*graph.height)].transporter = &transporter
}
