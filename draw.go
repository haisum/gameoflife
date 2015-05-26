package gameoflife

import (
	"fmt"
	"github.com/mgutz/ansi"
	"strings"
)

var normal = ansi.ColorFunc("green+h:black")
var highlight = ansi.ColorFunc("black+B:green")

func Draw(x int, y int, activePoints []Point) {
	fmt.Println(strings.Repeat(" -", x))
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			point := Point{i, j}

			if i == 0 {
				fmt.Print("|")
			}

			if point.InSlice(activePoints) {
				fmt.Printf(highlight(" "))
			} else {
				fmt.Printf(normal(" "))
			}
			fmt.Printf("|")
			if i == x-1 {
				fmt.Println()
				fmt.Println(strings.Repeat(" -", x))
			}
		}
	}
	fmt.Printf(ansi.ColorCode("reset"))
}

func Clear() {
	fmt.Printf("\033[H\033[2J")
}
