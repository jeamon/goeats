package core

import (
	"embed"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	PicturesFs embed.FS
	SoundsFs   embed.FS
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isWav(path string) bool {
	return filepath.Ext(path) == ".wav"
}

func loadPictures(dir string, m *[]item) {
	entries, err := PicturesFs.ReadDir(dir)
	checkerr(err)
	for _, e := range entries {
		name, _, found := strings.Cut(e.Name(), "_")
		if !found {
			name = strings.TrimSuffix(name, ".png")
		}
		imgBytes, err := PicturesFs.ReadFile(path.Join(dir, e.Name()))
		checkerr(err)
		rImg := rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
		*m = append(*m, item{name, rl.LoadTextureFromImage(rImg)})
	}
}

func loadFaces(faces map[string]rl.Texture2D) {
	imgByte, err := PicturesFs.ReadFile("assets/faces/up.png")
	checkerr(err)
	faces["up"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = PicturesFs.ReadFile("assets/faces/down.png")
	checkerr(err)
	faces["down"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = PicturesFs.ReadFile("assets/faces/left.png")
	checkerr(err)
	faces["left"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))

	imgByte, err = PicturesFs.ReadFile("assets/faces/right.png")
	checkerr(err)
	faces["right"] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte))))
}

func getImageFromPictures(path string) *rl.Image {
	imgByte, err := PicturesFs.ReadFile(path)
	checkerr(err)
	return rl.LoadImageFromMemory(".png", imgByte, int32(len(imgByte)))
}

func loadSounds(dir string, m map[string]rl.Sound) {
	entries, err := SoundsFs.ReadDir(dir)
	checkerr(err)
	for _, e := range entries {
		if !isWav(e.Name()) {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".wav")
		wavBytes, err := SoundsFs.ReadFile(path.Join(dir, e.Name()))
		checkerr(err)
		rWav := rl.LoadWaveFromMemory(".wav", wavBytes, int32(len(wavBytes)))
		m[name] = rl.LoadSoundFromWave(rWav)
	}
}

// LoadAssets - Load resources
func (g *Game) LoadAssets() {
	// load pictures
	loadPictures("assets/foods/fruits", &g.fruits)
	loadPictures("assets/foods/vegetables", &g.vegetables)
	loadPictures("assets/foods/donuts", &g.donuts)
	loadPictures("assets/foods/lives", &g.lives)
	loadPictures("assets/balls", &g.balls)
	loadFaces(g.faces)

	// load fruits - vegetables - donuts - actions audio
	loadSounds("assets/sounds/fruits", g.sounds)
	loadSounds("assets/sounds/vegetables", g.sounds)
	loadSounds("assets/sounds/donuts", g.sounds)
	loadSounds("assets/sounds/actions", g.Actions)

	// load boy 4D sprites idle and run positions
	var rImg *rl.Image
	for i := 0; i <= 5; i++ {
		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/back/%d.png", i+1))
		g.Sprite.idle[Back] = append(g.Sprite.idle[Back], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/front/%d.png", i+1))
		g.Sprite.idle[Front] = append(g.Sprite.idle[Front], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/left/%d.png", i+1))
		g.Sprite.idle[Left] = append(g.Sprite.idle[Left], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/idle/right/%d.png", i+1))
		g.Sprite.idle[Right] = append(g.Sprite.idle[Right], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/back/%d.png", i+1))
		g.Sprite.run[Back] = append(g.Sprite.run[Back], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/front/%d.png", i+1))
		g.Sprite.run[Front] = append(g.Sprite.run[Front], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/left/%d.png", i+1))
		g.Sprite.run[Left] = append(g.Sprite.run[Left], rl.LoadTextureFromImage(rImg))

		rImg = getImageFromPictures(fmt.Sprintf("assets/boy/run/right/%d.png", i+1))
		g.Sprite.run[Right] = append(g.Sprite.run[Right], rl.LoadTextureFromImage(rImg))
	}
}

// Unload - Unload resources
func (g *Game) Unload() {
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

	for _, rtxt := range g.Sprite.idle {
		for _, txt := range rtxt {
			rl.UnloadTexture(txt)
		}
	}

	for _, rtxt := range g.Sprite.run {
		for _, txt := range rtxt {
			rl.UnloadTexture(txt)
		}
	}

	for _, s := range g.sounds {
		rl.UnloadSound(s)
	}

	for _, s := range g.Actions {
		rl.UnloadSound(s)
	}

	for _, b := range g.balls {
		rl.UnloadTexture(b.picture)
	}
}
