package main

const (
	GREEN  int = 0
	YELLOW int = 1
	BLUE   int = 2
)

type Package struct {
	name  string
	color int
	x     int
	y     int
}

func (pack *Package) move(x, y int) {
	pack.x = x
	pack.y = y
}
