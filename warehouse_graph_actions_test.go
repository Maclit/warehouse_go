package main

import (
	"testing"
)

func TestMoveTransporterTowardNearestBox(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddPackage("name", 3, 3, GREEN)
	warehouse.AddTransporter("transporter1", 2, 3)

	transporter := Transporter{
		name:     "transporter1",
		isLoaded: false,
		box:      nil,
	}
	node1 := Node{
		transporter: &transporter,
		point: Point{
			x: 2,
			y: 3,
		},
	}
	err := warehouse.MoveTransporterTowardNearestBox(node1)
	if err != nil {
		t.Errorf("an error is detected during movement")
	}
}

func TestMoveTransporterTowardNearestTruck(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddTruck("name", 3, 3, 10, 10)
	warehouse.AddTransporter("transporter1", 1, 2)

	transporter := Transporter{
		name:     "transporter1",
		isLoaded: false,
		box:      nil,
	}
	node1 := Node{
		transporter: &transporter,
		point: Point{
			x: 1,
			y: 2,
		},
	}
	err := warehouse.MoveTransporterTowardNearestBox(node1)
	if err != nil {
		t.Errorf("an error is detected during movement")
	}
}

func TestUpdateTruckStatus(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddTruck("name", 3, 3, 10, 10)

	truck := Truck{
		name:         "name",
		maxLoad:      10,
		currentLoad:  5,
		maxTimer:     10,
		currentTimer: 0,
		isGone:       false,
	}
	node1 := Node{
		truck: &truck,
		point: Point{
			x: 3,
			y: 3,
		},
	}
	err := warehouse.UpdateTruckStatus(node1)
	if err != nil {
		t.Errorf("update status truck failed")
	}
}

func TestMoveTransportToNextPosition(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddTransporter("name", 3, 3)
	point1 := Point{
		x: 3,
		y: 3,
	}
	point2 := Point{
		x: 3,
		y: 2,
	}
	warehouse.moveTransporterToNextPosition(point1, point2)
	if warehouse.nodes[point2.x+(point2.y*6)].point.y != 2 {
		t.Errorf("the transporter does not change position")
	}
}

func TestLoadtransporter(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddPackage("name", 3, 3, GREEN)
	warehouse.AddTransporter("transporter1", 2, 3)
	point1 := Point{
		x: 3,
		y: 3,
	}
	point2 := Point{
		x: 2,
		y: 3,
	}
	warehouse.loadTransporter(point2, point1)
	if len(warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)) > 1 || len(warehouse.GetWharehouseNodeListWithObject(BOX)) > 0 {
		t.Errorf("object")
	}
}

func TestUnloadTransporter(t *testing.T) {
	warehouse := CreateWarehouseGraph(6, 6)
	warehouse.AddPackage("name", 3, 1, GREEN)
	warehouse.AddTransporter("transporter1", 3, 2)
	warehouse.AddTruck("truck", 3, 3, 200, 200)

	point1 := Point{
		x: 3,
		y: 1,
	}
	point2 := Point{
		x: 3,
		y: 2,
	}
	point3 := Point{
		x: 3,
		y: 3,
	}
	warehouse.loadTransporter(point2, point1)
	warehouse.moveTransporterToNextPosition(point2, point3)
	warehouse.unloadTransporter(point3)
	if warehouse.nodes[point3.x+(point3.y*6)].truck.currentLoad != 200 {
		t.Errorf("actuel load is different %d instead of %d", warehouse.nodes[point3.x+(point3.y*6)].truck.currentLoad, 200)
	}
}
