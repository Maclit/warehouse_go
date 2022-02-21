package main

import "fmt"

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, start_node Node) {
	closest_box := findClosestBox(graph, start_node)
	if closest_box.box == nil {
		return
	}
	shortestPath := shortestPath(graph, start_node, closest_box, make([]Node, 0))
	if len(shortestPath) == 2 {
		box_x := shortestPath[1].point.x
		box_y := shortestPath[1].point.y
		transporter_x := start_node.point.x
		transporter_y := start_node.point.y
		graph.nodes[transporter_x+(transporter_y*graph.height)].transporter.is_loaded = true
		graph.nodes[box_x+(box_y*graph.height)].box = nil
	} else {
		new_x := shortestPath[1].point.x
		new_y := shortestPath[1].point.y
		old_x := start_node.point.x
		old_y := start_node.point.y
		graph.nodes[new_x+(new_y*graph.height)].transporter = graph.nodes[old_x+(old_y*graph.height)].transporter
		graph.nodes[old_x+(old_y*graph.height)].transporter = nil
	}
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, start_node Node) {
	closest_box := findClosestTruck(graph, start_node)
	if closest_box.truck == nil {
		return
	}
	shortestPath := shortestPath(graph, start_node, closest_box, make([]Node, 0))
	if len(shortestPath) == 1 {
		// Load truck
	} else {
		new_x := shortestPath[1].point.x
		new_y := shortestPath[1].point.y
		old_x := start_node.point.x
		old_y := start_node.point.y
		graph.nodes[new_x+(new_y*graph.height)].transporter = graph.nodes[old_x+(old_y*graph.height)].transporter
		graph.nodes[old_x+(old_y*graph.height)].transporter = nil
	}
}

func shortestPath(graph *WarehouseSquareGraph, start Node, end Node, path []Node) []Node {
	path = append(path, start)
	if isNodeInArray(path, end) {
		return path
	}
	shortest := make([]Node, 0)
	neighbors := getAllNeighborsNode(graph, start)
	for _, node := range neighbors {
		if !isNodeInArray(path, node) {
			newPath := shortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}

func findClosestTruck(graph *WarehouseSquareGraph, start_node Node) Node {
	if start_node.truck != nil {
		return start_node
	}
	to_visit := getAllNeighborsNode(graph, start_node)
	for {
		if len(to_visit) == 0 {
			fmt.Println()
			return Node{}
		}
		next_node := to_visit[0]
		to_visit = to_visit[1:]
		if next_node.truck != nil {
			return next_node
		}
		neightbors_list := getAllNeighborsNode(graph, next_node)
		if len(neightbors_list) > 0 {
			to_visit = append(to_visit, neightbors_list...)
		}
	}
}

func findClosestBox(graph *WarehouseSquareGraph, start_node Node) Node {
	if start_node.box != nil {
		return start_node
	}
	to_visit := getAllNeighborsNode(graph, start_node)
	for {
		if len(to_visit) == 0 {
			fmt.Println()
			return Node{}
		}
		next_node := to_visit[0]
		to_visit = to_visit[1:]
		if next_node.box != nil {
			return next_node
		}
		neightbors_list := getAllNeighborsNode(graph, next_node)
		if len(neightbors_list) > 0 {
			to_visit = append(to_visit, neightbors_list...)
		}
	}
}

func getAllNeighborsNode(graph *WarehouseSquareGraph, node Node) []Node {
	var x, y int

	neighbors := []Node{}
	if node.point.x >= 0 && node.point.x < graph.width-1 {
		x = node.point.x + 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.x > 0 && node.point.x <= graph.width-1 {
		x = node.point.x - 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.y >= 0 && node.point.y < graph.height-1 {
		x = node.point.x
		y = node.point.y + 1
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.y > 0 && node.point.y <= graph.height-1 {
		x = node.point.x
		y = node.point.y - 1
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
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
