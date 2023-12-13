package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Score struct {
	fruits  int
	veggies int
	donuts  int
	lives   int
	level   int
	sound   rl.Sound
}

func (s *Score) draw() {
	f := fmt.Sprintf("Fruits: %d", s.fruits)
	v := fmt.Sprintf("Veggies: %d", s.veggies)
	d := fmt.Sprintf("Donuts: %d", s.donuts)
	lives := fmt.Sprintf("Lives: %d", s.lives)
	l := fmt.Sprintf("Level: %d", s.level)
	rl.DrawText(l, screenW-rl.MeasureText(l, 20)-7, 7, 20, rl.Red)
	rl.DrawText(lives, screenW-rl.MeasureText(d, 20)-7, 30, 20, rl.Yellow)
	rl.DrawText(d, screenW-rl.MeasureText(d, 20)-7, 53, 20, rl.DarkPurple)
	rl.DrawText(f, screenW-rl.MeasureText(f, 20)-7, 76, 20, rl.DarkPurple)
	rl.DrawText(v, screenW-rl.MeasureText(v, 20)-7, 99, 20, rl.DarkPurple)
}

func (s *Score) update(k kind) {
	switch k {
	case V:
		s.veggies++
	case F:
		s.fruits++
	case D:
		s.donuts++
	case L:
		s.lives += 100
	}
	// increment level after 10 items eaten
	if (s.level + 1) == (s.fruits+s.veggies+s.donuts)/10 {
		s.level++
		rl.PlaySound(s.sound)
	}
}
