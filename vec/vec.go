package vec

var (
	North = Coord{-1, 0}
	South = Coord{1, 0}
	East  = Coord{0, 1}
	West  = Coord{0, -1}

	Up    = Coord{-1, 0}
	Down  = Coord{1, 0}
	Right = Coord{0, 1}
	Left  = Coord{0, -1}
)

type Coord struct {
	X, Y int
}

func (lhs *Coord) Add(rhs Coord) Coord {
	return Coord{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func ManhattanDist(u Coord, v Coord) int {
	return absInt(u.X-v.X) + absInt(u.Y-v.Y)
}

func AllDirections() []Coord {
	return []Coord{Up, Down, Left, Right}
}

func (a *Coord) Dot(b Coord) int {
	return a.X*b.X + a.Y*b.Y
}

func absInt(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
