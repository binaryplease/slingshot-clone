package main

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
)

type SlingshotGame struct {
	xSize   int
	ySize   int
	planets []Planet
}

func NewSlingshotGame(numPlanets int) *SlingshotGame {

	var planets []Planet
	xSize := 800
	ySize := 800

	sg := &SlingshotGame{xSize, ySize, planets}

	for i := 0; i < numPlanets; i++ {
		xPos := rand.Intn(xSize)
		yPos := rand.Intn(ySize)
		diam := rand.Intn(100)
		image := "p1.png"
		sg.addPlanet(xPos, yPos, diam, image)
	}

	return sg
}

func (r *SlingshotGame) addPlanet(xPos, yPos, diameter int, graphic string) {

}

func (sg SlingshotGame) draw(win *pixelgl.Window) {
	win.Clear(colornames.Skyblue)
}
