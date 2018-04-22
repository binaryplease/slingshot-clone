package main

import (
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"time"
)

func main() {
	// Run
	pixelgl.Run(run)
}

func run() {
	fps := time.Tick(time.Second / 120)

	sg := NewSlingshotGame(3, 2, 1000, 800)
	for !sg.win.Closed() {
		sg.Update()
		sg.win.Update()
		<-fps
	}

}
