package main

import (
	"flag"
	"fmt"
	"github.com/haisum/gameoflife"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	var active = flag.String("a", "1:2,2:2,3:2", "List of alive cells. Format: x1:y1,x2:y2,x3:y3,....,xn:yn.")
	var columns = flag.Int("c", 5, "Number of columns in life space")
	var rows = flag.Int("r", 5, "Number of rows in life space")
	var textOnly = flag.Bool("t", false, "If passed, text only output is shown without any colors. Useful for systems where ansi coloring is not supported and program outputs garbage text.")

	flag.Parse()

	cells := make(map[int]map[int]gameoflife.Point)

	tokens := strings.Split(*active, ",")
	for _, p := range tokens {
		c := strings.Split(p, ":")
		//ignore bad coordinates
		if len(c) == 2 {
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])
			if _, ok := cells[x]; !ok {
				cells[x] = make(map[int]gameoflife.Point)
			}
			cells[x][y] = 1
		}
	}
	fmt.Println()

	g := gameoflife.Grid{
		Rows:    *rows,
		Columns: *columns,
		Alive:   cells,
	}

	ui := gameoflife.Terminal{TextOnly: runtime.GOOS == "windows" || *textOnly}

	g.Draw(ui)
	for {
		time.Sleep(time.Second / 2)
		g.Next()
		g.Draw(ui)
	}

}
