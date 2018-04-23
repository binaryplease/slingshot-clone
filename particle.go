package main

import "math"
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

func (so SpaceObject) distanceTo(other SpaceObject) float64 {
	return math.Sqrt(math.Pow(so.pos.X-other.pos.X, 2) - math.Pow(so.pos.Y-other.pos.Y, 2))
}

func (so SpaceObject) collides(other SpaceObject) bool {
	return so.distanceTo(other) <= so.size()+other.size()
}

func (so SpaceObject) size() float64 {
	pic, err := loadPicture(so.image)

	if err != nil {
		panic(err)
	}

	if pic.Bounds().H() > pic.Bounds().W() {
		return pic.Bounds().H()
	}
	return pic.Bounds().W()

}
