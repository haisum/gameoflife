// Package gameoflife is implementation of Game of Life game by John Conway
//
// Installation
//
// go get github.com/haisum/gameoflife
// go run src/github.com/haisum/gameoflife/_simulator/terminal.go
//
// Usage
//
// Compile files in _simulator folder
// Enjoy yourself smooth simulation
package gameoflife

import (
	"fmt"
	"github.com/mgutz/ansi"
	"runtime"
	"strings"
)

type UI interface {
	Draw(x int, y int, c map[int]map[int]Point)
}

type Terminal struct {
	// bash ansi colors that we are going to use for displaying dead cells
	Dead, Alive func(string) string
}

// This function draws a x X y grid on terminal
// and highlights alive points in c, 2D map
func (t Terminal) Draw(x int, y int, c map[int]map[int]Point) {
	if t.Alive == nil {
		t.Alive = ansi.ColorFunc("177+b:18")
	}
	if t.Dead == nil {
		t.Dead = ansi.ColorFunc("black+B:green")
	}
	fmt.Println(strings.Repeat(" -", x))
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			//enabling this line while debugging helps
			//fmt.Printf("{%d,%d}", i, j)
			if i == 0 {
				fmt.Print("|")
			}
			//if i,j is present in c, cell is alive
			_, alive := c[i][j]
			t.printCell(alive)
			fmt.Printf("|")
			if i == x-1 {
				fmt.Println()
				fmt.Println(strings.Repeat(" -", x))
			}
		}
	}
	fmt.Printf(ansi.ColorCode("reset"))
}

// Prints a cell of a grid on terminal
func (t *Terminal) printCell(isAlive bool) {
	if runtime.GOOS == "windows" {
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

// Clears current console
// Equivalent to clscr or clear
func (t *Terminal) Clear() {
	fmt.Printf("\033[H\033[2J")
}
