package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Walker type
type Walker struct {
	moving    bool
	frames    int
	direction rl.Texture2D
	velocity  float32
	position  rl.Vector2
	size      rl.Vector2
	srcRec    rl.Rectangle
	dstRec    rl.Rectangle
}

func (w *Walker) draw() {
	rl.DrawTexturePro(w.direction, w.srcRec, w.dstRec, w.position, 0, rl.White)
	// rec := rl.NewRectangle(w.dstRec.X-w.position.X, w.srcRec.Y-w.position.Y, w.size.X*3, w.size.Y*3)
	// rl.DrawRectangleLinesEx(rec, 3, rl.Red)
	//	rl.DrawCircleLines(int32(rec.X+rec.Width/2), int32(rec.Y+rec.Height/2), rec.Width*float32(math.Sqrt(2))/2, rl.Red)
	centerX := int32(w.dstRec.X - w.position.X + w.size.X*3/2)
	centerY := int32(w.dstRec.Y - w.position.Y + w.size.Y*3/2)
	radius := (w.size.X * 3) * float32(math.Sqrt(2)) / 2
	rl.DrawCircleLines(centerX, centerY, radius, rl.Red)
}
