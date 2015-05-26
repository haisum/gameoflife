package gameoflife

// Point struct represents a single cell at location x,y of the life grid
type Point struct {
	X int
	Y int
}

// Checks if a particular Point object is in supplied Point array
func (a *Point) InSlice(points []Point) bool {
	for _, b := range points {
		if b.X == a.X && b.Y == a.Y {
			return true
		}
	}
	return false
}
