package main

type Grid struct {
	width int
	height int
	pixels [][]byte
}

type Location struct {
	x int
	y int
	pixel byte
}
func GridInit(w int, h int) *Grid {
	//var p [][]byte
	g := Grid{ width: w, height: h}
	return &g
}

func (g Grid) LoadOne(x int, y int, value byte) {
	g.pixels[x][y] = value
}

func (g Grid) Load(locations []Location){
	for _, location := range locations {
		g.LoadOne(location.x,location.y, location.pixel)
	}
}

