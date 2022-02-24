package main

func (graph *WarehouseSquareGraph) shortestPath(start Node, end Node, path []Node) []Node {
	path = append(path, start)
	if graph.isNodeInArray(path, end) {
		return path
	}
	shortest := make([]Node, 0)
	neighbors := getAllNeighborsNode(graph, start)
	for _, node := range neighbors {
		if !graph.isNodeInArray(path, node) {
			newPath := graph.shortestPath(node, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func (graph *WarehouseSquareGraph) findClosestObject(startNode Node, toFind int) Node {
	toVisit := getAllNeighborsNode(graph, startNode)
	visited := []Node{}
	for len(toVisit) > 0 {
		nextNode := toVisit[0]
		toVisit = toVisit[1:]
		visited = append(visited, nextNode)
		if (toFind == TRUCK && nextNode.truck != nil) || (toFind == BOX && nextNode.box != nil) {
			return nextNode
		}
		neightborsList := getAllNeighborsNode(graph, nextNode)
		for _, node := range neightborsList {
			if !graph.isNodeInArray(visited, node) {
				toVisit = append(toVisit, node)
			}
		}
	}
	return Node{}
}
