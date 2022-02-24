package main

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddPackage("box", 1, 1, GREEN)
	warehouse.AddTransporter("name", 3, 1)
	node1 := Node{
		point: Point{
			x: 3,
			y: 1,
		},
	}
	node2 := Node{
		point: Point{
			x: 2,
			y: 1,
		},
	}
	node3 := Node{
		point: Point{
			x: 1,
			y: 1,
		},
	}
	nodes := []Node{node1, node2, node3}
	shortest := warehouse.shortestPath(warehouse.nodes[3+(1*6)], warehouse.nodes[2+(1*6)], make([]Node, 0))
	for n := 0; n < len(shortest); n++ {
		if shortest[n].point.x != nodes[n].point.x || shortest[n].point.y != nodes[n].point.y {
			t.Errorf("shortest Path is different x: %d and %d    y: %d and %d", shortest[n].point.x, nodes[n].point.x, shortest[n].point.y, nodes[n].point.y)
			return
		}
	}
}

func TestFindClosestObject(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddTransporter("trans", 2, 2)
	warehouse.AddPackage("box", 2, 0, GREEN)
	warehouse.AddTruck("truck", 2, 3, 200, 200)
	node := warehouse.findClosestObject(warehouse.nodes[2+(2*6)], BOX)
	if node.point.x != warehouse.nodes[2+(0*6)].point.x || node.point.y != warehouse.nodes[2+(0*6)].point.y {
		t.Errorf("shortest Path is different x: %d and %d    y: %d and %d", warehouse.nodes[2+(0*6)].point.x, node.point.x, warehouse.nodes[2+(0*6)].point.y, node.point.y)
	}
}
