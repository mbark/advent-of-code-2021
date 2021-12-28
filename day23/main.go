package main

import (
	"container/heap"
	"fmt"
	"github.com/mbark/advent-of-code-2021/maps"
	"github.com/mbark/advent-of-code-2021/util"
	"math/rand"
	"strings"
)

const (
	testDone = `
#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
`

	testData = `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
`

	testMove = `
#############
#.A.........#
###A#.#C#D###
  #B#B#C#D#
  #########
`

	testData2 = `
#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########
`

	in = `
#############
#...........#
###D#A#A#D###
  #C#C#B#B#
  #########
`

	in2 = `
#############
#...........#
###D#A#A#D###
  #D#C#B#A#
  #D#B#A#C#
  #C#C#B#B#
  #########
`
)

type inNode struct {
	empty bool
	amphi string
}

func main() {
	f := util.WithProfiling()
	defer f()
	state := parseInput(in)
	state2 := parseInput(in2)

	fmt.Printf("first: %d\n", solve(state))
	fmt.Printf("second: %d\n", solve(state2))
}

func parseInput(in string) State {
	m := make(map[maps.Coordinate]inNode)
	for y, l := range util.ReadInput(in, "\n") {
		for x, c := range l {
			if c == ' ' {
				continue
			}

			isEmpty := c != '#'
			var amphi string
			if c != '.' && c != '#' {
				amphi = string(c)
			}

			m[maps.Coordinate{Y: y, X: x}] = inNode{empty: isEmpty, amphi: amphi}
		}
	}

	pos := make(map[Amphipod][]maps.Coordinate)
	graph := make(Graph)
	for c, nn := range m {
		if !nn.empty {
			continue
		}

		if nn.amphi != "" {
			pos[Amphipod(nn.amphi)] = append(pos[Amphipod(nn.amphi)], c)
		}

		var neighbors []maps.Coordinate
		for _, ac := range c.Adjacent() {
			v, ok := m[ac]
			if !ok || !v.empty {
				continue
			}

			neighbors = append(neighbors, ac)
		}

		var forAmphipod Amphipod
		switch {
		case c.Y > 1 && c.X == 3:
			forAmphipod = "A"
		case c.Y > 1 && c.X == 5:
			forAmphipod = "B"
		case c.Y > 1 && c.X == 7:
			forAmphipod = "C"
		case c.Y > 1 && c.X == 9:
			forAmphipod = "D"
		}

		graph[c] = Node{Coordinate: c, RoomFor: forAmphipod}
	}

	var positions Positions
	for a, c := range pos {
		positions = append(positions, AmphipodPositions{
			Amphipod:    a,
			Coordinates: c,
		})
	}

	return State{Positions: positions, Graph: graph}
}

type State struct {
	Positions Positions
	Graph     Graph
}

type AmphipodPositions struct {
	Amphipod    Amphipod
	Coordinates []maps.Coordinate
}

type Positions []AmphipodPositions

func (p Positions) At(c maps.Coordinate) Amphipod {
	for _, s := range p {
		for _, ac := range s.Coordinates {
			if ac == c {
				return s.Amphipod
			}
		}
	}

	return ""
}

func (s State) String() string {
	var coords []maps.Coordinate
	for c := range s.Graph {
		coords = append(coords, c)
	}

	m := maps.MapFromCoordinates(coords)
	var sb strings.Builder

	for y, row := range m.Cells {
		for x := range row {
			c := maps.Coordinate{X: x, Y: y}
			_, ok := s.Graph[c]
			a := s.Positions.At(c)
			switch {
			case !ok:
				sb.WriteString("#")

			case a != "":
				sb.WriteString(string(a))

			default:
				sb.WriteString(" ")
			}
		}
		sb.WriteString("#")
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	return sb.String()
}

var bitstrings [7 * 13][4]uint32
var boardSize = 7 * 13

var amphipodIndex = map[Amphipod]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
}

func init() {
	for i := 0; i < boardSize; i++ {
		var bits [4]uint32
		for j := 0; j < 4; j++ {
			bits[j] = rand.Uint32()
		}

		bitstrings[i] = bits
	}
}

// Hash calculates a Zobrist hash
func (s State) Hash() uint32 {
	var h uint32
	for _, a := range s.Positions {
		j := amphipodIndex[a.Amphipod]
		for _, c := range a.Coordinates {
			h = h ^ bitstrings[13*c.Y+c.X][j]
		}
	}

	return h
}

func (s State) IsDone() bool {
	for _, a := range s.Positions {
		for _, c := range a.Coordinates {
			n := s.Graph[c]
			if !n.IsRoom() || n.RoomFor != a.Amphipod {
				return false
			}
		}
	}

	return true
}

func (s State) IsEmpty(c maps.Coordinate) bool {
	a := s.Positions.At(c)
	return a == ""
}

func (s State) IsRoomDone(at maps.Coordinate) bool {
	node := s.Graph[at]
	for down, ok := s.Graph[at.Down()]; ok; down, ok = s.Graph[down.Coordinate.Down()] {
		amphiDown := s.Positions.At(down.Coordinate)
		if amphiDown != node.RoomFor {
			return false
		}
	}

	return true
}

func (s State) CanMoveIntoRoom(a Amphipod) bool {
	for down, ok := s.Graph[maps.Coordinate{Y: 2, X: a.RoomX()}]; ok; down, ok = s.Graph[down.Coordinate.Down()] {
		amphiDown := s.Positions.At(down.Coordinate)
		if amphiDown != "" && amphiDown != down.RoomFor {
			return false
		}
	}

	return true
}

func (s State) CanMoveToHallway(at maps.Coordinate) bool {
	for n := s.Graph[at.Up()]; n.Coordinate.Y > 1; n = s.Graph[n.Coordinate.Up()] {
		if !s.IsEmpty(n.Coordinate) {
			return false
		}
	}

	return true
}

func (s State) MoveInHallway(at maps.Coordinate) []maps.Coordinate {
	var coords []maps.Coordinate
	leftat := at.Left()
	for ; ; leftat = leftat.Left() {
		n, ok := s.Graph[leftat]
		if !ok || !s.IsEmpty(n.Coordinate) {
			break
		}

		if n.IsLeftCorridor() || !n.IsAboveRoom() {
			coords = append(coords, n.Coordinate)
		}
	}

	rightat := at.Right()
	for ; ; rightat = rightat.Right() {
		n, ok := s.Graph[rightat]
		if !ok || !s.IsEmpty(n.Coordinate) {
			break
		}

		if n.IsRightCorridor() || !n.IsAboveRoom() {
			coords = append(coords, n.Coordinate)
		}
	}

	return coords
}

func (s State) MoveToRoom(at maps.Coordinate, a Amphipod) *maps.Coordinate {
	// no point moving there
	if !s.CanMoveIntoRoom(a) {
		return nil
	}

	// assume we start in the hallway somewhere
	delta := 1
	if at.X > a.RoomX() {
		delta = -1
	}

	x := at.X + delta
	for ; x != a.RoomX(); x += delta {
		if !s.IsEmpty(maps.Coordinate{Y: at.Y, X: x}) {
			return nil
		}
	}

	y := at.Y
	for ; ; y++ {
		_, ok := s.Graph[maps.Coordinate{Y: y, X: x}]
		if !ok || !s.IsEmpty(maps.Coordinate{Y: y, X: x}) {
			y -= 1
			break
		}
	}

	return &maps.Coordinate{X: x, Y: y}
}

func (s State) MoveFrom(at maps.Coordinate, a Amphipod) []Move {
	node := s.Graph[at]
	switch {
	case node.IsRoom() && node.RoomFor != a:
		if !s.CanMoveToHallway(at) {
			return nil
		}

		aboveRoom := maps.Coordinate{X: at.X, Y: 1}
		steps := aboveRoom.ManhattanDistance(at)

		hallwayStops := s.MoveInHallway(aboveRoom)
		var moves []Move
		for _, c := range hallwayStops {
			moves = append(moves, NewMove(at, c, steps+c.ManhattanDistance(aboveRoom), a))
		}
		if c := s.MoveToRoom(aboveRoom, a); c != nil {
			moves = append(moves, NewMove(at, *c, steps+c.ManhattanDistance(aboveRoom), a))
		}

		return moves

	case node.IsRoom() && node.RoomFor == a:
		if s.IsRoomDone(at) {
			return nil
		}

		if !s.CanMoveToHallway(at) {
			return nil
		}

		aboveRoom := maps.Coordinate{X: at.X, Y: 1}
		steps := aboveRoom.ManhattanDistance(at)

		hallwayStops := s.MoveInHallway(aboveRoom)
		var moves []Move
		for _, c := range hallwayStops {
			moves = append(moves, NewMove(at, c, steps+c.ManhattanDistance(aboveRoom), a))
		}
		return moves

	case node.IsHallway():
		if c := s.MoveToRoom(at, a); c != nil {
			return []Move{NewMove(at, *c, c.ManhattanDistance(at), a)}
		}
	}

	return nil
}

func (a Amphipod) MoveCost(steps int) int {
	return a.Cost() * (steps)
}

func (a Amphipod) RoomX() int {
	switch a {
	case "A":
		return 3
	case "B":
		return 5
	case "C":
		return 7
	case "D":
		return 9
	}

	panic("invalid amphipod" + a)
	return -1
}

func (s State) Moves() []Move {
	var moves []Move
	for _, a := range s.Positions {
		for _, c := range a.Coordinates {
			movesFor := s.MoveFrom(c, a.Amphipod)
			moves = append(moves, movesFor...)
		}
	}

	return moves
}

func EndState(g Graph) State {
	pos := make(map[Amphipod][]maps.Coordinate)
	for c, n := range g {
		if n.IsRoom() {
			pos[n.RoomFor] = append(pos[n.RoomFor], c)
		}
	}

	var positions Positions
	for a, c := range pos {
		positions = append(positions, AmphipodPositions{Amphipod: a, Coordinates: c})
	}

	return State{Positions: positions, Graph: g}
}

func (a Amphipod) Cost() int {
	switch a {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	default:
		panic("Invalid Amphipod: " + a)
	}
}

type Move struct {
	From maps.Coordinate
	To   maps.Coordinate
	Cost int
}

func (m Move) String() string {
	return fmt.Sprintf("(%d) %s => %s", m.Cost, m.From, m.To)
}

func NewMove(from maps.Coordinate, to maps.Coordinate, steps int, a Amphipod) Move {
	return Move{To: to, From: from, Cost: a.MoveCost(steps)}
}

func (m Move) Do(s State) State {
	news := make(Positions, len(s.Positions))
	for i, a := range s.Positions {
		var moved bool
		for _, c := range a.Coordinates {
			if c == m.From {
				moved = true
				break
			}
		}

		if moved {
			coordinates := make([]maps.Coordinate, len(a.Coordinates))
			for i, c := range a.Coordinates {
				if c == m.From {
					coordinates[i] = m.To
				} else {
					coordinates[i] = c
				}
			}
			news[i] = AmphipodPositions{Amphipod: a.Amphipod, Coordinates: coordinates}
		} else {
			news[i] = a
		}
	}

	return State{Positions: news, Graph: s.Graph}
}

func (s State) Heuristic() int {
	var sum int

	for c, n := range s.Graph {
		a := s.Positions.At(c)
		if n.IsRoom() && a != "" && n.RoomFor != a {
			yMove := c.Y - 1
			xMove := util.AbsInt(c.X - a.RoomX())
			sum += (yMove + xMove) * a.Cost()
		}

		if n.IsHallway() && a != "" {
			xMove := util.AbsInt(c.X - a.RoomX())
			sum += xMove * a.Cost()
		}

		if n.IsRoom() && a == "" {
			sum += (c.Y - 1) * n.RoomFor.Cost()
		}
	}

	return sum
}

type Node struct {
	Coordinate maps.Coordinate
	RoomFor    Amphipod
}

func (n Node) IsHallway() bool {
	return n.Coordinate.Y == 1
}

func (n Node) IsLeftCorridor() bool {
	c := n.Coordinate
	return c.Y == 1 && c.X < 3
}

func (n Node) IsRightCorridor() bool {
	c := n.Coordinate
	return c.X > 9 && c.X < 3
}

func (n Node) IsRoom() bool {
	c := n.Coordinate
	return c.Y > 1 && (c.X == 3 || c.X == 5 || c.X == 7 || c.X == 9)
}

func (n Node) IsAboveRoom() bool {
	c := n.Coordinate
	return c.Y == 1 && (c.X == 3 || c.X == 5 || c.X == 7 || c.X == 9)
}

type Amphipod string

type Graph map[maps.Coordinate]Node

func solve(state State) int {
	return djikstra(state, EndState(state.Graph), state.Graph)
}

func djikstra(start State, end State, graph Graph) int {
	f := util.WithTime()
	defer f()

	cost := make(map[uint32]int)
	prev := make(map[uint32]State)

	var pq PriorityQueue
	heap.Init(&pq)
	heap.Push(&pq, &Item{State: start, Priority: 0})

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Item)
		if n.State.IsDone() {
			break
		}

		for _, move := range n.State.Moves() {
			nextState := move.Do(n.State)
			hash := nextState.Hash()

			newCost := cost[n.State.Hash()] + move.Cost
			currCost, ok := cost[hash]
			if !ok || newCost < currCost {
				prev[hash] = n.State
				cost[hash] = newCost
				heap.Push(&pq, &Item{State: nextState, Priority: newCost + nextState.Heuristic()})
			}
		}
	}

	p := EndState(graph)
	path := []State{p}
	for p.Hash() != start.Hash() {
		if len(p.Positions) == 0 {
			panic("no way back!")
		}

		path = append(path, p)
		hash := p.Hash()
		pr := prev[hash]
		p = pr
	}
	path = append(path, p)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return cost[end.Hash()]
}

type Item struct {
	State    State
	Priority int
	Index    int
}

type PriorityQueue []*Item

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
	item := x.(*Item)
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

func (pq *PriorityQueue) Update(item *Item, value State, priority int) {
	item.State = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
