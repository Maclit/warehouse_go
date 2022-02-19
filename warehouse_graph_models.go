package main

const (
	GREEN  int = 0
	YELLOW int = 1
	BLUE   int = 2
)

type Box struct {
	name  string
	color int
}

type Transporter struct {
	name      string
	is_loaded bool
}

type Truck struct {
	name         string
	max_load     int
	waiting_time int
}

type Node struct {
	box         *Box
	transporter *Transporter
	truck       *Truck
	x           int
	y           int
}

type WarehouseSquareGraph struct {
	nodes  []Node
	width  int
	height int
}
