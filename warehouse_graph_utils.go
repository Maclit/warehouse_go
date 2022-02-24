package main

func (graph *WarehouseSquareGraph) isNodeInArray(array []Node, node Node) bool {
	for _, element := range array {
		if element.point.x == node.point.x && element.point.y == node.point.y {
			return true
		}
	}
	return false
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

func (graph *WarehouseSquareGraph) areCoordinatesValid(x, y int) bool {
	if x >= graph.width || y >= graph.height {
		return false
	}
	return true
}
