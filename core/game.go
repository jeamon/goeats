package core

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const cellsize = 40

var (
	ScreenW int32 = 1280
	ScreenH int32 = 768
)

type item struct {
	name    string
	picture rl.Texture2D
}

type Game struct {
	num        int
	faces      map[string]rl.Texture2D
	Actions    map[string]rl.Sound
	fruits     []item
	vegetables []item
	donuts     []item
	lives      []item
	balls      []item // balls to use as enemies
	sounds     map[string]rl.Sound
	Sprite     Sprite
	foods      []*Food
	enemies    []enemy
	Score      Score
}

func (g *Game) Init() {
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
	g.Actions = make(map[string]rl.Sound, 3)
	g.sounds = make(map[string]rl.Sound)

	g.Sprite.Velocity = 0.0
	g.Sprite.speed = 7
	g.Sprite.Moving = false
	g.Sprite.position = rl.NewVector2(2, float32(ScreenH/2))
	g.Sprite.face = Right
	g.Sprite.idle = make(map[direction][]rl.Texture2D, 4)
	g.Sprite.run = make(map[direction][]rl.Texture2D, 4)
	g.Sprite.radius = float32(spriteW) * float32(math.Sqrt(2)) / 2

	g.balls = make([]item, 0, 4)
}

// Draw game textures
func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Beige)
	g.Score.draw()
	g.drawFoods()
	g.drawEnemies()
	g.Sprite.Draw()
	rl.EndDrawing()
}

func (g *Game) drawFoods() {
	for _, food := range g.foods {
		food.draw()
	}
}

func (g *Game) drawEnemies() {
	center := rl.Vector2{X: g.Sprite.position.X + spriteW/2, Y: g.Sprite.position.Y + spriteH/2}
	// add 1 ball-based enemy each 3rd level with a max of available balls
	if g.Score.level > 0 && g.Score.level%3 == 0 && (g.Score.level/3) != len(g.enemies) && len(g.enemies) < len(g.balls) {
		g.addEnemy()
	}

	for _, e := range g.enemies {
		// increase enemy speed with a step of half the level + 1 on both axises while
		// making sure it is done until level 12 rather than 3*len(g.balls) in order to
		// cap the speed to a max of 7.
		if g.Score.level <= 12 {
			e.speed(float32(g.Score.level/2 + 1))
		}
		e.draw()

		if e.collision(center, g.Sprite.radius) {
			rl.PlaySound(g.Actions["hurt"])
			g.Score.lives--
		}
	}
}

func (g *Game) addEnemy() {
	e := &ball{}
	item := g.balls[len(g.enemies)]
	e.picture = item.picture
	g.enemies = append(g.enemies, e)
}

func (g *Game) update() {
	center := rl.Vector2{X: g.Sprite.position.X + spriteW/2, Y: g.Sprite.position.Y + spriteH/2}
	for _, food := range g.foods {
		food.change = false
		if food.collision(center, g.Sprite.radius) {
			if food.kind == L {
				rl.PlaySound(g.Actions["life"])
			} else {
				rl.PlaySound(g.Actions["eat"])
				rl.PlaySound(g.sounds[food.name])
			}
			g.Score.update(food.kind)
			food.change = true
		}
	}
	g.Randomize()
}

func (g *Game) CheckExpire() {
	for _, food := range g.foods {
		food.change = false
		if time.Now().Unix() >= food.expire {
			food.change = true
			continue
		}
	}
	g.Randomize()
}

func (g *Game) Randomize() {
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
