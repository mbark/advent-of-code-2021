package maps

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) Up() Coordinate {
	return Coordinate{X: c.X, Y: c.Y - 1}
}

func (c Coordinate) Right() Coordinate {
	return Coordinate{X: c.X + 1, Y: c.Y}
}

func (c Coordinate) Down() Coordinate {
	return Coordinate{X: c.X, Y: c.Y + 1}
}

func (c Coordinate) Left() Coordinate {
	return Coordinate{X: c.X - 1, Y: c.Y}
}

func (c Coordinate) Adjacent() []Coordinate {
	return []Coordinate{
		{X: c.X, Y: c.Y + 1}, // up
		{X: c.X + 1, Y: c.Y}, // right
		{X: c.X, Y: c.Y - 1}, // down
		{X: c.X - 1, Y: c.Y}, // left
	}
}

func (c Coordinate) Surrounding() []Coordinate {
	return []Coordinate{
		{X: c.X, Y: c.Y - 1},     // N
		{X: c.X, Y: c.Y + 1},     // S
		{X: c.X + 1, Y: c.Y},     // W
		{X: c.X - 1, Y: c.Y},     // E
		{X: c.X + 1, Y: c.Y - 1}, // NE
		{X: c.X + 1, Y: c.Y + 1}, // SE
		{X: c.X - 1, Y: c.Y + 1}, // SW
		{X: c.X - 1, Y: c.Y - 1}, // NW
	}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(x=%d,y=%d)", c.X, c.Y)
}

