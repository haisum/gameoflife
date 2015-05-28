package gameoflife

import (
	"strconv"
	"strings"
)

// Gets a string in format x1:y1,x2:y2,x3:y3,....,xn:yn then splits it in pairs of x,y locations
// and builds a map of format [x][y] = 1
func GetCells(s string) map[int]map[int]Point {
	cells := make(map[int]map[int]Point)

	tokens := strings.Split(s, ",")
	for _, p := range tokens {
		c := strings.Split(p, ":")
		//ignore bad coordinates
		if len(c) == 2 {
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])
			if _, ok := cells[x]; !ok {
				cells[x] = make(map[int]Point)
			}
			cells[x][y] = 1
		}
	}

	return cells
}
