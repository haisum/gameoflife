package main

import (
	"github.com/haisum/gameoflife"
)

func main() {
	active := make([]gameoflife.Point, 3)
	active[0] = gameoflife.Point{
		6,
		5,
	}
	active[1] = gameoflife.Point{
		5,
		5,
	}
	active[2] = gameoflife.Point{
		7,
		6,
	}
	gameoflife.Draw(10, 15, active)
}
