package main

import (
	"boughtnine/entities"
	"boughtnine/levels"
	"boughtnine/life"
	"embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

func main() {
	world := entities.NewWorld()
	game := life.NewGame(world)

	world.Levels = []life.Level{
		levels.One,
		levels.Two,
		levels.Three,
	}

	// Test audio with a generated tone
	world.AudioManager.CreateTestTone("test_beep", 440.0, 500*time.Millisecond)

	// Test playing the tone when 'T' is pressed
	world.On(life.EventMouseDown, func(data interface{}) {
		if err := world.PlaySound("test_beep"); err != nil {
			// Handle error - but we can't print it
		}
	})

	// Also test with keyboard
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			if world.IsKeyPressed(ebiten.KeyT) {
				world.PlaySound("test_beep")
				time.Sleep(600 * time.Millisecond) // Prevent spam
			}
		}
	}()

	game.Run()
}
