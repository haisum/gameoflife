package main

import (
	"flag"
	"fmt"
	"github.com/haisum/gameoflife"
	"runtime"
	"time"
)

func main() {

	var active = flag.String("a", "1:2,2:2,3:2", "List of alive cells. Format: x1:y1,x2:y2,x3:y3,....,xn:yn.")
	var columns = flag.Int("y", 5, "Number of columns in life space")
	var rows = flag.Int("x", 5, "Number of rows in life space")
	var textOnly = flag.Bool("t", false, "If passed, text only output is shown without any colors. Useful for systems where ansi coloring is not supported and program outputs garbage text.")
	var refreshRate = flag.Int64("r", 500, "Refresh rate for animation in milliseconds.")
	flag.Parse()

	fmt.Println()

	ui := gameoflife.Terminal{TextOnly: runtime.GOOS == "windows" || *textOnly}

	g := gameoflife.Grid{
		Rows:        *rows,
		Columns:     *columns,
		Alive:       gameoflife.GetCells(*active),
		RefreshRate: time.Millisecond * time.Duration(*refreshRate),
	}
	g.Draw(ui)

}
