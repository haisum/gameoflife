// Package gameoflife is implementation of Game of Life game by John Conway.
//
// This implementation of game of life is cross platform and can be run on any system for which golang programs can be compiled.
// Simulation is supported on terminal and as a http page and can be easily extended for other displays by implementing an interface for that display (see draw.go for examples)
//
// This simulation is memory efficient as it only records alive cells in a map.
// Iteration speed can definitely be further improved by utilizing HashLife and other algorithms in future.
//
// Installation
//
// 	cd $GOPATH
// 	go get github.com/haisum/gameoflife
// 	go build src/github.com/haisum/gameoflife/_simulator/simulator.go
//	./simulator
//
// Usage
//
// Build binary
//
// 	cd src/github.com/haisum/gameoflife
// 	go build _simulator/simulator.go
//
// Usage of ./simulator
//
// 	-a="1:2,2:2,3:2": List of alive cells. Format: x1:y1,x2:y2,x3:y3,....,xn:yn.
//  -d="terminal": Display interface for simulation. Two values are supported right now: "terminal" and "http"
//  -p=8001: Port number to listen on for http requests. Only applicable when -d is set to http
//  -r=500: Refresh rate for animation in milliseconds.
//  -t=false: Only applicable when -d is set to terminal. If passed, text only output is shown without any colors. Useful for systems where ansi coloring is not supported and program outputs garbage text.
//  -x=5: Number of rows in life space
//  -y=5: Number of columns in life space

//
// Example commands
//
// 	./simulator -x 9 -y 9 -a "3:3,4:3,5:3,3:4,4:4,5:4,3:5,4:5,5:5" -r 1000
// 	./simulator -x 15 -y 10 -r 500
// 	./simulator -h
//	./simulator -d http
//	./simulator -t
//
// Author: Haisum
package gameoflife

import (
	"encoding/json"
	"fmt"
	"github.com/mgutz/ansi"
	"log"
	"net/http"
	"strings"
	"time"
)

type UI interface {
	Draw(g Grid)
}

type Terminal struct {
	// bash ansi colors that we are going to use for displaying dead cells
	Dead, Alive func(string) string
	TextOnly    bool
}

// This function draws a x X y grid on terminal
// and highlights alive points in c, 2D map.
// It also regenerates grid after "r" nanoseconds
func (t *Terminal) Draw(g Grid) {
	if !t.TextOnly {
		if t.Alive == nil {
			t.Alive = ansi.ColorFunc("177+b:18")
		}
		if t.Dead == nil {
			t.Dead = ansi.ColorFunc("black+B:green")
		}
	}
	fmt.Println(strings.Repeat(" -", g.Rows))
	for j := 0; j < g.Columns; j++ {
		for i := 0; i < g.Rows; i++ {
			//enabling this line while debugging helps
			//fmt.Printf("{%d,%d}", i, j)
			if i == 0 {
				fmt.Print("|")
			}
			//if i,j is present in c, cell is alive
			_, alive := g.Alive[i][j]
			t.printCell(alive)
			fmt.Printf("|")
			if i == g.Rows-1 {
				fmt.Println()
				fmt.Println(strings.Repeat(" -", g.Rows))
			}
		}
	}
	fmt.Printf(ansi.ColorCode("reset"))
	time.Sleep(g.RefreshRate)
	g.Next()
	t.Draw(g)
}

// Prints a cell of a grid on terminal
func (t Terminal) printCell(isAlive bool) {
	if t.TextOnly {
		if isAlive {
			fmt.Printf("O")
		} else {
			fmt.Printf(" ")
		}
	} else {
		if isAlive {
			fmt.Printf(t.Alive(" "))
		} else {
			fmt.Printf(t.Dead(" "))
		}
	}
}

// This type defines UI Draw method for displaying simulation on http requests
type Http struct {
	serving bool
	Port    int
}

// Starts http server and draws simulation on http requests
func (h *Http) Draw(g Grid) {
	if !h.serving {
		//next outputs next state of the grid
		http.HandleFunc("/next", func(w http.ResponseWriter, req *http.Request) {
			data, err := json.Marshal(GridtoJGrid(g))
			if err != nil {
				fmt.Fprint(w, err)
			} else {
				fmt.Fprintf(w, "%s", data)
			}
			g.Next()
		})
		//serve html javascript from static dir
		http.Handle("/", http.FileServer(http.Dir("static")))
		h.serve()
	}
}

// Starts the http server if server hasn't yet been started. This function will only be called on first call to Draw method
func (h *Http) serve() {
	fmt.Printf("\nListening on http://localhost:%d you can see simulation by opening this url in a browser.\n", h.Port)
	h.serving = true
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.Port), nil)
	if err != nil {
		h.serving = false
		log.Fatal(fmt.Sprintf("Could not start server on port %d. Error: ", h.Port), err)
	}
}
