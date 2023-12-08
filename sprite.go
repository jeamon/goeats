package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type direction string

const (
	Front direction = "front"
	Back  direction = "back"
	Left  direction = "left"
	Right direction = "Right"
)

const (
	spriteW = 80
	spriteH = 80
)

type Sprite struct {
	run      map[direction][]rl.Texture2D
	idle     map[direction][]rl.Texture2D
	moving   bool
	face     direction
	velocity float32
	speed    float32
	position rl.Vector2
}

func (s *Sprite) draw() {
	if s.moving {
		txt := s.run[s.face][int(s.velocity)]
		rl.DrawTexture(txt, int32(s.position.X), int32(s.position.Y), rl.White)
	} else {
		txt := s.idle[s.face][int(rl.GetTime())%6]
		rl.DrawTexture(txt, int32(s.position.X), int32(s.position.Y), rl.White)
	}

	centerX := int32(s.position.X + spriteW/2)
	centerY := int32(s.position.Y + spriteH/2)
	radius := float32(spriteW) * float32(math.Sqrt(2)) / 2
	rl.DrawCircleLines(centerX, centerY, radius, rl.Red)
}
