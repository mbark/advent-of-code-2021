package maps

import "C"
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

func (m IntMap) ArraySize() int {
	return (m.Rows + 1) * (m.Columns + 1)
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

func (m IntMap) ArrPos(c Coordinate) int {
	return c.Y*m.Rows + c.X
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

func (m IntMap) CopyWith(fn func(val int) int) IntMap {
	cells := make([][]int, len(m.Cells))

	for i := range m.Cells {
		row := make([]int, len(m.Cells[i]))
		for j, cell := range m.Cells[i] {
			row[j] = fn(cell)
		}

		cells[i] = row
	}

	return IntMap{Columns: m.Columns, Rows: m.Rows, Cells: cells}
}

func Merged(maps [][]IntMap) IntMap {
	var cells [][]int
	var columns, rows int

	for _, row := range maps {
		rows += row[0].Rows
	}
	for _, col := range maps[0] {
		columns += col.Columns
	}

	// for each map in the row
	for _, mapRow := range maps {
		// for each row in the map
		for i := 0; i < mapRow[0].Rows; i++ {
			var row []int
			for _, mapCol := range mapRow {
				row = append(row, mapCol.Cells[i]...)
			}

			cells = append(cells, row)
		}
	}

	return IntMap{Columns: columns, Rows: rows, Cells: cells}
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
	coordinates := make([]Coordinate, 4)
	var at int
	for _, x := range []int{-1, 1} {
		c := Coordinate{X: c.X + x, Y: c.Y}
		if m.Exists(c) {
			coordinates[at] = c
			at += 1
		}
	}
	for _, y := range []int{-1, 1} {
		c := Coordinate{X: c.X, Y: c.Y + y}
		if m.Exists(c) {
			coordinates[at] = c
			at += 1
		}
	}

	return coordinates[:at]
}

func (m IntMap) Surrounding(c Coordinate) []Coordinate {
	var coordinates []Coordinate
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			if x == 0 && y == 0 {
				continue
			}

			c := Coordinate{X: c.X + x, Y: c.Y + y}
			if m.Exists(c) {
				coordinates = append(coordinates, c)
			}
		}
	}

	return coordinates
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
