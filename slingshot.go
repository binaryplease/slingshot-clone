package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
)

type SlingshotGame struct {
	xSize   int
	ySize   int
	planets []Planet
	players []SlingshotPlayer
}

func NewSlingshotGame(numPlanets, numPlayers, xSize, ySize int) *SlingshotGame {

	var planets []Planet
	var players []SlingshotPlayer

	sg := &SlingshotGame{xSize, ySize, planets, players}

	// Add Players
	for i := 0; i < numPlanets; i++ {
		xPos := rand.Intn(xSize)
		yPos := rand.Intn(ySize)
		diam := rand.Intn(100)
		image := "p1.png"
		sg.addPlanet(xPos, yPos, diam, image)
	}

	// Add Players
	for i := 0; i < numPlayers; i++ {
		sg.addPlayer()
	}

	return sg
}

func (r *SlingshotGame) addPlanet(xPos, yPos, diameter int, graphic string) {

}

func (s *SlingshotGame) addPlayer() {

}

// Draw all images on the screen
func (sg SlingshotGame) draw(win *pixelgl.Window) {
	win.Clear(colornames.Skyblue)
	drawPlanets(sg.planets)
	drawShips(sg.players)
	drawScore(sg.players)
}

func drawPlanets(planets []Planet) {
	for _, v := range planets {
		drawPicture(v.image)
	}

}

func drawShips(players []SlingshotPlayer) {
	for _, v := range players {
		drawPicture(v.ship.image)
	}
}

func drawScore(players []SlingshotPlayer) {
	for k, v := range players {
		//TODO display on screen, not on console
		fmt.Println("Player " + string(k) + ": " + string(v.score) + "Points")
	}

}
