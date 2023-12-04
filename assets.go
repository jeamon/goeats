package main

import (
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	f.Close()

	g.faces["up"] = rl.LoadTexture("assets/faces/up.png")
	g.faces["down"] = rl.LoadTexture("assets/faces/down.png")
	g.faces["left"] = rl.LoadTexture("assets/faces/left.png")
	g.faces["right"] = rl.LoadTexture("assets/faces/right.png")

	g.actions["eat"] = rl.LoadSound("assets/sounds/actions/eat.wav")
	g.actions["life"] = rl.LoadSound("assets/sounds/actions/life.mp3")
	g.actions["level"] = rl.LoadSound("assets/sounds/actions/level.mp3")
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

	for _, s := range g.sounds {
		rl.UnloadSound(s)
	}

	for _, s := range g.actions {
		rl.UnloadSound(s)
	}
}
