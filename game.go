package main

import (
	"math"
	"os"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type item struct {
	name    string
	picture rl.Texture2D
}

type Game struct {
	num        int
	faces      map[string]rl.Texture2D
	actions    map[string]rl.Sound
	fruits     []item
	vegetables []item
	donuts     []item
	lives      []item
	balls      []item // balls to use as enemies
	sounds     map[string]rl.Sound
	sprite     Sprite
	foods      []*Food
	enemies    []enemy
	score      Score
}

// loadsettings loads environment variables to set screen width and height.
// min value for width is 800 and for height is 500
func loadsettings() {
	w := os.Getenv("GOEATS_SCREEN_WIDTH")
	h := os.Getenv("GOEATS_SCREEN_HEIGHT")
	if w != "" {
		if v, err := strconv.Atoi(w); err == nil && v >= 800 {
			screenW = int32(v)
		}
	}
	if h != "" {
		if v, err := strconv.Atoi(h); err == nil && v >= 500 {
			screenH = int32(v)
		}
	}
}

func (g *Game) init() {
	loadsettings()
	g.num = 11 // 10 foods + 1 life
	g.foods = make([]*Food, g.num)
	for i := 0; i < g.num-1; i++ {
		g.foods[i] = &Food{change: true, kind: kind(i % 3)}
	}
	// append 1 life items to the foods list
	g.foods[10] = &Food{change: true, kind: L}

	g.faces = make(map[string]rl.Texture2D, 4)
	g.vegetables = make([]item, 0, 98)
	g.fruits = make([]item, 0, 44)
	g.donuts = make([]item, 0, 12)
	g.lives = make([]item, 0, 5)
	g.actions = make(map[string]rl.Sound, 3)
	g.sounds = make(map[string]rl.Sound)

	g.sprite.velocity = 0.0
	g.sprite.speed = 7
	g.sprite.moving = false
	g.sprite.position = rl.NewVector2(2, float32(screenH/2))
	g.sprite.face = Right
	g.sprite.idle = make(map[direction][]rl.Texture2D, 4)
	g.sprite.run = make(map[direction][]rl.Texture2D, 4)

	g.balls = make([]item, 0, 4)
}

// Draw game textures
func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	g.score.draw()
	g.drawFoods()
	g.drawEnemies()
	g.sprite.draw()
	rl.EndDrawing()
}

func (g *Game) drawFoods() {
	for _, food := range g.foods {
		food.draw()
	}
}

func (g *Game) drawEnemies() {
	center := rl.Vector2{X: g.sprite.position.X + spriteW/2, Y: g.sprite.position.Y + spriteH/2}
	radius := float32(spriteW) * float32(math.Sqrt(2)) / 2
	// add 1 ball-based enemy each 3rd level with a max of available balls
	if g.score.level > 0 && g.score.level%3 == 0 && (g.score.level/3) != len(g.enemies) && len(g.enemies) < len(g.balls) {
		g.addEnemy()
	}

	for _, e := range g.enemies {
		if g.score.level <= 7 {
			e.speed(float32(g.score.level))
		}
		e.draw()

		if e.collision(center, radius) {
			rl.PlaySound(g.actions["hurt"])
			g.score.lives--
		}
	}
}

func (g *Game) addEnemy() {
	e := &ball{}
	item := g.balls[len(g.enemies)]
	e.picture = item.picture
	e.speedX = float32(g.score.level)
	e.speedY = float32(g.score.level)
	g.enemies = append(g.enemies, e)
}

func (g *Game) update() {
	center := rl.Vector2{X: g.sprite.position.X + spriteW/2, Y: g.sprite.position.Y + spriteH/2}
	radius := float32(spriteW) * float32(math.Sqrt(2)) / 2
	for _, food := range g.foods {
		food.change = false
		if food.collision(center, radius) {
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
