package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Food type
type Food struct {
	picture  rl.Texture2D
	position rl.Vector2
	name     string
	change   bool
	expire   int64
	kind     kind
}

func (f *Food) randomize(items *[]item) {
	f.position = rl.NewVector2(
		float32(rl.GetRandomValue(1, (screenW/cellsize)-2))*cellsize,
		float32(rl.GetRandomValue(1, (screenH/cellsize)-2))*cellsize,
	)
	item := (*items)[rl.GetRandomValue(0, int32(len(*items)-1))]
	f.picture = item.picture
	f.name = item.name
	f.expire = time.Now().Add(30 * time.Second).Unix()
}

func (f *Food) draw() {
	if f.kind == L {
		rl.DrawTextureEx(f.picture, f.position, float32(rl.GetTime()*180), 1.0, rl.White)
	} else if f.kind == V {
		rl.DrawTextureEx(f.picture, f.position, float32(rl.GetTime()*90), 1.0, rl.White)
	} else if f.kind == F {
		rl.DrawTextureEx(f.picture, f.position, float32(rl.GetTime()*-90), 1.0, rl.White)
	} else if f.kind == D {
		rl.DrawTextureEx(f.picture, f.position, -float32(rl.GetTime()*90), 1.0, rl.White)
	} else {
		rl.DrawTextureV(f.picture, f.position, rl.White)
	}
}

func (f *Food) collision(center rl.Vector2, radius float32) bool {
	return rl.CheckCollisionPointCircle(f.position, center, radius)
}
