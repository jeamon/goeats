package main

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

func (g *Game) controls() {
	if rl.IsKeyDown(rl.KeyRight) || gamePadButtonPressed() == RT {
		g.sprite.face = Right
		if g.sprite.position.X+g.sprite.speed <= (float32(screenW) - spriteW) {
			g.sprite.position.X += g.sprite.speed
		}
		g.sprite.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyLeft) || gamePadButtonPressed() == LT {
		g.sprite.face = Left
		if g.sprite.position.X-g.sprite.speed >= 0 {
			g.sprite.position.X -= g.sprite.speed
		}
		g.sprite.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyUp) || gamePadButtonPressed() == UP {
		g.sprite.face = Back
		if g.sprite.position.Y-g.sprite.speed >= 0 {
			g.sprite.position.Y -= g.sprite.speed
		}
		g.sprite.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyDown) || gamePadButtonPressed() == DN {
		g.sprite.face = Front
		if g.sprite.position.Y+g.sprite.speed <= float32(screenH)-spriteH {
			g.sprite.position.Y += g.sprite.speed
		}
		g.sprite.moving = true
		g.update()
	}
}

func gamePadButtonPressed() gamepadkey {
	if rl.IsGamepadAvailable(gamepad) {
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
