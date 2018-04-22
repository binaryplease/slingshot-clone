package main

import "github.com/faiface/pixel"
import "time"

type SpaceObject struct {
	pos         pixel.Vec
	angle       float64
	speed       pixel.Vec
	image       string
	spawnTime   time.Time
	updateMilis time.Duration
}

func NewSpaceObject(pos pixel.Vec, angle float64, speed pixel.Vec, image string) *SpaceObject {
	return &SpaceObject{pos, angle, speed, image, time.Now(), 10 * time.Millisecond}
}

type Shot struct {
	*SpaceObject
	timeout int
}

func (so *SpaceObject) update() {
	so.pos = so.pos.Add(so.speed.Rotated(so.angle))
}

func (so SpaceObject) distanceTo(other SpaceObject) {

}

func (so SpaceObject) collides(other SpaceObject) {

}
