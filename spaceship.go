package main

import "github.com/faiface/pixel"

type SpaceShip struct {
	SpaceObject
	power  float64
	weapon int
}

func (ss *SpaceShip) shoot() SpaceObject {
	shot := Shot{NewSpaceObject(ss.pos, ss.angle, pixel.V(ss.power*1, 0), "img/shots/shot1.png"), 5}
	return *shot.SpaceObject
}
