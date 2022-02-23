package main

import "fmt"

// Print A method to print the graph on the terminal.
func (graph *WarehouseSquareGraph) Print() {
	for i := 0; i < graph.width*graph.height; i++ {
		if i != 0 && i%graph.height == 0 {
			fmt.Print("\n")
		}
		node := graph.nodes[i]
		if node.box != nil {
			fmt.Print("[x]")
			continue
		}
		if node.transporter != nil {
			fmt.Print("[t]")
			continue
		}
		if node.truck != nil {
			fmt.Print("[g]")
			continue
		}
		fmt.Print("[ ]")
	}
	fmt.Print("\n")
}
