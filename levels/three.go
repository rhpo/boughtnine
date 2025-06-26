package levels

import (
	"boughtnine/life"
	"image/color"
)

var Three life.Level = life.Level{

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
				Background:   color.RGBA{R: 200, G: 200, B: 200, A: 255}, // Light gray for ground
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
				Background: color.RGBA{R: 255, G: 0, B: 0, A: 255}, // Red color for finish line
				X:          position.X,
				Y:          position.Y,
				Width:      width,
				Height:     height,
				OnCollisionFunc: func(who *life.Shape) {
					if who == player {
						// Use the safe switching method instead of direct SelectLevel
						world.SwitchToLevel(0) // Go back to level 1 (index 0)
					}
				},
			})

			world.Register(s)
		},

		"@": func(position life.Vector2, width float64, height float64) {
			player.SetX(position.X)
			player.SetY(position.Y)
		},
	},

	Init: func(world_ *life.World) {
		world = world_

		playerEntity = NewPlayerEntity(world, assets)
		player = playerEntity.Shape
	},

	Tick: func(ld life.LoopData) {
		if playerEntity != nil {
			playerEntity.Update(ld)
		}
	},

	Map: life.Map{
		"#############################",
		"#      @                    #",
		"#      '''                  #",
		"#                           #",
		"#  '''                      #",
		"#          '''              #",
		"#                           #",
		"#             FFF           #",
		"'''''''''''''''''''''''''''''",
	},
}
