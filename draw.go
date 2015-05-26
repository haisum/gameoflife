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

// bash ansi color that we are going to use for displaying dead cells
var dead = ansi.ColorFunc("black+B:green")
var alive = ansi.ColorFunc("177+b:18")

// This function draws a x X y grid on terminal
// and highlights points at location defined by Point objects inside activePoints array
func Draw(x int, y int, activePoints []Point) {
	fmt.Println(strings.Repeat(" -", x))
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			point := Point{i, j}

			if i == 0 {
				fmt.Print("|")
			}

			printCell(point.InSlice(activePoints))

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
func printCell(isAlive bool) {
	if runtime.GOOS == "windows" {
		if isAlive {
			fmt.Printf("O")
		} else {
			fmt.Printf(" ")
		}
	} else {
		if isAlive {
			fmt.Printf(alive(" "))
		} else {
			fmt.Printf(dead(" "))
		}
	}
}

// Clears current console
// Equivalent to clscr or clear
func Clear() {
	fmt.Printf("\033[H\033[2J")
}
