package maps

import (
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strconv"
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

type IntMap struct {
	Columns int
	Rows    int
	Cells   [][]int
}

func NewIntMap(definition string) IntMap {
	var cells [][]int

	var rows, cols int
	for y, l := range util.ReadInput(definition, "\n") {
		rows = y
		var row []int
		for x, n := range util.NumberList(l, "") {
			cols = x
			row = append(row, n)
		}

		cells = append(cells, row)
	}

	return IntMap{Columns: cols + 1, Rows: rows + 1, Cells: cells}
}

func MapFromCoordinates(coordinates []Coordinate) IntMap {
	var rows, cols int
	for _, c := range coordinates {
		if c.Y > rows {
			rows = c.Y
		}
		if c.X > cols {
			cols = c.X
		}
	}

	rows, cols = rows+1, cols+1

	cells := make([][]int, rows)
	for i := range cells {
		cells[i] = make([]int, cols)
	}

	for _, c := range coordinates {
		cells[c.Y][c.X] = 1
	}

	return IntMap{Columns: cols, Rows: rows, Cells: cells}
}

func (m IntMap) At(c Coordinate) int {
	return m.Cells[c.Y][c.X]
}

func (m IntMap) Coordinates() []Coordinate {
	coordinates := make([]Coordinate, m.Length())
	for y, row := range m.Cells {
		for x := range row {
			coordinates[y*m.Rows+x] = Coordinate{Y: y, X: x}
		}
	}

	return coordinates
}

func (m *IntMap) Set(c Coordinate, val int) {
	m.Cells[c.Y][c.X] = val
}

func (m *IntMap) Inc(c Coordinate) {
	m.Cells[c.Y][c.X] += 1
}

func (m IntMap) Exists(c Coordinate) bool {
	return c.X >= 0 && c.X < m.Columns &&
		c.Y >= 0 && c.Y < m.Rows
}

func (m IntMap) filterNonExistent(coords []Coordinate) []Coordinate {
	var cs []Coordinate
	for _, c := range coords {
		if m.Exists(c) {
			cs = append(cs, c)
		}
	}

	return cs
}

func (m IntMap) Adjacent(c Coordinate) []Coordinate {
	var coordinates []Coordinate
	for _, x := range []int{-1, 1} {
		coordinates = append(coordinates, Coordinate{X: c.X + x, Y: c.Y})
	}
	for _, y := range []int{-1, 1} {
		coordinates = append(coordinates, Coordinate{X: c.X, Y: c.Y + y})
	}

	return m.filterNonExistent(coordinates)
}

func (m IntMap) Surrounding(c Coordinate) []Coordinate {
	var coordinates []Coordinate
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			if x == 0 && y == 0 {
				continue
			}

			coordinates = append(coordinates, Coordinate{X: c.X + x, Y: c.Y + y})
		}
	}

	return m.filterNonExistent(coordinates)
}

func (m IntMap) String() string {
	var sb strings.Builder
	for _, row := range m.Cells {
		for _, cell := range row {
			sb.WriteString(strconv.Itoa(cell))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m IntMap) Length() int {
	return m.Rows * m.Columns
}
