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

//JCell is representation of single Alive cell's position in json friendly format
type JCell struct {
	X, Y int
}

// This type is representation of Grid object in json friendly structure and is used in converting Grid oject to JSON
type JGrid struct {
	Rows, Columns int
	Alive         []JCell
	RefreshRate   time.Duration
}

func GridtoJGrid(g Grid) JGrid {
	j := make([]JCell, 5)
	for k1, _ := range g.Alive {
		for k2, _ := range g.Alive[k1] {
			o := JCell{X: k1, Y: k2}
			j = append(j, o)
		}
	}
	jg := JGrid{
		Rows:        g.Rows,
		Columns:     g.Columns,
		Alive:       j,
		RefreshRate: g.RefreshRate / time.Millisecond,
	}
	return jg
}
