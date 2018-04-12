package main

import "github.com/faiface/pixel"

type Particle interface {
	draw()
	collides() bool
}

type ParticleType struct {
	pos   pixel.Matrix
	speed pixel.Matrix
}

type ShipShot struct {
	pos   pixel.Matrix
	speed pixel.Matrix
	image string
}

type Asteroid struct {
	pos   pixel.Matrix
	speed pixel.Matrix
	image string
}
