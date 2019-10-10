package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	location1 := Location{1, 1, true}
	location2 := Location{2, 2, true}
	location3 := Location{3, 3, true}
	locations := []Location{location1, location2, location3}
	grid := GridInit(10, 10)
	grid.Load(locations)
	cfg := pixelgl.WindowConfig{
		Title:  "Conway Game of Life",
		Bounds: pixel.R(0, 0, float64(grid.width*50), float64(grid.height*50)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.EndShape = imdraw.SharpEndShape
	for i := 0.0; i < float64(len(grid.pixels)); i++ {
		imd.Color = colornames.Black
		fmt.Println(grid.pixels)
		for j := 0.0; j < float64(len(grid.pixels[int(i)])); j++ {
			if grid.pixels[int(i)][int(j)] {
				imd.Push(pixel.V(i, float64(len(grid.pixels[int(i)]))))
				imd.Line(50)
			}
		}

	}

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {

	pixelgl.Run(run)
}
