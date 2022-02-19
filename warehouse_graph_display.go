package main

import "fmt"

func (graph *WarehouseSquareGraph) Print() {
	for i := 0; i < graph.width*graph.height; i++ {
		if i != 0 && i%graph.height == 0 {
			fmt.Print("\n")
		}
		node := graph.nodes[i]
		if node.box != nil {
			fmt.Print("[x]")
		} else if node.transporter != nil {
			fmt.Print("[t]")
		} else if node.truck != nil {
			fmt.Print("[g]")
		} else {
			fmt.Print("[ ]")
		}
	}
	fmt.Print("\n")
}
