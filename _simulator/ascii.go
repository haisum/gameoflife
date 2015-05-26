package main

import (
	"github.com/haisum/gameoflife"
)

func main() {
	active := make([]gameoflife.Point, 3)
	active[0] = gameoflife.Point{
		1,
		2,
	}
	active[1] = gameoflife.Point{
		8,
		7,
	}
	active[2] = gameoflife.Point{
		0,
		0,
	}
	gameoflife.Draw(50, 50, active)
}
