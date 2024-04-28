package main

import (
	"embed"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jeamon/goeats/core"
)

//go:embed assets/*/*.png
//go:embed assets/*/*/*.png
//go:embed assets/*/*/*/*.png
var picturesFs embed.FS

//go:embed assets/sounds/*/*.wav
var soundsFs embed.FS

// loadsettings loads environment variables to set screen width and height.
// min value for width is 800 and for height is 500
func loadsettings() {
	w := os.Getenv("GOEATS_SCREEN_WIDTH")
	h := os.Getenv("GOEATS_SCREEN_HEIGHT")
	if w != "" {
		if v, err := strconv.Atoi(w); err == nil && v >= 800 {
			core.ScreenW = int32(v)
		}
	}
	if h != "" {
		if v, err := strconv.Atoi(h); err == nil && v >= 500 {
			core.ScreenH = int32(v)
		}
	}
}

func main() {
	loadsettings()
	core.PicturesFs = picturesFs
	core.SoundsFs = soundsFs

	game := core.Game{}
	game.Init()
	rl.InitWindow(core.ScreenW, core.ScreenH, "Go & Eats")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)
	game.LoadAssets()
	game.Score.Sound = game.Actions["level"]
	game.Randomize()
	game.Sprite.Draw()
	for !rl.WindowShouldClose() {
		game.Draw()
		game.Sprite.Moving = false
		game.Controls()
		game.CheckExpire()

		if game.Sprite.Moving {
			game.Sprite.Velocity += 0.3
			if game.Sprite.Velocity >= 6 {
				game.Sprite.Velocity = 0
			}
		}
	}

	game.Unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
