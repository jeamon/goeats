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

const cellsize = 40

var (
	screenW int32 = 1280
	screenH int32 = 768
)

var gamepad int32 = 0 // gamepad to track

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	game := Game{}
	game.init()
	rl.InitWindow(screenW, screenH, "Go & Eats")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)
	game.load()
	game.score.sound = game.actions["level"]
	game.randomize()
	game.sprite.draw()
	for !rl.WindowShouldClose() {
		game.draw()
		game.sprite.moving = false
		game.controls()
		game.checkExpire()

		if game.sprite.moving {
			game.sprite.velocity += 0.3
			if game.sprite.velocity >= 6 {
				game.sprite.velocity = 0
			}
		}
	}

	game.unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
