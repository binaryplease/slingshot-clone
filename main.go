package main

import (
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

func main() {
	// Run
	pixelgl.Run(run)
}

func run() {

	sg := NewSlingshotGame(3, 2, 1000, 800)
	for !sg.win.Closed() {
		sg.Update()
		sg.win.Update()
	}

}
