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
		g.walker.direction = g.faces["right"]
		if g.walker.position.X-g.walker.velocity > -(screenW - g.walker.size.X*3) {
			g.walker.position.X -= g.walker.velocity
		}
		g.walker.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyLeft) || gamePadButtonPressed() == LT {
		g.walker.direction = g.faces["left"]
		if g.walker.position.X+g.walker.velocity < 0 {
			g.walker.position.X += g.walker.velocity
		}
		g.walker.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyUp) || gamePadButtonPressed() == UP {
		g.walker.direction = g.faces["up"]
		if g.walker.position.Y+g.walker.velocity < 0 {
			g.walker.position.Y += g.walker.velocity
		}
		g.walker.moving = true
		g.update()
	}

	if rl.IsKeyDown(rl.KeyDown) || gamePadButtonPressed() == DN {
		g.walker.direction = g.faces["down"]
		if g.walker.position.Y-g.walker.velocity > -(screenH - g.walker.size.Y*3) {
			g.walker.position.Y -= g.walker.velocity
		}
		g.walker.moving = true
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
