package main

type Grid struct {
	width  int
	height int
	pixels [][]bool
}

func (g Grid) LoadOne(x int, y int, value bool) {
	row := make([]bool, g.width)
	row[x] = value
	g.pixels[y] = row
}

type Location struct {
	x     int
	y     int
	pixel bool
}

func GridInit(w int, h int) *Grid {

	p := make([][]bool, w)
	for i := range p {
		p[i] = make([]bool, h)
	}
	g := Grid{width: w, height: h, pixels: p}
	return &g
}

func (g Grid) Load(locations []Location) {
	for _, location := range locations {
		g.LoadOne(location.x, location.y, location.pixel)
	}
}
