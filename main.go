package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run () {
	location1 := Location{1, 1, true}
	location2 := Location{1, 2, true}
	location3 := Location{1, 3, true}
	location4 := Location{2, 1, true}
	location5 := Location{2, 2, true}
	location6 := Location{2, 3, true}
	location7 := Location{3, 1, true}
	location8 := Location{3, 2, true}
	location9 := Location{3, 3, true}

	locations := []Location{location1, location2, location3,location4, location5, location6,location7, location8, location9}
	grid := GridInit(10, 10)
	grid.Load(locations)
	cfg := pixelgl.WindowConfig{
		Title:  "Conway Game of Life",
		Bounds: pixel.R(0, 0, float64(grid.width*50), float64(grid.height*50)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.SharpEndShape
	imd.Color = colornames.Black

	if err != nil {
		panic(err)
	}
	for i := 0; i < len(grid.pixels); i++ {
		for j := 0; j < len(grid.pixels[i]); j++ {
			if grid.pixels[i][j] {
				imd.Push(pixel.V(float64(i*50), float64(j*60)))
				imd.Circle(10, 25)
			}
		}
	}
	grid.Evaluate()

	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		imd.Draw(win)
		win.Update()
	}
}
func main() {
	pixelgl.Run(run)
}
