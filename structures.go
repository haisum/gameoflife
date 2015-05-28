package gameoflife

import (
	"time"
)

// Point struct represents a single alive cell at location x,y of the life grid
type Point byte

//Grid represents current life space of dimension Rows X Columns and Alive points
type Grid struct {
	Rows, Columns int
	Alive         map[int]map[int]Point
	RefreshRate   time.Duration
}

//checks 3x3 matix around the given point and returns total number of alive neighbours of point
func (g *Grid) totalNeighbors(x int, y int) int {
	n := 0
	//loop from position y-1 to y+1 to check neighbor columns
	for j := -1; j <= 1; j++ {
		//loop from position x-1 to x+1 to check neighbor rows
		for i := -1; i <= 1; i++ {
			nX := x + i
			nY := y + j
			//don't run out of bonds of life
			if nX >= 0 && nY >= 0 && nY < g.Columns && nX < g.Rows && (i != 0 || j != 0) {
				if _, alive := g.Alive[nX][nY]; alive {
					n += 1
				}
			}
		}
	}
	return n
}

// This function draws grid on UI
func (g Grid) Draw(u UI) {
	u.Draw(g)
}

// This function iterates through all cells of grid and sets alive cells for next generation of grid
func (g *Grid) Next() {
	alive := make(map[int]map[int]Point)
	for y := 0; y < g.Columns; y++ {
		for x := 0; x < g.Rows; x++ {
			n := g.totalNeighbors(x, y)
			_, isAlive := g.Alive[x][y]

			//Any live cell with fewer than two live neighbours dies, as if caused by under-population.

			//Any live cell with two or three live neighbours lives on to the next generation.
			if isAlive && (n == 2 || n == 3) {
				if _, ok := alive[x]; !ok {
					alive[x] = make(map[int]Point)
				}
				alive[x][y] = 1
			} else if !isAlive && n == 3 {
				//Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
				if _, ok := alive[x]; !ok {
					alive[x] = make(map[int]Point)
				}
				alive[x][y] = 1
			}
			//Any live cell with more than three live neighbours dies, as if by overcrowding
		}
	}
	g.Alive = alive
}
