package maps

import "C"
import (
	"container/heap"
	"fmt"
	"github.com/mbark/advent-of-code-2021/util"
	"strconv"
	"strings"
)

type Coordinate3D struct {
	X int
	Y int
	Z int
}

func NewCoordinate3D(s string) Coordinate3D {
	split := strings.Split(s, ",")
	return Coordinate3D{
		X: util.Str2Int(split[0]),
		Y: util.Str2Int(split[1]),
		Z: util.Str2Int(split[2]),
	}
}

func (c Coordinate3D) String() string {
	return fmt.Sprintf("(x=%d,y=%d,z=%d)", c.X, c.Y, c.Z)
}

func (c Coordinate3D) Diff(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: util.AbsInt(c.X - to.X),
		Y: util.AbsInt(c.Y - to.Y),
		Z: util.AbsInt(c.Z - to.Z),
	}
}

func (c Coordinate3D) ManhattanDistance(to Coordinate3D) int {
	d := c.Diff(to)
	return d.X + d.Y + d.Z
}

func (c Coordinate3D) Sub(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X - to.X,
		Y: c.Y - to.Y,
		Z: c.Z - to.Z,
	}
}

func (c Coordinate3D) Add(to Coordinate3D) Coordinate3D {
	return Coordinate3D{
		X: c.X + to.X,
		Y: c.Y + to.Y,
		Z: c.Z + to.Z,
	}
}

type Rotation interface {
	Apply(c Coordinate3D) Coordinate3D
}

type RotationDirection struct {
	X bool
	Y bool
	Z bool
}

func (r RotationDirection) Apply(c Coordinate3D) Coordinate3D {
	if r.X {
		c.X = c.X * -1
	}
	if r.Y {
		c.Y = c.Y * -1
	}
	if r.Z {
		c.Z = c.Z * -1
	}

	return c
}

type RotationFacing struct {
	X string
	Y string
	Z string

	Direction RotationDirection
}

func (r RotationFacing) Apply(c Coordinate3D) Coordinate3D {
	cnew := Coordinate3D{}
	switch r.X {
	case "x":
		cnew.X = c.X
	case "y":
		cnew.X = c.Y
	case "z":
		cnew.X = c.Z
	}

	switch r.Y {
	case "x":
		cnew.Y = c.X
	case "y":
		cnew.Y = c.Y
	case "z":
		cnew.Y = c.Z
	}

	switch r.Z {
	case "x":
		cnew.Z = c.X
	case "y":
		cnew.Z = c.Y
	case "z":
		cnew.Z = c.Z
	}

	return r.Direction.Apply(cnew)
}

func (c Coordinate3D) ApplyRotation(x, y, z int) Coordinate3D {
	return Coordinate3D{X: x * c.X, Y: y * c.Y, Z: z * c.Z}
}

func (c Coordinate3D) Rotations() []Coordinate3D {
	vals := []int{c.X, c.Y, c.Z}
	var permutations []Coordinate3D

	// each permutation of x,y,z
	for i, x := range vals {
		for j, y := range vals {
			for k, z := range vals {
				if i == j || j == k || i == k {
					continue
				}

				permutations = append(permutations, Coordinate3D{
					X: x,
					Y: y,
					Z: z,
				})
			}
		}
	}

	var coordinates []Coordinate3D
	for _, c := range permutations {
		for _, x := range []int{1, -1} {
			for _, y := range []int{1, -1} {
				for _, z := range []int{1, -1} {
					coordinates = append(coordinates,
						Coordinate3D{X: x * c.X, Y: y * c.Y, Z: z * c.Z})
				}
			}
		}
	}

	coordinates = append(coordinates, permutations...)
	fmt.Println(len(coordinates))

	return coordinates
}

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

type CoordinateItem struct {
	Coordinate Coordinate
	Priority int
	Index    int
}

type PriorityQueue []*CoordinateItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*CoordinateItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *CoordinateItem, value Coordinate, priority int) {
	item.Coordinate = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
