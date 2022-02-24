package main

import "testing"

func TestCreateWarehouseGrapg(t *testing.T) {
	warehouse := CreateWarehouseGraph(5, 5)
	if warehouse == nil {
		t.Errorf("there is an problem in the creation of WarehouseGraph")
	}
}
