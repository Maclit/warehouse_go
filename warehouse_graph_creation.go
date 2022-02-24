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

func (graph *WarehouseSquareGraph) addPackage(name string, x, y, color int) error {
	if !graph.areCoordinatesValid(x, y) {
		return BadGraphCoordinatesError("addPackage")
	}
	box := Box{
		name:  name,
		color: color,
	}
	graph.nodes[x+(y*graph.height)].box = &box
	return nil
}

func (graph *WarehouseSquareGraph) addTruck(name string, x, y, maxLoad, maxTimer int) error {
	if !graph.areCoordinatesValid(x, y) {
		return BadGraphCoordinatesError("addTruck")
	}
	truck := Truck{
		name:         name,
		maxLoad:      maxLoad,
		currentLoad:  0,
		maxTimer:     maxTimer,
		currentTimer: 0,
		isGone:       false,
	}
	graph.nodes[x+(y*graph.height)].truck = &truck
	return nil
}

func (graph *WarehouseSquareGraph) addTransporter(name string, x, y int) error {
	if !graph.areCoordinatesValid(x, y) {
		return BadGraphCoordinatesError("addTransporter")
	}
	transporter := Transporter{
		name:     name,
		isLoaded: false,
		box:      nil,
	}
	graph.nodes[x+(y*graph.height)].transporter = &transporter
	return nil
}
