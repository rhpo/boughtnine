package entities

import (
	"boughtnine/life"
	"image/color"
)

func NewWorld() *life.World {
	world := life.NewWorld(&life.WorldProps{
		Width:         1920,
		Height:        1080,
		G:             life.NewVector2(0, 11),
		Background:    color.RGBA{R: 65, G: 105, B: 225, A: 255},
		Paused:        false,
		HasLimits:     false,
		AirResistance: 0.0,
		Title:         "Bought Nine Lifes",
	})

	world.CreateBorders()
	return world
}
