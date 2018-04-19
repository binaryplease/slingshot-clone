package main

import "github.com/faiface/pixel"

type SpaceShip struct {
	// pos   pixel.Matrix
	// speed pixel.Vec
	SpaceObject
	power  float64
	weapon int
}

func (ss *SpaceShip) shoot() SpaceObject {
	shot := Shot{NewSpaceObject(ss.pos, ss.angle, pixel.ZV, "img/shots/shot1.png"), 5}
	return *shot.SpaceObject
}
