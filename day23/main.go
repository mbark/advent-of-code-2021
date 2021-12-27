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
	fmt.Printf("first: %d\n", solve(state2))
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

	pos := make(map[maps.Coordinate]Amphipod)
	graph := make(Graph)
	for c, nn := range m {
		if !nn.empty {
			continue
		}

		if nn.amphi != "" {
			pos[c] = Amphipod(nn.amphi)
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

		graph[c] = Node{
			Coordinate:      c,
			IsHallway:       c.Y == 1,
			IsRoom:          c.Y > 1,
			IsLeftCorridor:  c.X < 3,
			IsRightCorridor: c.X > 9,
			IsAboveRoom:     c.Y == 1 && (c.X == 3 || c.X == 5 || c.X == 7 || c.X == 9),
			RoomFor:         forAmphipod,
		}
	}

	for c, n := range graph {
		var neighbors []Node
		for _, ac := range c.Adjacent() {
			nn, ok := graph[ac]
			if !ok {
				continue
			}

			neighbors = append(neighbors, nn)
		}

		n.Neighbors = neighbors
	}

	return State{Positions: pos, Graph: graph}
}

type State struct {
	Positions map[maps.Coordinate]Amphipod
	Graph     Graph
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
			a := s.Positions[c]
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

var ampphipodIndex = map[Amphipod]int{
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
	for c, a := range s.Positions {
		j := ampphipodIndex[a]
		h = h ^ bitstrings[13*c.Y+c.X][j]
	}

	return h
}

func (s State) IsDone() bool {
	for c, a := range s.Positions {
		n := s.Graph[c]
		if !n.IsRoom || n.RoomFor != a {
			return false
		}
	}

	return true
}

func (s State) IsEmpty(c maps.Coordinate) bool {
	_, ok := s.Positions[c]
	return !ok
}

func (s State) MovesFor(visited map[maps.Coordinate]struct{}, from maps.Coordinate, at maps.Coordinate, steps int, a Amphipod) []Move {
	var moves []Move

	// we've already been here!
	if _, ok := visited[at]; ok {
		return nil
	}

	visited[at] = struct{}{}
	node, exists := s.Graph[at]
	switch {
	case !exists:
		// can't go where there is nowhere to go
		return nil

	case !s.IsEmpty(at):
		// can't go here, it's already occupied
		return nil

	case node.IsRoom && node.RoomFor != a:
		// if we're in the wrong room move to the corridor
		return s.MovesFor(visited, from, maps.Coordinate{X: at.X, Y: 1}, steps+at.Y-1, a)

	case node.IsRoom && node.RoomFor == a:
		roomDone := true
		for down, ok := s.Graph[at.Down()]; ok; down, ok = s.Graph[down.Coordinate.Down()] {
			amphiDown := s.Positions[down.Coordinate]
			if amphiDown != node.RoomFor {
				roomDone = false
				break
			}
		}

		if roomDone {
			return []Move{NewMove(from, at, steps, a)}
		}

		moveDown := s.MovesFor(visited, from, at.Down(), steps+1, a)
		if moveDown != nil {
			return moveDown
		}

		return s.MovesFor(visited, from, at.Up(), steps+1, a)

	case node.IsHallway:
		down := s.Graph[at.Down()]
		if node.IsAboveRoom && down.RoomFor == a {
			// check if we can move into our room
			moveDown := s.MovesFor(visited, from, down.Coordinate, steps+1, a)
			if moveDown != nil {
				return moveDown
			}
		}

		started := s.Graph[from]
		if !node.IsAboveRoom && !started.IsHallway {
			moves = append(moves, NewMove(from, at, steps, a))
		}

		moves = append(moves, s.MovesFor(visited, from, at.Left(), steps+1, a)...)
		moves = append(moves, s.MovesFor(visited, from, at.Right(), steps+1, a)...)
	}

	return moves
}

func (s State) MoveFrom(at maps.Coordinate, a Amphipod) []Move {
	var moves []Move

	node := s.Graph[at]
	visited := make(map[maps.Coordinate]struct{})
	visited[at] = struct{}{}

	switch {
	case node.IsRoom && node.RoomFor != a:
		// if we're in the wrong room just move upwards
		return s.MovesFor(visited, at, at.Up(), 1, a)

	case node.IsRoom && node.RoomFor == a:
		down := s.Graph[at.Down()]
		amphiDown := s.Positions[down.Coordinate]

		switch {
		case amphiDown != down.RoomFor:
			// we're in the right room but the wrong amphipod is below, we need to move up
			return s.MovesFor(visited, at, at.Up(), 1, a)

		default:
			return nil
		}

	case node.IsHallway:
		switch {
		case node.IsLeftCorridor:
			// we can only move right from the left corridor
			return s.MovesFor(visited, at, at.Right(), 1, a)

		case node.IsRightCorridor:
			// we can only move left from the right corridor
			return s.MovesFor(visited, at, at.Left(), 1, a)

		case at.X > a.RoomX():
			return s.MovesFor(visited, at, at.Left(), 1, a)

		case at.X < a.RoomX():
			return s.MovesFor(visited, at, at.Right(), 1, a)
		}
	}

	return moves
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
	for c, a := range s.Positions {
		movesFor := s.MoveFrom(c, a)
		moves = append(moves, movesFor...)
	}

	return moves
}

func EndState(g Graph) State {
	pos := make(map[maps.Coordinate]Amphipod)
	for c, n := range g {
		if n.IsRoom {
			pos[c] = n.RoomFor
		}
	}

	return State{Positions: pos, Graph: g}
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
	news := make(map[maps.Coordinate]Amphipod, len(s.Positions))
	for k, v := range s.Positions {
		if k == m.From {
			continue
		}
		news[k] = v
	}
	news[m.To] = s.Positions[m.From]

	return State{Positions: news, Graph: s.Graph}
}

func (s State) Heuristic() int {
	var sum int

	for c, n := range s.Graph {
		a := s.Positions[c]
		if n.IsRoom && a != "" && n.RoomFor != a {
			yMove := c.Y - 1
			xMove := util.AbsInt(c.X - a.RoomX())
			sum += (yMove + xMove) * a.Cost()
		}

		if n.IsHallway && a != "" {
			xMove := util.AbsInt(c.X - a.RoomX())
			sum += xMove * a.Cost()
		}

		if n.IsRoom && a == "" {
			sum += (c.Y - 1) * n.RoomFor.Cost()
		}
	}

	return sum
}

type Node struct {
	Coordinate maps.Coordinate
	Neighbors  []Node

	IsHallway       bool
	IsLeftCorridor  bool // the corridor is the room to the left or right of the rooms
	IsRightCorridor bool // the corridor is the room to the left or right of the rooms
	IsAboveRoom     bool
	IsRoom          bool
	RoomFor         Amphipod
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
