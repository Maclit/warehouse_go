package main

import "fmt"

type Array []string

func (arr Array) contains(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, start_node Node) {
	closest_box := findClosestBox(graph, start_node)
	if closest_box.box == nil {
		fmt.Println("No box found.")
		return
	}
	fmt.Println("Found box : ", closest_box)
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, start_node Node) error {
	return nil
}

func findClosestTruck(graph *WarehouseSquareGraph, start_node Node) (Node, error) {
	return Node{}, nil
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
	if node.point.x >= 0 && node.point.x < graph.width-1 { // Left
		x = node.point.x + 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.x > 0 && node.point.x <= graph.width-1 { // Left
		x = node.point.x - 1
		y = node.point.y
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.y >= 0 && node.point.y < graph.height-1 { // Left
		x = node.point.x
		y = node.point.y + 1
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	if node.point.y > 0 && node.point.y <= graph.height-1 { // Left
		x = node.point.x
		y = node.point.y - 1
		node := graph.nodes[x+(y*graph.height)]
		neighbors = append(neighbors, node)
	}
	return neighbors
}
