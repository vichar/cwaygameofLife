package main

import "fmt"

type Grid struct {
	width  int
	height int
	pixels [][]bool
}

func (g Grid) LoadOne(x int, y int, value bool) {
	var row []bool
	if g.pixels[y] == nil {
		row = make([]bool, g.width)
	} else {
		row = g.pixels[y]
	}
	row[x] = row[x] || value
	g.pixels[y] = row
}

type Location struct {
	x     int
	y     int
	pixel bool
}

func GridInit(w int, h int) *Grid {
	p := make([][]bool, h)
	for i := range p {
		p[i] = make([]bool, w)
	}
	g := Grid{width: w, height: h, pixels: p}
	return &g
}

func (g Grid) Load(locations []Location) {
	for _, location := range locations {
		g.LoadOne(location.x, location.y, location.pixel)
	}
}

func (g Grid) Evaluate() {
	for i := 0; i < len(g.pixels); i++ {
		for j := 0; j < len(g.pixels[i]); j++ {
			fmt.Print(g.pixels[i][j])
		}
		fmt.Println()
	}
}
