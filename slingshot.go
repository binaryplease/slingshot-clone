package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"io/ioutil"
	"math/rand"
)

type SlingshotGame struct {
	xSize   int
	ySize   int
	planets []Planet
	players []SlingshotPlayer
	win     *pixelgl.Window
}

func (sg *SlingshotGame) drawPicture(xPos, yPos, angle float64, path string) {
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	mat := pixel.IM
	mat = mat.Moved(pixel.V(xPos, yPos))

	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(sg.win, mat)

	for !sg.win.Closed() {
		sg.win.Update()
	}
}

func NewSlingshotGame(numPlanets, numPlayers, xSize, ySize int) *SlingshotGame {

	//Load planet images
	var planetImages []string
	files, err := ioutil.ReadDir("img/planets")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		planetImages = append(planetImages, "./img/planets/"+f.Name())
	}

	//Load ship images
	var shipImages []string
	files, err = ioutil.ReadDir("img/ships")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		shipImages = append(shipImages, "./img/ships/"+f.Name())
	}

	//Create window
	cfg := pixelgl.WindowConfig{
		Title:  "Slingshot",
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

	// Add Planets
	for i := 0; i < numPlanets; i++ {
		xPos := rand.Float64() * float64(xSize)
		yPos := rand.Float64() * float64(ySize)
		diam := rand.Float64() * float64(100)
		sg.addPlanet(xPos, yPos, diam, planetImages[i%len(planetImages)])
	}

	// Add Players
	for i := 0; i < numPlayers; i++ {
		xPos := rand.Float64() * float64(xSize)
		yPos := rand.Float64() * float64(ySize)
		sg.addPlayer(xPos, yPos, 90, shipImages[i%len(shipImages)])
	}

	return sg
}

// Add a planet to the game
func (sg *SlingshotGame) addPlanet(xPos, yPos, diameter float64, image string) {
	p := Planet{xPos, yPos, diameter, image}
	sg.planets = append(sg.planets, p)
}

// Add a player to the game
func (sg *SlingshotGame) addPlayer(xPos, yPos, angle float64, image string) {
	ship := SpaceShip{xPos, yPos, angle, image}
	player := SlingshotPlayer{ship, 0}
	sg.players = append(sg.players, player)
}

// Draw all images on the screen
func (sg SlingshotGame) draw(win *pixelgl.Window) {
	win.Clear(colornames.Skyblue)
	sg.drawPlanets()
	sg.drawShips()
	sg.drawScore()
}

// Draw the planets
func (sg *SlingshotGame) drawPlanets() {
	for _, v := range sg.planets {
		fmt.Println("Drawing Planet at:" + FloatToString(v.xPos) + "" + FloatToString(v.yPos))
		sg.drawPicture(v.xPos, v.yPos, 0, v.image)
	}
}

// Draw the ships
func (sg *SlingshotGame) drawShips() {
	for _, v := range sg.players {
		sg.drawPicture(v.ship.xPos, v.ship.yPos, v.ship.angle, v.ship.image)
	}
}

// Draw the players score
func (sg SlingshotGame) drawScore() {
	for k, v := range sg.players {
		//TODO display on screen, not on console
		fmt.Println("Player " + string(k) + ": " + string(v.score) + "Points")
	}
}
