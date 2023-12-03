package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Score struct {
	fruits  int
	veggies int
	donuts  int
	level   int
}

func (s *Score) draw() {
	f := fmt.Sprintf("Fruits: %d", s.fruits)
	v := fmt.Sprintf("Veggies: %d", s.veggies)
	d := fmt.Sprintf("Donuts: %d", s.donuts)
	l := fmt.Sprintf("Level: %d", s.level)
	rl.DrawText(l, screenW-rl.MeasureText(l, 20)-7, 7, 20, rl.Red)
	rl.DrawText(d, screenW-rl.MeasureText(d, 20)-7, 30, 20, rl.DarkPurple)
	rl.DrawText(f, screenW-rl.MeasureText(f, 20)-7, 53, 20, rl.DarkPurple)
	rl.DrawText(v, screenW-rl.MeasureText(v, 20)-7, 76, 20, rl.DarkPurple)

}

func (s *Score) update(k kind) {
	switch k {
	case V:
		s.veggies++
	case F:
		s.fruits++
	case D:
		s.donuts++
	}
	// increment level after 10 items eaten
	s.level = (s.fruits + s.veggies + s.donuts) / 10
}
