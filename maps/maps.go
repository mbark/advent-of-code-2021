package maps

import (
	"fmt"
	"sort"
	"strings"
)

type Coordinate struct {
	X int
	Y int
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

func MapString(m map[Coordinate]int) string {
	yms := make(map[int]struct{})
	for c := range m {
		yms[c.Y] = struct{}{}
	}

	var ys []int
	for y := range yms {
		ys = append(ys, y)
	}

	sort.Ints(ys)

	var sb strings.Builder
	for _, y := range ys {
		xms := make(map[int]struct{})
		for c := range m {
			if c.Y != y {
				continue
			}
			xms[c.X] = struct{}{}
		}

		var xs []int
		for x := range xms {
			xs = append(xs, x)
		}

		sort.Ints(xs)
		for _, x := range xs {
			sb.WriteString(fmt.Sprintf("%2d", m[Coordinate{X: x, Y: y}]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
