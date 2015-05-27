package main

import (
	"fmt"
	"github.com/haisum/gameoflife"
	"strconv"
	"strings"
)

func main() {
	active := "1:1,2:3,0:0"

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
	fmt.Printf("%v", cells)
	gameoflife.Draw(10, 15, cells)
}
