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

type CoordinateArray struct {
	Coordinates []Coordinate

	Y int
	X int
}

func NewCoordinateArray(coords []Coordinate) CoordinateArray {
	var x, y int
	for _, c := range coords {
		if c.X > x {
			x = c.X
		}
		if c.Y > y {
			y = c.Y
		}
	}

	return CoordinateArray{Coordinates: coords, X: x, Y: y}
}

func (arr CoordinateArray) Size() int {
	return arr.Y * arr.X
}

func (arr CoordinateArray) Index(c Coordinate) int {
	return c.Y*arr.Y + c.X
}

func (arr CoordinateArray) Coordinate(i int) Coordinate {
	y := i / 13
	x := i % 13

	return Coordinate{Y: y, X: x}
}