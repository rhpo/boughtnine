package levels

import (
	"boughtnine/life"
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

var (
	player       *life.Shape
	playerEntity *PlayerEntity
	world        *life.World
)

var One life.Level = life.Level{

	MapItems: life.MapItems{
		"#": func(position life.Vector2, width float64, height float64) {
			s := life.NewShape(&life.ShapeProps{
				Name:         "wall",
				Type:         life.ShapeRectangle,
				Pattern:      life.PatternColor,
				Physics:      false,
				IsBody:       false,
				Background:   color.Black,
				X:            position.X,
				Y:            position.Y,
				Width:        width,
				Height:       height,
				Friction:     0.5,
				Rebound:      0,
				RotationLock: true,
			})

			world.Register(s)
		},

		"'": func(position life.Vector2, width float64, height float64) {
			s := life.NewShape(&life.ShapeProps{
				Tag:          "ground",
				Type:         life.ShapeRectangle,
				Pattern:      life.PatternColor,
				Physics:      false,
				IsBody:       false,
				Background:   color.Black,
				X:            position.X,
				Y:            position.Y,
				Width:        width,
				Height:       height,
				Friction:     0.5,
				Rebound:      0,
				RotationLock: true,
			})

			world.Register(s)
		},

		"F": func(position life.Vector2, width float64, height float64) {
			s := life.NewShape(&life.ShapeProps{
				Type:       life.ShapeRectangle,
				Pattern:    life.PatternColor,
				Background: color.RGBA{R: 0, G: 255, B: 0, A: 255}, // Green color for finish line
				X:          position.X,
				Y:          position.Y,
				Width:      width,
				Height:     height,
				OnCollisionFunc: func(who *life.Shape) {
					if who == player {
						// Play level complete sound
						world.PlaySound("level_complete")
						world.NextLevel()
					}
				},
			})

			world.Register(s)
		},

		"@": func(position life.Vector2, width float64, height float64) {
			player.SetX(position.X)
			player.SetY(position.Y)
		},

		"o": func(position life.Vector2, width float64, height float64) {
			s := life.NewShape(&life.ShapeProps{
				Type:         life.ShapeCircle,
				Pattern:      life.PatternColor,
				Background:   color.RGBA{R: 255, G: 215, B: 0, A: 255}, // Gold color for collectible
				X:            position.X + width/2,
				Y:            position.Y + height/2,
				Radius:       width / 2,
				Physics:      true,
				IsBody:       true,
				Friction:     0.5,
				Mass:         100,
				Rebound:      0.5,
				RotationLock: false,
				OnCollisionFunc: func(who *life.Shape) {
					if who == player {
						world.PlaySound("jump") // Play jump sound on collection
					}
				},
			})

			world.Register(s)
		},
	},

	Init: func(world_ *life.World) {
		world = world_

		// Load audio files
		// Note: You'll need to add actual audio files to your assets folder
		world.LoadSound("jump", assets, "assets/sounds/jump.wav")
		world.LoadSound("level_complete", assets, "assets/sounds/complete.wav")
		world.LoadMusic("background", assets, "assets/music/background.mp3")

		playerEntity = NewPlayerEntity(world, assets)
		player = playerEntity.Shape

		// Play background music
		// world.PlayMusic("background")
	},

	Tick: func(ld life.LoopData) {
		playerEntity.Update(ld)
	},

	Render: func(screen *ebiten.Image) {
		life.DrawText(screen, &life.TextProps{
			Text:  "Press Space to jump, Left/Right to move",
			X:     player.X + 10,
			Y:     player.Y - 15,
			Color: color.White,
		})
	},

	Map: life.Map{
		"#############################",
		"#                           #",
		"#                           #",
		"#  @ o                      #",
		"''''''F                     #",
		"#                           #",
		"#                           #",
		"#                           #",
		"'''''''''''''''''''''''''''''",
	},
}
