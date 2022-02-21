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
	name      string
	is_loaded bool
	weight    int
	box_name  string
}

type Truck struct {
	name          string
	max_load      int
	current_load  int
	waiting_time  int
	current_timer int
	is_gone       bool
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
