package main

import "github.com/faiface/pixel"

type SpaceShip struct {
	// pos   pixel.Matrix
	// speed pixel.Vec
	pos    pixel.Vec
	angle  float64
	power  float64
	weapon int
	image  string
}

func (ss *SpaceShip) shoot() {
	speed := pixel.V(0, 1)
	shot := &Shot{ss.pos, speed, 5, "img/shots/shot1.png"}
	shot.move()
}
