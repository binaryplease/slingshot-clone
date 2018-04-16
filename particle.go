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
	pos     pixel.Vec
	speed   pixel.Vec
	angle   float64
	timeout int
	image   string
}

type SpaceObject interface {
	getPos() pixel.Vec
	getAngle() float64
	getImage() string
}

type Shot struct {
	pos     pixel.Vec
	speed   pixel.Vec
	timeout int
	image   string
}

func (s *Shot) move() {
	spawnTime := time.Now()
	go func() {
		for int(time.Now().Sub(spawnTime)) < s.timeout {
			s.pos = s.pos.Add(s.speed)
			time.Sleep(100 * time.Millisecond) //TODO implement proper fps
		}
	}()
}

type Asteroid struct {
	ParticleType
}

// Move asteroid until timeout is reached
func (a *Asteroid) move() {
	spawnTime := time.Now()
	go func() {
		for int(time.Now().Sub(spawnTime)) < a.timeout {
			a.pos = a.pos.Add(a.speed)
		}
	}()
}
