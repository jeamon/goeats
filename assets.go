package main

import (
	"embed"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets/*/*.png
//go:embed assets/*/*/*/*.png
var pictures embed.FS

//go:embed assets/sounds/*/*.wav
var sounds embed.FS

func isWav(path string) bool {
	return filepath.Ext(path) == ".wav"
}

func loadPictures(dir string, m *[]item) {
	entries, err := pictures.ReadDir(dir)
	checkerr(err)
	for _, e := range entries {
		name, _, found := strings.Cut(e.Name(), "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		imgBytes, err := pictures.ReadFile(path.Join(dir, e.Name()))
		checkerr(err)
		rImg := rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
		*m = append(*m, item{name, rl.LoadTextureFromImage(rImg)})
	}
}

func loadFaces(faces map[string]rl.Texture2D) {
	imgByte, err := pictures.ReadFile("assets/faces/up.png")
	checkerr(err)
	faces["up"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = pictures.ReadFile("assets/faces/down.png")
	checkerr(err)
	faces["down"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = pictures.ReadFile("assets/faces/left.png")
	checkerr(err)
	faces["left"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = pictures.ReadFile("assets/faces/right.png")
	checkerr(err)
	faces["right"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))
}

func getImageFromPictures(path string) *rl.Image {
	imgByte, err := pictures.ReadFile(path)
	checkerr(err)
	return rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte)))
}

func loadSounds(dir string, m map[string]rl.Sound) {
	entries, err := sounds.ReadDir(dir)
	checkerr(err)
	for _, e := range entries {
		if !isWav(e.Name()) {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".wav")
		wavBytes, err := sounds.ReadFile(path.Join(dir, e.Name()))
		checkerr(err)
		rWav := rl.LoadWaveFromMemory(".wav", wavBytes, int32(len(wavBytes)))
		m[name] = rl.LoadSoundFromWave(rWav)
	}
}

// loadAssets - Load resources
func (g *Game) loadAssets() {
	// load pictures
	loadPictures("assets/fruits", &g.fruits)
	loadPictures("assets/vegetables", &g.vegetables)
	loadPictures("assets/donuts", &g.donuts)
	loadPictures("assets/lives", &g.lives)
	loadPictures("assets/balls", &g.balls)
	loadFaces(g.faces)

	// load fruits - vegetables - donuts - actions audio
	loadSounds("assets/sounds/fruits", g.sounds)
	loadSounds("assets/sounds/vegetables", g.sounds)
	loadSounds("assets/sounds/donuts", g.sounds)
	loadSounds("assets/sounds/actions", g.actions)

	// load boy 4D sprites idle and run positions
	var rImg *rl.Image
	for i := 0; i <= 5; i++ {
		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/back/%d.png", i+1))
		g.sprite.idle[Back] = append(g.sprite.idle[Back], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/front/%d.png", i+1))
		g.sprite.idle[Front] = append(g.sprite.idle[Front], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/left/%d.png", i+1))
		g.sprite.idle[Left] = append(g.sprite.idle[Left], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/right/%d.png", i+1))
		g.sprite.idle[Right] = append(g.sprite.idle[Right], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/back/%d.png", i+1))
		g.sprite.run[Back] = append(g.sprite.run[Back], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/front/%d.png", i+1))
		g.sprite.run[Front] = append(g.sprite.run[Front], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/left/%d.png", i+1))
		g.sprite.run[Left] = append(g.sprite.run[Left], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/right/%d.png", i+1))
		g.sprite.run[Right] = append(g.sprite.run[Right], rl.LoadTextureFromImage(rImg))
	}
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
