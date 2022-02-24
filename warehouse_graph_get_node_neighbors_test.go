package main

import (
	"testing"
)

func TestGetLeftEmptyNeighbor(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node, err := getLeftEmptyNeighbor(warehouse, transporterList[0])

	if (err != nil || node == Node{}) {
		t.Errorf("the programm detect an error in getLeftEmpty")
	}
}

func TestGetRightEmptyNeighbor(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node, err := getRightEmptyNeighbor(warehouse, transporterList[0])

	if (err != nil || node == Node{}) {
		t.Errorf("the programm detect an error in getRightEmpty")
	}
}

func TestGetTopEmptyNeighbor(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node, err := getTopEmptyNeighbor(warehouse, transporterList[0])

	if (err != nil || node == Node{}) {
		t.Errorf("the programm detect an error in getTopEmpty")
	}
}

func TestGetBotEmptyNeighbor(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node, err := getBotEmptyNeighbor(warehouse, transporterList[0])

	if (err != nil || node == Node{}) {
		t.Errorf("the programm detect an error in getBotEmpty")
	}
}

func TestGetEmptyNeighbor(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node, err := getEmptyNeighbor(warehouse, transporterList[0])

	if (err != nil || node == Node{}) {
		t.Errorf("the programm detect an error in getEmptyNeighbor")
	}
}

func TestGetAllNeighborsNode(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	warehouse.AddTransporter("name", 3, 2)

	transporterList := warehouse.GetWharehouseNodeListWithObject(TRANSPORTER)
	node := getAllNeighborsNode(warehouse, transporterList[0])

	if node == nil {
		t.Errorf("the programm detect an error in getAllNeighborsNode")
	}
}
