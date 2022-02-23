package main

const (
	GREEN  int = 200
	YELLOW int = 100
	BLUE   int = 500
)

type Point struct {
	x int
	y int
}

type Box struct {
	name  string
	color int
}

type Transporter struct {
	name     string
	isLoaded bool
	box      *Box
}

type Truck struct {
	name         string
	maxLoad      int
	currentLoad  int
	maxTimer     int
	currentTimer int
	isGone       bool
}

type Node struct {
	box         *Box
	transporter *Transporter
	truck       *Truck
	point       Point
}

type WarehouseSquareGraph struct {
	nodes  []Node
	width  int
	height int
}
