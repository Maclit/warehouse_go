package main

import "fmt"

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, start_node Node) {
	fmt.Print(start_node.transporter.name)
	if isWarehouseEmpty(graph) {
		fmt.Print(" WAIT")
		return
	}
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
		graph.nodes[transporter_x+(transporter_y*graph.height)].transporter.weight = graph.nodes[box_x+(box_y*graph.height)].box.color
		graph.nodes[transporter_x+(transporter_y*graph.height)].transporter.box_name = graph.nodes[box_x+(box_y*graph.height)].box.name
		fmt.Printf(" TAKE %s ", graph.nodes[box_x+(box_y*graph.height)].box.name)
		switch graph.nodes[box_x+(box_y*graph.height)].box.color {
		case YELLOW:
			fmt.Print("YELLOW\n")
		case GREEN:
			fmt.Print("GREEN\n")
		case BLUE:
			fmt.Print("BLUE\n")
		}
		graph.nodes[box_x+(box_y*graph.height)].box = nil
	} else {
		new_x := shortestPath[1].point.x
		new_y := shortestPath[1].point.y
		old_x := start_node.point.x
		old_y := start_node.point.y
		graph.nodes[new_x+(new_y*graph.height)].transporter = graph.nodes[old_x+(old_y*graph.height)].transporter
		graph.nodes[old_x+(old_y*graph.height)].transporter = nil
		fmt.Printf(" GO [%d,%d]\n", new_x, new_y)
	}
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, start_node Node) {
	fmt.Print(start_node.transporter.name)
	closest_box := findClosestTruck(graph, start_node)
	if closest_box.truck == nil {
		fmt.Print(" WAIT")
		return
	}
	shortestPath := shortestPath(graph, start_node, closest_box, make([]Node, 0))
	if len(shortestPath) == 1 {
		x := shortestPath[0].point.x
		y := shortestPath[0].point.y
		if graph.nodes[x+(y*graph.height)].truck.is_gone {
			return
		}
		graph.nodes[x+(y*graph.height)].truck.current_load += graph.nodes[x+(y*graph.height)].transporter.weight
		fmt.Printf(" LEAVE %s ", graph.nodes[x+(y*graph.height)].transporter.box_name)
		switch graph.nodes[x+(y*graph.height)].transporter.weight {
		case YELLOW:
			fmt.Print("YELLOW\n")
		case GREEN:
			fmt.Print("GREEN\n")
		case BLUE:
			fmt.Print("BLUE\n")
		}
		graph.nodes[x+(y*graph.height)].transporter.weight = 0
		graph.nodes[x+(y*graph.height)].transporter.is_loaded = false
		if isWarehouseEmpty(graph) {
			graph.nodes[x+(y*graph.height)].truck.is_gone = true
		}
	} else {
		new_x := shortestPath[1].point.x
		new_y := shortestPath[1].point.y
		old_x := start_node.point.x
		old_y := start_node.point.y
		graph.nodes[new_x+(new_y*graph.height)].transporter = graph.nodes[old_x+(old_y*graph.height)].transporter
		graph.nodes[old_x+(old_y*graph.height)].transporter = nil
		fmt.Printf(" GO [%d,%d]\n", new_x, new_y)
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
