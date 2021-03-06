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
	var textOnly = flag.Bool("t", false, "Only applicable when -d is set to terminal. If passed, text only output is shown without any colors. Useful for systems where ansi coloring is not supported and program outputs garbage text.")
	var refreshRate = flag.Int64("r", 500, "Refresh rate for animation in milliseconds.")
	var display = flag.String("d", "terminal", "Display interface for simulation. Two values are supported right now: \"terminal\" and \"http\"")
	var port = flag.Int("p", 8001, "Port number to listen on for http requests. Only applicable when -d is set to http")

	flag.Parse()

	fmt.Println()

	var ui gameoflife.UI

	//we need a reference to ui rather than value so we could re-use the object instead of copying it, hence & was used
	switch *display {
	case "http":
		ui = &gameoflife.Http{Port: *port}
	default:
		ui = &gameoflife.Terminal{TextOnly: runtime.GOOS == "windows" || *textOnly}
	}

	g := gameoflife.Grid{
		Rows:        *rows,
		Columns:     *columns,
		Alive:       gameoflife.GetCells(*active),
		RefreshRate: time.Millisecond * time.Duration(*refreshRate),
	}
	g.Draw(ui)

}
