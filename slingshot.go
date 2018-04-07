package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
)

type SlingshotGame struct {
	xSize   int
	ySize   int
	planets []Planet
	players []SlingshotPlayer
	win     *pixelgl.Window
}

func (sg *SlingshotGame) run() {

	for !sg.win.Closed() {
		sg.draw(sg.win)
		sg.win.Update()
	}

}

func (sg *SlingshotGame) drawPicture(path string) {
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(sg.win, pixel.IM.Moved(sg.win.Bounds().Center()))

	for !sg.win.Closed() {
		sg.win.Update()
	}
}

func NewSlingshotGame(numPlanets, numPlayers, xSize, ySize int) *SlingshotGame {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var planets []Planet
	var players []SlingshotPlayer

	sg := &SlingshotGame{xSize, ySize, planets, players, win}

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
	sg.drawPlanets()
	sg.drawShips()
	sg.drawScore()
}

func (sg *SlingshotGame) drawPlanets() {
	for _, v := range sg.planets {
		sg.drawPicture(v.image)
	}

}

func (sg *SlingshotGame) drawShips() {
	for _, v := range sg.players {
		sg.drawPicture(v.ship.image)
	}
}

func (sg SlingshotGame) drawScore() {
	for k, v := range sg.players {
		//TODO display on screen, not on console
		fmt.Println("Player " + string(k) + ": " + string(v.score) + "Points")
	}

}
