package gameoflife

import (
	"strconv"
	"strings"
	"time"
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

// This type is representation of Grid object in json friendly structure and is used in converting Grid oject to JSON
type JGrid struct {
	Cells       [][]bool
	RefreshRate time.Duration
}

func GridtoJGrid(g Grid) JGrid {
	var j = make([][]bool, g.Rows)
	for x := 0; x < g.Rows; x++ {
		j[x] = make([]bool, g.Columns)
		for y := 0; y < g.Columns; y++ {
			if _, ok := g.Alive[x][y]; ok {
				j[x][y] = true
			} else {
				j[x][y] = false
			}
		}
	}
	jg := JGrid{
		Cells:       j,
		RefreshRate: g.RefreshRate / time.Millisecond,
	}
	return jg
}
