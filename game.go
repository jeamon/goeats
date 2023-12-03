package main

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type item struct {
	name    string
	picture rl.Texture2D
}

type Game struct {
	num           int
	faces         map[string]rl.Texture2D
	actions       map[string]rl.Sound
	fruits        []item
	vegetables    []item
	donuts        []item
	lives         []item
	sounds        map[string]rl.Sound
	walker        Walker
	foods         []*Food
	score         Score
	framesCounter int
}

func (g *Game) Init() {
	g.num = 11 // 10 foods + 1 life
	g.foods = make([]*Food, g.num)
	for i := 0; i < g.num-1; i++ {
		g.foods[i] = &Food{change: true, kind: kind(i % 3)}
	}
	// append one life item to the foods list
	// g.foods = append(g.foods, &Food{change: true, kind: L})
	g.foods[10] = &Food{change: true, kind: L}

	g.faces = make(map[string]rl.Texture2D, 4)
	g.vegetables = make([]item, 0, 68)
	g.fruits = make([]item, 0, 44)
	g.donuts = make([]item, 0, 12)
	g.lives = make([]item, 0, 4)
	g.actions = make(map[string]rl.Sound, 3)
	g.sounds = make(map[string]rl.Sound)

	g.walker.velocity = 4
	g.walker.size = rl.NewVector2(48/2, 48/2)
	g.walker.position = rl.NewVector2(-2, 0)
	g.walker.srcRec = rl.NewRectangle(0, 0, g.walker.size.X, g.walker.size.Y)
	g.walker.dstRec = rl.NewRectangle(0, 0, g.walker.size.X*3, g.walker.size.Y*3)
}

// Draw game textures
func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	g.score.draw()
	for _, food := range g.foods {
		food.draw()
	}
	g.walker.draw()
	rl.EndDrawing()
}

func (g *Game) collision(food *Food) bool {
	centerX := g.walker.dstRec.X - g.walker.position.X + g.walker.size.X*3/2
	centerY := g.walker.dstRec.Y - g.walker.position.Y + g.walker.size.Y*3/2
	radius := (g.walker.size.X * 3) * float32(math.Sqrt(2)) / 2
	return rl.CheckCollisionPointCircle(food.position, rl.NewVector2(centerX, centerY), radius)
}

func (g *Game) update() {
	for _, food := range g.foods {
		food.change = false
		if g.collision(food) {
			if food.kind == L {
				rl.PlaySound(g.actions["life"])
			} else {
				rl.PlaySound(g.actions["eat"])
				rl.PlaySound(g.sounds[food.name])
			}
			g.score.update(food.kind)
			food.change = true
		}
	}
	g.randomize()
}

func (g *Game) checkExpire() {
	for _, food := range g.foods {
		food.change = false
		if time.Now().Unix() >= food.expire {
			food.change = true
			continue
		}
	}
	g.randomize()
}

func (g *Game) randomize() {
	for _, food := range g.foods {
		if !food.change {
			continue
		}
		switch food.kind {
		case F:
			food.randomize(&(g.fruits))
		case V:
			food.randomize(&(g.vegetables))
		case D:
			food.randomize(&(g.donuts))
		case L:
			food.randomize(&(g.lives))
		}
	}
}
