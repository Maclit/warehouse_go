package main

import "fmt"

func moveTransporterTowardNearestBox(graph *WarehouseSquareGraph, start_node Node) {
	fmt.Print(start_node.transporter.name)
	if isWarehouseEmpty(graph) {
		if graph.doesNodeHasObject(start_node.point, TRUCK) {
			neighbor, err := getEmptyNeighbor(graph, start_node)
			if err != nil {
				fmt.Println(err.Error())
			}
			moveTransporterToNextPosition(graph, start_node.point, neighbor.point)
		} else {
			fmt.Print(" WAIT\n")
		}
		return
	}
	closest_box := findClosestObject(graph, start_node, BOX)
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
		moveTransporterToNextPosition(graph, start_node.point, shortestPath[1].point)
	}
}

func moveTransporterTowardNearestTruck(graph *WarehouseSquareGraph, start_node Node) {
	fmt.Print(start_node.transporter.name)
	if graph.doesNodeHasObject(start_node.point, TRUCK) {
		x := start_node.point.x
		y := start_node.point.y
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
	} else {
		closest_box := findClosestObject(graph, start_node, TRUCK)
		if closest_box.truck == nil {
			fmt.Print(" WAIT\n")
			return
		}
		shortestPath := shortestPath(graph, start_node, closest_box, make([]Node, 0))
		moveTransporterToNextPosition(graph, start_node.point, shortestPath[1].point)
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

func findClosestObject(graph *WarehouseSquareGraph, start_node Node, to_find int) Node {
	if to_find == TRUCK && start_node.truck != nil {
		return start_node
	}
	if to_find == BOX && start_node.box != nil {
		return start_node
	}
	to_visit := getAllNeighborsNode(graph, start_node)
	visited := []Node{}
	for {
		if len(to_visit) == 0 {
			return Node{}
		}
		next_node := to_visit[0]
		to_visit = to_visit[1:]
		visited = append(visited, next_node)
		if to_find == TRUCK && next_node.truck != nil {
			return next_node
		}
		if to_find == BOX && next_node.box != nil {
			return next_node
		}
		neightbors_list := getAllNeighborsNode(graph, next_node)
		if len(neightbors_list) > 0 {
			for _, node := range neightbors_list {
				if !isNodeInArray(visited, node) {
					to_visit = append(to_visit, node)
				}
			}
		}
	}
}
