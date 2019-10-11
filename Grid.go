package main

import (
	"fmt"
)

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
			fmt.Print(g.nextAlive(i,j))
			g.pixels[i][j] = g.nextAlive(i,j)
		}
		fmt.Println()
	}
}
func (g Grid) isalive(x int, y int)bool {

	return g.pixels[x][y]
}
func (g Grid) nextAlive(x int, y int) bool {
	count := 0
	if y+1 < len(g.pixels[x]) && g.isalive(x,y+1) {
		count++
	}
	if y-1 >= 0 && g.isalive(x,y-1) {
		count++
	}

	if  y-1 >= 0 &&  x+1 < len(g.pixels) && g.isalive(x+1,y-1) {
		count++

	}
	if  x+1 < len(g.pixels) && y+1 < len(g.pixels[x]) && g.isalive(x+1,y+1) {
		count++

	}
	if x-1 >= 0 && g.isalive(x-1,y) {
		count++

	}

	if x-1 >= 0 &&  g.isalive(x-1,y) {
		count++

	}
	if x-1 >= 0 && y+1 < len(g.pixels[x]) && g.isalive(x-1,y+1) {
		count++

	}
	if x-1 >= 0 && y-1 >= 0 && g.isalive(x-1,y-1) {
		count++

	}
	if count <  2 {
		return false
	}
	return g.pixels[x][y]
}

