package main

import (
	"fmt"
	"github.com/haisum/gameoflife"
	"strconv"
	"strings"
	"time"
)

func main() {
	active := "1:2,2:2,3:2"

	cells := make(map[int]map[int]gameoflife.Point)

	tokens := strings.Split(active, ",")
	for _, p := range tokens {
		c := strings.Split(p, ":")
		//ignore bad coordinates
		if len(c) == 2 {
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])
			cells[x] = make(map[int]gameoflife.Point)
			cells[x][y] = 1
		}
	}
	fmt.Println()

	g := gameoflife.Grid{
		Rows:    5,
		Columns: 5,
		Alive:   cells,
	}

	ui := gameoflife.Terminal{}

	g.Draw(ui)
	for {
		time.Sleep(time.Second / 2)
		g.Next()
		g.Draw(ui)
	}

}
