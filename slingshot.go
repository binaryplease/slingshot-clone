package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	_ "image"
	_ "image/jpeg"
	"math"
	"math/rand"
	"time"
)

type SlingshotGame struct {
	xSize      int
	ySize      int
	planets    []Planet
	players    []SlingshotPlayer
	win        *pixelgl.Window
	turn       int
	cam        *SlingshotCamera
	atlas      *text.Atlas
	background string
	toDraw     chan SpaceObject
	fps        int
	particles  []SpaceObject
}

func (sg *SlingshotGame) Update() {
	sg.draw()
	sg.getInput()
	// time.Sleep(100 * time.Millisecond)

}

func (sg *SlingshotGame) getInput() {
	// Control the camera
	dt := time.Since(sg.cam.last).Seconds()
	sg.cam.last = time.Now()

	sg.cam.cam = pixel.IM.Scaled(sg.cam.camPos, sg.cam.camZoom).Moved(sg.win.Bounds().Center().Sub(sg.cam.camPos))
	sg.win.SetMatrix(sg.cam.cam)

	mouse := (sg.win.MousePosition())

	if mouse.X+100 > float64(sg.xSize) {
		sg.cam.camPos.X += sg.cam.camSpeed * dt
	}

	if mouse.X-100 < 0.0 {
		sg.cam.camPos.X -= sg.cam.camSpeed * dt
	}

	if mouse.Y+100 > float64(sg.ySize) {
		sg.cam.camPos.Y += sg.cam.camSpeed * dt
	}

	if mouse.Y-100 < 0.0 {
		sg.cam.camPos.Y -= sg.cam.camSpeed * dt
	}

	sg.cam.camZoom *= math.Pow(sg.cam.camZoomSpeed, sg.win.MouseScroll().Y)
	if sg.cam.camZoom > sg.cam.camMaxZoom {
		sg.cam.camZoom = sg.cam.camMaxZoom
	}

	if sg.cam.camZoom < sg.cam.camMinZoom {
		sg.cam.camZoom = sg.cam.camMinZoom
	}
	// Turn ship left
	if sg.win.Pressed(pixelgl.Key1) {
		sg.players[sg.turn].ship.angle += math.Pi / 100
	}

	// Turn ship right
	if sg.win.Pressed(pixelgl.Key2) {
		sg.players[sg.turn].ship.angle -= math.Pi / 100
	}

	// More power
	if sg.win.Pressed(pixelgl.Key3) {
		sg.players[sg.turn].ship.power += 10
	}

	//Less power
	if sg.win.Pressed(pixelgl.Key4) {
		sg.players[sg.turn].ship.power -= 10
	}

	if sg.win.Pressed(pixelgl.KeySpace) {
		so := sg.players[sg.turn].ship.shoot()
		sg.particles = append(sg.particles, so)
		sg.turn = (sg.turn + 1) % len(sg.players)
		time.Sleep(100 * time.Millisecond)
	}
}

func (sg *SlingshotGame) drawParticles() {
	for _, v := range sg.particles {
		//TODO remove after timeout
		sg.toDraw <- v
	}
}
func (sg *SlingshotGame) drawChan() {
	for {
		so := <-sg.toDraw
		so.update()
		for _, v := range sg.particles {
			v = v.update()
		}
		pic, err := loadPicture(so.image)
		if err != nil {
			panic(err)
		}

		sprite := pixel.NewSprite(pic, pic.Bounds())
		sprite.Draw(sg.win, pixel.IM.Moved(so.pos).Rotated(so.pos, so.angle))
	}
}

func NewSlingshotGame(numPlanets, numPlayers, xSize, ySize int) *SlingshotGame {

	// Pseudo random init
	rand.Seed(time.Now().Unix())

	//Load images
	backgroundImages := loadImageDir("./img/background")
	planetImages := loadImageDir("./img/planets")
	shipImages := loadImageDir("img/ships")

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

	//Create text for info
	face, err := loadTTF("font.ttf", 30)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)

	var planets []Planet
	var players []SlingshotPlayer
	var particles []SpaceObject

	cam := NewSlingshotCamera()
	background := backgroundImages[rand.Intn(len(backgroundImages))]
	toDraw := make(chan SpaceObject)
	sg := &SlingshotGame{xSize, ySize, planets, players, win, 0, cam, atlas, background, toDraw, 60, particles}

	go sg.drawChan()

	// Add Planets
	for i := 0; i < numPlanets; i++ {
		xPos := rand.Float64() * float64(xSize)
		yPos := rand.Float64() * float64(ySize)
		pos := (pixel.V(xPos, yPos))
		diam := rand.Float64() * float64(100)
		sg.addPlanet(pos, diam, planetImages[i%len(planetImages)])
	}

	// Add Players
	for i := 0; i < numPlayers; i++ {
		xPos := rand.Float64() * float64(xSize)
		yPos := rand.Float64() * float64(ySize)

		pos := pixel.V(xPos, yPos)
		sg.addPlayer(pos, 0, shipImages[i%len(shipImages)])
	}

	return sg
}

// Add a planet to the game
func (sg *SlingshotGame) addPlanet(pos pixel.Vec, diameter float64, image string) {
	p := Planet{*NewSpaceObject(pos, 0, pixel.ZV, image), diameter}
	sg.planets = append(sg.planets, p)
}

// Add a player to the game
func (sg *SlingshotGame) addPlayer(pos pixel.Vec, angle float64, image string) {
	ship := SpaceShip{*NewSpaceObject(pos, angle, pixel.ZV, image), 10, 1}
	player := SlingshotPlayer{ship, 0}
	sg.players = append(sg.players, player)
}

// Draw all images on the screen
func (sg SlingshotGame) draw() {
	sg.win.Clear(colornames.Black)
	sg.drawBackground()
	sg.drawPlanets()
	sg.drawShips()
	sg.drawScore()
	sg.drawParticles()
	time.Sleep(time.Second / time.Duration(sg.fps))
}

func (sg *SlingshotGame) drawBackground() {

	pic, err := loadPicture(sg.background)
	if err != nil {
		panic(err)
	}

	maxI := int((sg.win.Bounds().W()/pic.Bounds().W()+1)/sg.cam.camZoom) + 1
	maxJ := int((sg.win.Bounds().H()/pic.Bounds().H()+1)/sg.cam.camZoom) + 1

	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {

			w := float64(i) * pic.Bounds().W()
			h := float64(j) * pic.Bounds().H()
			mat := pixel.IM
			mat = mat.Moved(sg.cam.cam.Unproject(sg.win.Bounds().Min))
			mat = mat.Moved(pixel.V(w, h))
			sprite := pixel.NewSprite(pic, pic.Bounds())
			sprite.Draw(sg.win, mat)
		}

	}
}

// Draw the planets
func (sg *SlingshotGame) drawPlanets() {
	for _, v := range sg.planets {
		sg.toDraw <- v.SpaceObject
	}
}

// Draw the ships
func (sg *SlingshotGame) drawShips() {
	for _, v := range sg.players {
		sg.toDraw <- v.ship.SpaceObject
	}
}

// Draw the players score
func (sg SlingshotGame) drawScore() {
	txt := text.New(sg.cam.cam.Unproject(pixel.V(50, 500)), sg.atlas)

	// Print Players info in respective color
	for k, v := range sg.players {
		if sg.turn == k {
			txt.Color = colornames.Red
		} else {
			txt.Color = colornames.Grey
		}
		txt.WriteString("Payer: " + string(k+1) + ": " + string(v.score) + "\n")
	}

	mat := pixel.IM
	mat = mat.Moved(sg.cam.cam.Unproject(sg.win.Bounds().Min))
	txt.Draw(sg.win, pixel.IM.Moved(sg.cam.cam.Unproject(sg.win.Bounds().Max).Sub(txt.Bounds().Max)))
}
