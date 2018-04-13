package main

import (
	"github.com/faiface/pixel"
	"time"
)

type SlingshotCamera struct {
	cam          pixel.Matrix
	camPos       pixel.Vec
	camSpeed     float64
	camZoom      float64
	camZoomSpeed float64
	last         time.Time
}

func NewSlingshotCamera() *SlingshotCamera {

	var (
		camZoom      = 1.0
		camZoomSpeed = 1.2
		camPos       = pixel.ZV
		camSpeed     = 500.0
	)

	cam := &SlingshotCamera{pixel.IM, camPos, camSpeed, camZoom, camZoomSpeed, time.Now()}
	return cam
}
