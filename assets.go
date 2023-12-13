package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets/faces/up.png
var upBytes []byte

//go:embed assets/faces/down.png
var downBytes []byte

func isPng(path string) bool {
	return filepath.Ext(path) == ".png"
}

func isMp3(path string) bool {
	return filepath.Ext(path) == ".mp3"
}

// Load - Load resources
func (g *Game) load() {
	f, err := os.Open("assets/fruits")
	checkerr(err)
	filenames, err := f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isPng(fn) {
			continue
		}
		name, _, found := strings.Cut(fn, "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		g.fruits = append(g.fruits, item{name, rl.LoadTexture("assets/fruits/" + fn)})
	}
	f.Close()

	f, err = os.Open("assets/vegetables")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isPng(fn) {
			continue
		}
		name, _, found := strings.Cut(fn, "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		g.vegetables = append(g.vegetables, item{name, rl.LoadTexture("assets/vegetables/" + fn)})
	}
	f.Close()

	f, err = os.Open("assets/donuts")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isPng(fn) {
			continue
		}
		name, _, found := strings.Cut(fn, "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		g.donuts = append(g.donuts, item{name, rl.LoadTexture("assets/donuts/" + fn)})
	}

	f, err = os.Open("assets/lives")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isPng(fn) {
			continue
		}
		name, _, found := strings.Cut(fn, "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		g.lives = append(g.lives, item{name, rl.LoadTexture("assets/lives/" + fn)})
	}

	// load foods sounds
	f, err = os.Open("assets/sounds/fruits")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isMp3(fn) {
			continue
		}
		name := strings.TrimSuffix(fn, ".mp3")
		g.sounds[name] = rl.LoadSound("assets/sounds/fruits/" + fn)
	}

	f, err = os.Open("assets/sounds/vegetables")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isMp3(fn) {
			continue
		}
		name := strings.TrimSuffix(fn, ".mp3")
		g.sounds[name] = rl.LoadSound("assets/sounds/vegetables/" + fn)
	}

	f, err = os.Open("assets/sounds/donuts")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isMp3(fn) {
			continue
		}
		name := strings.TrimSuffix(fn, ".mp3")
		g.sounds[name] = rl.LoadSound("assets/sounds/donuts/" + fn)
	}

	f, err = os.Open("assets/balls")
	checkerr(err)
	filenames, err = f.Readdirnames(0)
	checkerr(err)
	for _, fn := range filenames {
		if !isPng(fn) {
			continue
		}
		g.balls = append(g.balls, item{picture: rl.LoadTexture("assets/balls/" + fn)})
	}

	f.Close()

	rImage := rl.LoadImageFromMemory(".png", upBytes, int32(len(upBytes)))
	// g.faces["up"] = rl.LoadTexture("assets/faces/up.png")
	g.faces["up"] = rl.LoadTextureFromImage(rImage)

	g.faces["down"] = rl.LoadTexture("assets/faces/down.png")
	g.faces["left"] = rl.LoadTexture("assets/faces/left.png")
	g.faces["right"] = rl.LoadTexture("assets/faces/right.png")

	// load boy 4D sprites idle and run positions
	var img string
	for i := 0; i <= 5; i++ {
		img = fmt.Sprintf("assets/boy/idle/back/%d.png", i+1)
		g.sprite.idle[Back] = append(g.sprite.idle[Back], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/idle/front/%d.png", i+1)
		g.sprite.idle[Front] = append(g.sprite.idle[Front], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/idle/left/%d.png", i+1)
		g.sprite.idle[Left] = append(g.sprite.idle[Left], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/idle/right/%d.png", i+1)
		g.sprite.idle[Right] = append(g.sprite.idle[Right], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/run/back/%d.png", i+1)
		g.sprite.run[Back] = append(g.sprite.run[Back], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/run/front/%d.png", i+1)
		g.sprite.run[Front] = append(g.sprite.run[Front], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/run/left/%d.png", i+1)
		g.sprite.run[Left] = append(g.sprite.run[Left], rl.LoadTexture(img))

		img = fmt.Sprintf("assets/boy/run/right/%d.png", i+1)
		g.sprite.run[Right] = append(g.sprite.run[Right], rl.LoadTexture(img))
	}

	g.actions["eat"] = rl.LoadSound("assets/sounds/actions/eat.wav")
	g.actions["life"] = rl.LoadSound("assets/sounds/actions/life.wav")
	g.actions["level"] = rl.LoadSound("assets/sounds/actions/level.mp3")
	g.actions["hurt"] = rl.LoadSound("assets/sounds/actions/hurt.wav")
}

// Unload - Unload resources
func (g *Game) unload() {
	for _, fruit := range g.fruits {
		rl.UnloadTexture(fruit.picture)
	}

	for _, veg := range g.vegetables {
		rl.UnloadTexture(veg.picture)
	}

	for _, do := range g.donuts {
		rl.UnloadTexture(do.picture)
	}

	for _, fruit := range g.fruits {
		rl.UnloadTexture(fruit.picture)
	}

	for _, face := range g.faces {
		rl.UnloadTexture(face)
	}

	for _, rtxt := range g.sprite.idle {
		for _, txt := range rtxt {
			rl.UnloadTexture(txt)
		}
	}

	for _, rtxt := range g.sprite.run {
		for _, txt := range rtxt {
			rl.UnloadTexture(txt)
		}
	}

	for _, s := range g.sounds {
		rl.UnloadSound(s)
	}

	for _, s := range g.actions {
		rl.UnloadSound(s)
	}

	for _, b := range g.balls {
		rl.UnloadTexture(b.picture)
	}
}
