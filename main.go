package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type kind uint8

const (
	V kind = 0 // veggie item
	F kind = 1 // fruit item
	D kind = 2 // donut item
	L kind = 3 // life item
)

const (
	cellsize = 40
	screenW  = 960
	screenH  = 640
)

var gamepad int32 = 0 // gamepad to track

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	game := Game{}
	game.Init()
	rl.InitWindow(screenW, screenH, "Go & Eats")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)
	game.load()
	game.score.sound = game.actions["level"]
	game.randomize()

	game.walker.direction = game.faces["right"]
	game.walker.draw()
	for !rl.WindowShouldClose() {
		game.draw()
		game.walker.moving = false
		game.framesCounter++
		game.controls()
		game.checkExpire()

		if game.walker.moving {
			if game.framesCounter%8 == 1 {
				game.walker.frames++
				game.walker.frames %= 2
			}
		}

		game.walker.srcRec.X = game.walker.srcRec.Width * float32(game.walker.frames)
	}

	game.unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
