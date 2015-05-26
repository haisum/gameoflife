package gameoflife

type Point struct {
	X int
	Y int
}

func (a *Point) InSlice(points []Point) bool {
	for _, b := range points {
		if b.X == a.X && b.Y == a.Y {
			return true
		}
	}
	return false
}
