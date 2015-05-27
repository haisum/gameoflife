package gameoflife

// Point struct represents a single alive cell at location x,y of the life grid
type Point byte

type Grid struct {
	Rows, Columns uint
	Alive         map[int]map[int]Point
}

//checks 3x3 matix around the given point and returns total number of alive neighbours of point
func (g *Grid) TotalNeighbors(x int, y int) int {
	n := 0
	//loop from position x-1 to x+1 to check neighbor rows
	for i := -1; i <= 1; i++ {
		//loop from position y-1 to y+1 to check neighbor columns
		for j := -1; j <= 1; j++ {
			_, alive = g.Alive[i][j]
			if alive {
				n += 1
			}
		}
	}
}

func (g *Grid) Draw(UI u) {
	u.Draw(g.Rows, g.Columns, g.Alive)
}
