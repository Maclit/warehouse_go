package main

func main() {
	width := 5
	height := 5

	warehouse_graph := WarehouseSquareGraph{width: width, height: height}
	warehouse_graph.initializeNodes()
	warehouse_graph.addPackage("box_1", 1, 1, GREEN)
	warehouse_graph.addTransporter("transporter_1", 2, 0)
	warehouse_graph.addTruck("truck_1", 3, 4, 10, 50)
	warehouse_graph.print()
}
