package core

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type enemy interface {
	speed(float32)
	draw()
	collision(center rl.Vector2, radius float32) bool // check collision with player circle
}

type ball struct {
	picture  rl.Texture2D
	speedX   float32
	speedY   float32
	position rl.Vector2
}

func (b *ball) draw() {
	b.update()
	rl.DrawTextureV(b.picture, b.position, rl.White)
}

func (b *ball) update() {
	b.position.X += b.speedX
	b.position.Y += b.speedY
	if b.position.Y <= 0 || b.position.Y+float32(b.picture.Height) >= float32(ScreenH) {
		b.speedY *= -1
	}
	if b.position.X <= 0 || b.position.X+float32(b.picture.Width) >= float32(ScreenW) {
		b.speedX *= -1
	}
}

func (b *ball) speed(s float32) {
	if b.speedX >= 0 {
		b.speedX = s
	} else {
		b.speedX = -s
	}
	if b.speedY >= 0 {
		b.speedY = s
	} else {
		b.speedY = -s
	}
}

func (b *ball) collision(center rl.Vector2, radius float32) bool {
	return rl.CheckCollisionCircles(
		rl.Vector2{X: b.position.X + float32(b.picture.Width)/2, Y: b.position.Y + float32(b.picture.Height)/2},
		(float32(b.picture.Width)*float32(math.Sqrt(2))/2)-9,
		center, radius)
}
