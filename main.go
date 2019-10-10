package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)
func run()  {
	grid := GridInit(10,10)
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, float64(grid.width*50), float64(grid.height*50)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.EndShape = imdraw.SharpEndShape
	for i := 25.0; i <= float64(grid.height*50.0); i += 50.0 {
		imd.Color = colornames.Black
		imd.Push(pixel.V(i, 470))
		imd.Line(50)
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

