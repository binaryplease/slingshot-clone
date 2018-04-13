package main

import "github.com/faiface/pixel"
import "time"

type Particle interface {
	// Spawns particle
	draw()

	//Checks for collision
	collides(Particle) bool
}

type ParticleType struct {
	pos     pixel.Matrix
	speed   pixel.Vec
	timeout int
	image   string
}

type Asteroid struct {
	ParticleType
}

// Move asteroid until timeout is reached
func (a *Asteroid) move() {
	spawnTime := time.Now()
	go func() {
		for int(time.Now().Sub(spawnTime)) < a.timeout {
			a.pos = a.pos.Moved(a.speed)
		}
	}()
}
