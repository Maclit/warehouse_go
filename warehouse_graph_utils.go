package main

import "fmt"

const (
	TRUCK       = 0
	BOX         = 1
	TRANSPORTER = 2
)

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
	new_x := end.x
	new_y := end.y
	old_x := start.x
	old_y := start.y
	graph.nodes[new_x+(new_y*graph.height)].transporter = graph.nodes[old_x+(old_y*graph.height)].transporter
	graph.nodes[old_x+(old_y*graph.height)].transporter = nil
	fmt.Printf(" GO [%d,%d]\n", new_x, new_y)
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
	return Node{}, ErrStuck(node.transporter.name)
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
		if node.transporter != nil && node.transporter.is_loaded {
			return false
		}
	}
	return true
}

func isTruckGone(graph *WarehouseSquareGraph) bool {
	for i := 0; i < graph.width*graph.height; i++ {
		node := graph.nodes[i]
		if node.truck != nil && !node.truck.is_gone {
			return false
		}
	}
	return true
}
