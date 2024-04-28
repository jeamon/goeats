package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type gamepadkey uint8

const (
	RT gamepadkey = 0
	LT gamepadkey = 1
	UP gamepadkey = 2
	DN gamepadkey = 3
	NN gamepadkey = 20
)

const gamepadID int32 = 0 // gamepad to track

func (g *Game) Controls() {
	if rl.IsKeyDown(rl.KeyRight) || gamePadButtonPressed() == RT {
		g.Sprite.face = Right
		if g.Sprite.position.X+g.Sprite.speed <= (float32(ScreenW) - spriteW) {
			g.Sprite.position.X += g.Sprite.speed
		}
		g.Sprite.Moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyLeft) || gamePadButtonPressed() == LT {
		g.Sprite.face = Left
		if g.Sprite.position.X-g.Sprite.speed >= 0 {
			g.Sprite.position.X -= g.Sprite.speed
		}
		g.Sprite.Moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyUp) || gamePadButtonPressed() == UP {
		g.Sprite.face = Back
		if g.Sprite.position.Y-g.Sprite.speed >= 0 {
			g.Sprite.position.Y -= g.Sprite.speed
		}
		g.Sprite.Moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyDown) || gamePadButtonPressed() == DN {
		g.Sprite.face = Front
		if g.Sprite.position.Y+g.Sprite.speed <= float32(ScreenH)-spriteH {
			g.Sprite.position.Y += g.Sprite.speed
		}
		g.Sprite.Moving = true
		g.update()
	}
}

func gamePadButtonPressed() gamepadkey {
	if rl.IsGamepadAvailable(gamepadID) {
		switch rl.GetGamepadButtonPressed() {
		case rl.GamepadButtonLeftFaceRight, rl.GamepadButtonRightFaceRight:
			return RT
		case rl.GamepadButtonLeftFaceLeft, rl.GamepadButtonRightFaceLeft:
			return LT
		case rl.GamepadButtonLeftFaceUp, rl.GamepadButtonRightFaceUp:
			return UP
		case rl.GamepadButtonLeftFaceDown, rl.GamepadButtonRightFaceDown:
			return DN
		}
	}
	return NN
}
