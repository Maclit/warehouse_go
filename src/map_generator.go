package main

const (
	GREEN  int = 0
	YELLOW int = 1
	BLUE   int = 2
)

type Coord struct {
	x int
	y int
}

type Package struct {
	name   string
	coords Coord
	color  int
}

type Transporter struct {
	name   string
	coords Coord
}

type Truck struct {
	name         string
	coord        Coord
	max_load     int
	waiting_time int
}

type Map struct {
	packages     []Package
	transporters []Transporter
	trucks       []Truck
	width        int
	height       int
}
