package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

func run () {
	grid := GridInit(10, 10)

	location1 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location2 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location3 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location4 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location5 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location6 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location7 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location8 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}
	location9 := Location{randomCoordinate(1,grid.width), randomCoordinate(1,grid.height), randomCellValue()}

	locations := []Location{location1, location2, location3,location4, location5, location6,location7, location8, location9}
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


func randomCellValue() bool {
	return rand.Float32() < 0.5
}

func randomCoordinate(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}