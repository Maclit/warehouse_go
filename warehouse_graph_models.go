package main

const (
	// TRUCK : A truck object
	TRUCK = 0

	// BOX : A box object
	BOX = 1

	// TRANSPORTER : A transporter object
	TRANSPORTER = 2

	// GREEN A green box has a weight of 200
	GREEN int = 200

	// YELLOW A yellow box has a weight of 100
	YELLOW int = 100

	// BLUE A blue box has a weight of 500
	BLUE int = 500
)

// Point A point object, containing 2d coordiantes (x, y)
type Point struct {
	x int
	y int
}

// Box A box object, to be placed in a node or transporter
type Box struct {
	name  string
	color int
}

// Transporter A transporter object, to be placed in a node.
type Transporter struct {
	name     string
	isLoaded bool
	box      *Box
}

// Truck A truck object, to be placed in a node.
type Truck struct {
	name         string
	maxLoad      int
	currentLoad  int
	maxTimer     int
	currentTimer int
	isGone       bool
}

// Node A node of the WarehouseSquareGraph, represent a case in the square.
type Node struct {
	box         *Box
	transporter *Transporter
	truck       *Truck
	point       Point
}

// WarehouseSquareGraph A graph shaped as a square representing a warehouse.
type WarehouseSquareGraph struct {
	nodes  []Node
	width  int
	height int
}
