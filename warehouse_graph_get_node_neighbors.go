package main

func getLeftEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	var x, y int

	if node.point.x > 0 && node.point.x <= graph.width-1 {
		x = node.point.x - 1
		y = node.point.y
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			return graph.nodes[x+(y*graph.height)], nil
		}
	}
	return Node{}, NoNeighborNodeError("Left")
}

func getRightEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	var x, y int

	if node.point.x >= 0 && node.point.x < graph.width-1 {
		x = node.point.x + 1
		y = node.point.y
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			return graph.nodes[x+(y*graph.height)], nil
		}
	}
	return Node{}, NoNeighborNodeError("Right")
}

func getTopEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	var x, y int

	if node.point.y > 0 && node.point.y <= graph.height-1 {
		x = node.point.x
		y = node.point.y - 1
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			return graph.nodes[x+(y*graph.height)], nil
		}
	}
	return Node{}, NoNeighborNodeError("Top")
}

func getBotEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	var x, y int

	if node.point.y >= 0 && node.point.y < graph.height-1 {
		x = node.point.x
		y = node.point.y + 1
		if graph.nodes[x+(y*graph.height)].transporter == nil {
			return graph.nodes[x+(y*graph.height)], nil
		}
	}
	return Node{}, NoNeighborNodeError("Bot")
}

func getEmptyNeighbor(graph *WarehouseSquareGraph, node Node) (Node, error) {
	leftNeighbor, err := getLeftEmptyNeighbor(graph, node)
	if err == nil {
		return leftNeighbor, nil
	}
	rightNeighbor, err := getRightEmptyNeighbor(graph, node)
	if err == nil {
		return rightNeighbor, nil
	}
	topNeighbor, err := getTopEmptyNeighbor(graph, node)
	if err == nil {
		return topNeighbor, nil
	}
	botNeighbor, err := getBotEmptyNeighbor(graph, node)
	if err == nil {
		return botNeighbor, nil
	}
	return Node{}, StuckTransporterError(node.transporter.name)
}

func getAllNeighborsNode(graph *WarehouseSquareGraph, node Node) []Node {
	neighbors := []Node{}

	leftNeighbor, err := getLeftEmptyNeighbor(graph, node)
	if err == nil {
		neighbors = append(neighbors, leftNeighbor)
	}
	rightNeighbor, err := getRightEmptyNeighbor(graph, node)
	if err == nil {
		neighbors = append(neighbors, rightNeighbor)
	}
	topNeighbor, err := getTopEmptyNeighbor(graph, node)
	if err == nil {
		neighbors = append(neighbors, topNeighbor)
	}
	botNeighbor, err := getBotEmptyNeighbor(graph, node)
	if err == nil {
		neighbors = append(neighbors, botNeighbor)
	}
	return neighbors
}
