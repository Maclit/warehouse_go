package main

import (
	"testing"
)

func TestValidateTruck(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	err := warehouse.Validate()
	if err == nil {
		t.Errorf("error no element is not detected")
	}
}

func TestValidateTransporter(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTruck("name", 5, 5, 5, 5)
	err := warehouse.Validate()
	if err == nil {
		t.Errorf("error no transporters is not detected")
	}
}

func TestValidate(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTruck("name", 2, 2, 5, 5)
	warehouse.AddTransporter("name2", 3, 3)
	err := warehouse.Validate()
	if err != nil {
		t.Errorf("error valide map")
	}
}

func TestAreAlTrucksGone(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTruck("name", 2, 3, 5, 5)
	warehouse.nodes[2+(3*5)].truck.isGone = true
	if warehouse.AreAllTrucksGone() == false {
		t.Errorf("non detected truck")
	}
}

func TestAreAlTrucksGoneError(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTruck("name", 2, 3, 5, 5)
	warehouse.nodes[2+(3*5)].truck.isGone = false
	if warehouse.AreAllTrucksGone() == true {
		t.Errorf("detected truck")
	}
}

func TestGetWharehouseNodeListWithObject(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)
	warehouse.AddTransporter("name2", 4, 1)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	if len(transporterList) != 2 {
		t.Errorf("error size %d instead of 2", len(transporterList))
	}
}

func TestIsEmpty(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	if warehouse.isEmpty() == false {
		t.Errorf("test")
	}
}

func TestDoesNodeHasObject(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)
	warehouse.AddTruck("name2", 4, 1, 4, 4)
	warehouse.AddPackage("name3", 0, 0, GREEN)
	point1 := Point{
		x: 3,
		y: 2,
	}
	point2 := Point{
		x: 4,
		y: 1,
	}
	point3 := Point{
		x: 0,
		y: 0,
	}
	if warehouse.doesNodeHasObject(point1, TRANSPORTER) == false {
		t.Errorf("object not detected as transporter")
	}
	if warehouse.doesNodeHasObject(point2, TRUCK) == false {
		t.Errorf("object not detected as truck")
	}
	if warehouse.doesNodeHasObject(point3, BOX) == false {
		t.Errorf("object not detected as box")
	}
	if warehouse.doesNodeHasObject(point2, BOX) == true {
		t.Errorf("object detected as Box")
	}
}

func TestAreCoordinatesValid(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	if warehouse.areCoordinatesValid(2, 3) == false {
		t.Errorf("coordinates are not found in the graph")
	}
	if warehouse.areCoordinatesValid(6, 5) == true {
		t.Errorf("coordinates are found in the graph")
	}
}

func TestIsNodeArray(t *testing.T) {
	warehouse := CreateWarehouseGraph(7, 7)
	box := Box{
		name:  "box",
		color: GREEN,
	}
	node1 := Node{
		box: &box,
		point: Point{
			x: 4,
			y: 5,
		},
	}
	transporter := Transporter{
		name:     "transporter",
		isLoaded: false,
		box:      nil,
	}
	node2 := Node{
		transporter: &transporter,
		point: Point{
			x: 3,
			y: 4,
		},
	}
	transporter2 := Transporter{
		name:     "transporter",
		isLoaded: false,
		box:      nil,
	}
	node3 := Node{
		transporter: &transporter2,
		point: Point{
			x: 1,
			y: 1,
		},
	}
	nodes := []Node{node1, node2}
	if warehouse.isNodeInArray(nodes, node1) == false {
		t.Errorf("does not find node in array nodes")
	}
	if warehouse.isNodeInArray(nodes, node3) == true {
		t.Errorf("find node in array nodes")
	}
}
