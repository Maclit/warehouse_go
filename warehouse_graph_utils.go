package main

func getAllNeighborsNode(graph *WarehouseSquareGraph, node Node) []Node {
	var x, y int

	neighbors := []Node{}
	if node.point.x >= 0 && node.point.x < graph.width-1 {
		x = node.point.x + 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		if node.transporter == nil {
			neighbors = append(neighbors, node)
		}
	}
	if node.point.x > 0 && node.point.x <= graph.width-1 {
		x = node.point.x - 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		if node.transporter == nil {
			neighbors = append(neighbors, node)
		}
	}
	if node.point.y >= 0 && node.point.y < graph.height-1 {
		x = node.point.x
		y = node.point.y + 1
		node := graph.nodes[x+(y*graph.height)]
		if node.transporter == nil {
			neighbors = append(neighbors, node)
		}
	}
	if node.point.y > 0 && node.point.y <= graph.height-1 {
		x = node.point.x
		y = node.point.y - 1
		node := graph.nodes[x+(y*graph.height)]
		if node.transporter == nil {
			neighbors = append(neighbors, node)
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
		if node.truck != nil && !node.truck.is_gone {
			return false
		}
	}
	return true
}
