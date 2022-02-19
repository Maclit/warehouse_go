package main

func createWarehouseGraph(height, width int) *WarehouseSquareGraph {
	warehouse_graph := WarehouseSquareGraph{width: width, height: height}

	warehouse_graph.initializeNodes()
	return &warehouse_graph
}

func (graph *WarehouseSquareGraph) initializeNodes() {
	graph.nodes = make([]Node, graph.width*graph.height)

	for i := 0; i < graph.width*graph.height; i++ {
		graph.nodes[i] = Node{x: i % graph.width, y: i / graph.height, box: nil, transporter: nil, truck: nil}
	}
}

func (graph *WarehouseSquareGraph) AddPackage(name string, x, y, color int) {
	graph.nodes[x+(y*graph.height)].box = &Box{name: name, color: color}
}

func (graph *WarehouseSquareGraph) AddTruck(name string, x, y, max_load, waiting_time int) {
	graph.nodes[x+(y*graph.height)].truck = &Truck{name: name, max_load: max_load, waiting_time: waiting_time}
}

func (graph *WarehouseSquareGraph) AddTransporter(name string, x, y int) {
	graph.nodes[x+(y*graph.height)].transporter = &Transporter{name: name, is_loaded: false}
}
